package mygithub

const CreateIssuesURL = "https://api.github.com/repos/%s/issues"
const IssueURL = CreateIssuesURL + "/%s"

type IssueRequest struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
	State string `json:"state,omitempty"`
}
