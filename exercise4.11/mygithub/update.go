package mygithub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func UpdateIssue(repo, number string, request IssueRequest, oauth string) (*Issue, error) {
	content, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}

	resp, err := makeGithubRequest("PATCH", fmt.Sprintf(IssueURL, repo, number), bytes.NewBuffer(content), oauth)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("updating issue failed: %s", resp.Status)
	}

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

func CloseIssue(repo, number, oauth string) error {
	request := IssueRequest{State: "closed"}
	_, err := UpdateIssue(repo, number, request, oauth)
	return err
}
