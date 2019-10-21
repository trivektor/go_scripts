package structs

// Pull request
type PullRequest struct {
	URL    string  `json:"url"`
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	User   User    `json:"user"`
	Labels []Label `json:"labels"`
}
