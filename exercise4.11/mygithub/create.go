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

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}
