package blogger

import (
	"log"

	"github.com/pkg/errors"

	"github.com/simulot/aspirablog/blog"
)

type HTTPGetter interface {
	IsUpdated(url string) (bool, error)
	Get(url string) ([]byte, string, error)
}

type Blogger struct {
	service *Service
	http    HTTPGetter
	url     string
	name    string
}

func New(name string, url string, apiKey string, http HTTPGetter) *Blogger {
	return &Blogger{
		service: NewService(http, apiKey),
		name:    name,
		url:     url,
	}
}

func (b *Blogger) GetAll() (blog.Blog, error) {
	bl := blog.Blog{}
	blBlog, err := b.service.ByURL(b.url)
	if err != nil {
		return bl, errors.Wrapf(err, "Can't GetAll %s: %v", b.name, err)
	}

	bl = blog.Blog{
		Title:       blBlog.Name,
		Description: blBlog.Description,
		Published:   blBlog.Published,
		Updated:     blBlog.Updated,
		URL:         blBlog.URL,
		Posts:       []blog.Post{},
	}

	nextPageToken := ""
	for {
		list, err := b.service.List(nextPageToken)
		if err != nil {
			return bl, errors.Wrapf(err, "Can't GetAll %s: %v", b.name, err)
		}
		for _, item := range list.Items {
			par, err := b.parsePost(item.Content)
			if err != nil {
				log.Print("Can't parse blogger entry #", err)
				continue
			}
			bl.Posts = append(
				bl.Posts,
				blog.Post{
					Author:     item.Author.DisplayName,
					Categories: item.Labels,
					Published:  item.Published,
					Title:      blog.AllCapLinter(item.Title),
					URL:        item.URL,
					Paragraph:  par,
				},
			)
		}
		nextPageToken = list.NextPageToken
		if len(nextPageToken) == 0 {
			break
		}
	}

	return bl, err
}
