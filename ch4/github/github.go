package github

import (
	"time"
)

const IssuesURL = "https://docs.github.com/rest"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}
type Issue struct {
	Number   int
	HTMLURL  string `json:"html_url`
	Title    string
	State    string
	User     *User
	CreateAt time.Time `json:"create_at"`
	Body     string
}
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
