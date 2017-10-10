package github

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	HTMLURL   string `json:"html_url"`
	Number    int
	Title     string
	State     string
	User      *User
	MileStone *MileStone
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// MileStone
type MileStone struct {
	Title   string
	HTMLURL string `json:"html_url"`
}
