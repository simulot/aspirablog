package blogger

import (
	"fmt"
	"log"

	"github.com/pkg/errors"

	"github.com/simulot/aspirablog/blog"
	"github.com/simulot/aspirablog/blogger/atom"
)

type HTTPGetter interface {
	IsUpdated(url string) (bool, error)
	Get(url string) ([]byte, string, error)
}

type Blogger struct {
	http HTTPGetter
	url  string
	name string
}

func New(name string, url string, http HTTPGetter) *Blogger {
	return &Blogger{
		name: name,
		url:  url,
		http: http,
	}
}

func (b Blogger) feedUrl(number int) string {
	return fmt.Sprintf("%s/feeds/posts/default?alt=json&max-results=%d", b.url, number)
}

func (b *Blogger) GetAll() (blog.Blog, error) {
	bl := blog.Blog{}
	url := b.feedUrl(10)
	buffer, _, err := b.http.Get(url)
	if err != nil {
		return bl, errors.Wrapf(err, "Can't GetAll %s: %v", b.name, err)
	}

	feed, err := atom.Decode(buffer)
	if err != nil {
		return bl, err
	}

	bl.Title = blog.TextLinter(feed.Feed.Title.Content)
	bl.Posts = []blog.Post{}

	for e, entry := range feed.Feed.Entries {
		post := blog.Post{
			Title:      blog.AllCapLinter(entry.Title.Content),
			Published:  entry.Published.Time,
			Categories: []string{},
		}

		for _, c := range entry.Categories {
			post.Categories = append(post.Categories, string(c))
		}
		log.Printf("Parsing entry %d: %s\n", e, post.Title)

		par, err := b.parsePost(entry.Content.Content)
		if err != nil {
			log.Print("Can't parse blogger entry #", e)
			continue
		}
		post.Paragraph = par
		bl.Posts = append(bl.Posts, post)
	}
	return bl, err
}
