package blogger

import "time"

type Blog struct {
	Description string `json:"description"`
	ID          string `json:"id"`
	Kind        string `json:"kind"`
	Locale      struct {
		Country  string `json:"country"`
		Language string `json:"language"`
		Variant  string `json:"variant"`
	} `json:"locale"`
	Name  string `json:"name"`
	Pages struct {
		SelfLink   string `json:"selfLink"`
		TotalItems int    `json:"totalItems"`
	} `json:"pages"`
	Posts struct {
		SelfLink   string `json:"selfLink"`
		TotalItems int    `json:"totalItems"`
	} `json:"posts"`
	Published time.Time `json:"published"`
	SelfLink  string    `json:"selfLink"`
	Updated   time.Time `json:"updated"`
	URL       string    `json:"url"`
}

type BlogList struct {
	Kind          string `json:"kind"`
	NextPageToken string `json:"nextPageToken"`
	Items         []Post `json:"items"`
}

type Post struct {
	Kind string `json:"kind"`
	ID   string `json:"id"`
	Blog struct {
		ID string `json:"id"`
	} `json:"blog"`
	Published time.Time `json:"published"`
	Updated   time.Time `json:"updated"`
	Etag      string    `json:"etag"`
	URL       string    `json:"url"`
	SelfLink  string    `json:"selfLink"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    struct {
		ID          string `json:"id"`
		DisplayName string `json:"displayName"`
		URL         string `json:"url"`
		Image       struct {
			URL string `json:"url"`
		} `json:"image"`
	} `json:"author"`
	Replies struct {
		TotalItems string `json:"totalItems"`
		SelfLink   string `json:"selfLink"`
	} `json:"replies"`
	Labels []string `json:"labels"`
}
