package blogger

import (
	"encoding/json"
	"net/url"

	"github.com/pkg/errors"
)

const endPoint = "https://www.googleapis.com/blogger/v3"

type Service struct {
	applicationKey string
	blogID         string
	getter         HTTPGetter
}

func NewService(getter HTTPGetter, applicationKey string) *Service {
	return &Service{
		applicationKey: applicationKey,
		getter:         getter,
	}
}

func (s *Service) ByURL(blogURL string) (Blog, error) {
	bl := Blog{}
	base, err := url.Parse(endPoint + "/blogs/byurl")
	if err != nil {
		return bl, err
	}
	u := base
	q := u.Query()
	q.Set("url", blogURL)
	q.Set("key", s.applicationKey)
	u.RawQuery = q.Encode()

	b, _, err := s.getter.Get(u.String())
	if err != nil {
		return bl, errors.Wrapf(err, "Can't get %s: %v", base, err)
	}

	err = json.Unmarshal(b, &bl)
	s.blogID = bl.ID
	return bl, err
}

func (s *Service) List(pageToken string) (BlogList, error) {
	bl := BlogList{}
	base, err := url.Parse(endPoint + "/blogs/" + s.blogID + "/posts")
	if err != nil {
		return bl, err
	}

	u := base
	q := u.Query()
	q.Set("key", s.applicationKey)
	q.Set("maxResults", "100")
	if len(pageToken) > 0 {
		q.Set("pageToken", pageToken)
	}
	u.RawQuery = q.Encode()
	b, _, err := s.getter.Get(u.String())
	if err != nil {
		return bl, errors.Wrapf(err, "Can't get %s: %v", base, err)
	}
	err = json.Unmarshal(b, &bl)
	return bl, err
}
