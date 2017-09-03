package mygithub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"gopl.io/ch4/github"
)

func UpdateIssue(repo, number string, request IssueRequest, oauth string) (*Issue, error) {
	content, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf(IssueURL, repo, number), bytes.NewBuffer(content))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", oauth))
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("updating issue failed: %s", resp.Status)
	}

	var issue github.Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &Issue{issue}, nil
}

func CloseIssue(repo, number, oauth string) error {
	request := IssueRequest{State: "closed"}
	_, err := UpdateIssue(repo, number, request, oauth)
	return err
}
