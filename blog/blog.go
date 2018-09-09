package blog

import "time"

type Blog struct {
	Title       string            // Blog title as retrieved from service
	URL         string            // Blog url
	Authors     []string          // Author list
	LastUpdated time.Time         // Last update by authors
	LastChecked time.Time         // Last check
	Categories  []string          // Category list as retrieved from the service
	CategoryMap map[string]string // Simplify the categories
	Posts       []Post            // Posts retrieved from the service
}
