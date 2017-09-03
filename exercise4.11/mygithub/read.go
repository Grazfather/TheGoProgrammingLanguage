package mygithub

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopl.io/ch4/github"
)

const GetIssueURL = "https://api.github.com/repos/%s/issues/%s"

func ReadIssue(repo, number, oauth string) (*Issue, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(GetIssueURL, repo, number), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", oauth))
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("reading issue failed: %s", resp.Status)
	}

	var issue github.Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &Issue{issue}, nil
}
