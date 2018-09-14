package blog

import "time"

type Blog struct {
	Title       string    // Blog title as retrieved from service
	Description string    // Blog subtitle
	URL         string    // Blog url
	Authors     []string  // Author list
	Updated     time.Time // Last update by authors
	Published   time.Time // Blog opening time
	Label       []string  // Category list as retrieved from the service
	Posts       []Post    // Posts retrieved from the service
}
