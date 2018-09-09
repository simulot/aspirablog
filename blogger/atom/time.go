package atom

import (
	"encoding/json"
	"time"
)

const BlogSpotTS = "2006-01-02T15:04:05.999-07:00"

type DateBlogSpot struct {
	time.Time
}

func (t *DateBlogSpot) UnmarshalJSON(b []byte) error {
	s := &Field{}
	err := json.Unmarshal(b, s)
	if err != nil {
		return err
	}
	v, err := time.Parse(BlogSpotTS, s.Content)
	if err != nil {
		return err
	}
	t.Time = v
	return nil
}
