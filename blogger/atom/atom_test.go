package atom

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/alecthomas/repr"
)

func mustParseTime(s string) DateBlogSpot {
	t, err := time.Parse(BlogSpotTS, s)
	if err != nil {
		panic(err)
	}
	return DateBlogSpot{t}
}
func Test_ReadAtom(t *testing.T) {

	expected := &Atom{
		Encoding: "UTF-8",
		Version:  "1.0",
		Feed: Feed{
			ID:       "tag:blogger.com,1999:blog-ID",
			Title:    Field{Content: "Blog Title", Type: "text"},
			Subtitle: Field{Content: "Blog subtitle", Type: "html"},
			Updated:  mustParseTime("2018-08-31T03:23:32.630+02:00"),
			Authors: []Author{
				"Author1 name",
				"Author2 name",
			},
			Categories: []Category{
				"Cat1",
				"Cat2",
			},
			Entries: []Entry{
				{
					ID: "tag:blogger.com,1999:blog-id.post-postid",
					Authors: []Author{
						"Author1 name",
					},
					Categories: []Category{
						"Cat1",
					},
					Title: Field{
						Content: "Post title",
						Type:    "text",
					},
					Content: Field{
						Content: "Post content",
						Type:    "html",
					},
					Published: mustParseTime("2018-08-22T17:15:00.005+02:00"),
					Updated:   mustParseTime("2018-08-22T17:15:52.863+02:00"),
				},
			},
		},
	}
	f, err := os.Open(filepath.Join("testdata", "atom.json"))
	if err != nil {
		t.Error(err)
		return
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error(err)
		return
	}

	atom := &Atom{}
	err = json.Unmarshal(b, atom)
	if err != nil {
		t.Error(err)
		return
	}
	o := []repr.Option{
		repr.Hide(&time.Location{}),
	}
	if !reflect.DeepEqual(expected, atom) {
		t.Errorf("Expecting %s\n got %s", repr.String(expected, o...), repr.String(atom, o...))
	}
}
