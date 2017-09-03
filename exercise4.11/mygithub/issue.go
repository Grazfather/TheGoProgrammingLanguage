package mygithub

import (
	"fmt"

	"gopl.io/ch4/github"
)

type Issue struct {
	github.Issue
	Assignee *github.User
}

func (i *Issue) String() string {
	return fmt.Sprintf(
		`Issue %d (%s)
%s (%s)
Created by %s at %v
Assigned to %s

%s`, i.Number, i.HTMLURL, i.Title, i.State, i.User.Login, i.CreatedAt, i.Assignee.Login, i.Body)
}
