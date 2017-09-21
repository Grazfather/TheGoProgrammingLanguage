package mygithub

import (
	"fmt"
	"io"
	"net/http"
)

const IssuesURL = "https://api.github.com/repos/%s/issues"
const IssueURL = IssuesURL + "/%s"

type IssueRequest struct {
	Title    string `json:"title,omitempty"`
	Body     string `json:"body,omitempty"`
	State    string `json:"state,omitempty"`
	Assignee string `json:"assignee,omitempty"`
}

func makeGithubRequest(method, url string, body io.Reader, oauth string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", oauth))
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)

	return resp, nil
}
