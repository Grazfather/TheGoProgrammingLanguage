package mygithub

const IssuesURL = "https://api.github.com/repos/%s/issues"
const IssueURL = IssuesURL + "/%s"

type IssueRequest struct {
	Title    string `json:"title,omitempty"`
	Body     string `json:"body,omitempty"`
	State    string `json:"state,omitempty"`
	Assignee string `json:"assignee,omitempty"`
}
