package atom

import "encoding/json"

type Atom struct {
	Encoding string `json:"encoding"`
	Version  string `json:"version"`
	Feed     Feed   `json:"feed"`
}

type GDImage struct {
	Height string `json:"height"`
	Rel    string `json:"rel"`
	Src    string `json:"src"`
	Width  string `json:"width"`
}

type Property struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type Link struct {
	Href  string `json:"href"`
	Rel   string `json:"rel"`
	Title string `json:"title"`
	Type  string `json:"type"`
}

type Thumbnail struct {
	Height string `json:"height"`
	URL    string `json:"url"`
	Width  string `json:"width"`
}

type Entry struct {
	ID         ID           `json:"id"`
	Authors    []Author     `json:"author"`
	Categories []Category   `json:"category"`
	Title      Field        `json:"title"`
	Content    Field        `json:"content"`
	Published  DateBlogSpot `json:"published"`
	Updated    DateBlogSpot `json:"updated"`
	// Properties []Property `json:"gd$extendedProperty"`
	// // Links      []Link       `json:"link"`
	// Thumbnail Thumbnail    `json:"media$thumbnail"`
	// // ThrTotal   Field        `json:"thr$total"`
}

type Feed struct {
	ID         ID           `json:"id"`
	Authors    []Author     `json:"author"`
	Categories []Category   `json:"category"`
	Subtitle   Field        `json:"subtitle"`
	Title      Field        `json:"title"`
	Updated    DateBlogSpot `json:"updated"`
	Entries    []Entry      `json:"entry"`
	// Generator  Field        `json:"generator"`
	// Links      []Link       `json:"link"`
	// OpenSearch_itemsPerPage struct {
	// 	_t string `json:"$t"`
	// } `json:"openSearch$itemsPerPage"`
	// OpenSearch_startIndex struct {
	// 	_t string `json:"$t"`
	// } `json:"openSearch$startIndex"`
	// OpenSearch_totalResults struct {
	// 	_t string `json:"$t"`
	// } `json:"openSearch$totalResults"`
	// Xmlns            string `json:"xmlns"`
	// Xmlns_blogger    string `json:"xmlns$blogger"`
	// Xmlns_gd         string `json:"xmlns$gd"`
	// Xmlns_georss     string `json:"xmlns$georss"`
	// Xmlns_openSearch string `json:"xmlns$openSearch"`
	// Xmlns_thr        string `json:"xmlns$thr"`
}

func Decode(b []byte) (Atom, error) {

	atom := Atom{}
	err := json.Unmarshal(b, &atom)
	return atom, err
}
