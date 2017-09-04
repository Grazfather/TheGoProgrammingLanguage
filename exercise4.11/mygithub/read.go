package mygithub

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopl.io/ch4/github"
)

func ListIssues(repo, state, oauth string) (*[]Issue, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(IssuesURL, repo)+"?state="+state, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", oauth))
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("listing issues failed: %s", resp.Status)
	}

	// TODO: Handle pagination
	var issues []Issue
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		return nil, err
	}
	return &issues, nil
}

func SearchIssues(repo string, terms []string, oauth string) (*github.IssuesSearchResult, error) {
	terms = append(terms, "repo:"+repo)
	return github.SearchIssues(terms)
}

func ReadIssue(repo, number, oauth string) (*Issue, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(IssueURL, repo, number), nil)
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

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}
