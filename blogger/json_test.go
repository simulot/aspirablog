package blogger

import (
	"encoding/json"
	"testing"
)

func Test_Blog_JSON(t *testing.T) {
	text := `{
		"kind": "blogger#blog",
		"id": "123456768987",
		"name": "Blog Name",
		"description": "Blog descripion",
		"published": "2013-12-25T18:06:56+01:00",
		"updated": "2018-09-07T10:00:10+02:00",
		"url": "http://blog.blogspot.com/",
		"selfLink": "https://www.googleapis.com/blogger/v3/blogs/123456768987",
		"posts": {
		 "totalItems": 268,
		 "selfLink": "https://www.googleapis.com/blogger/v3/blogs/123456768987/posts"
		},
		"pages": {
		 "totalItems": 0,
		 "selfLink": "https://www.googleapis.com/blogger/v3/blogs/123456768987/pages"
		},
		"locale": {
		 "language": "fr",
		 "country": "",
		 "variant": ""
		}
	   }`

	b := &Blog{}
	json.Unmarshal([]byte(text), b)

	_ = b
}
