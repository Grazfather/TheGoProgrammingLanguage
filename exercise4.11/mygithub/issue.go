package mygithub

import (
	"fmt"

	"gopl.io/ch4/github"
)

type Issue struct {
	github.Issue
}

func (i *Issue) String() string {
	return fmt.Sprintf(
		`Issue %d (%s)
%s (%s)
Created by %s at %v

%s`, i.Number, i.HTMLURL, i.Title, i.State, i.User.Login, i.CreatedAt, i.Body)
}
