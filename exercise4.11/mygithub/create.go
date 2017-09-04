package mygithub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateIssue(repo string, request IssueRequest, oauth string) (*Issue, error) {
	content, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}

	resp, err := makeGithubRequest("POST", fmt.Sprintf(IssuesURL, repo), bytes.NewBuffer(content), oauth)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("creating issue failed: %s", resp.Status)
	}

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}
