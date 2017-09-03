package mygithub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"gopl.io/ch4/github"
)

const CreateIssuesURL = "https://api.github.com/repos/%s/issues"

type CreateIssueRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func CreateIssue(repo, title, body, oauth string) (*Issue, error) {
	content, err := json.Marshal(&CreateIssueRequest{
		Title: title,
		Body:  body,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf(CreateIssuesURL, repo), bytes.NewBuffer(content))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", oauth))
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("creating issue failed: %s", resp.Status)
	}

	var issue github.Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &Issue{issue}, nil
}
