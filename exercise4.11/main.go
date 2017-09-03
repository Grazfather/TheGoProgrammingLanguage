// Exercise 4.11:
// Build a tool that lets users create, read, update, and close GitHub issues
// from the command line, invoking their preferred text editor when substantian
// text input is required.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Grazfather/GoPL/exercise4.11/mygithub"
)

var (
	command string
)

func usage() {
	fmt.Println("usage: gocli <cmd> [<args>]")
}

func create_usage() {
	fmt.Println("usage: gitcli create <repo> <title> ...")
}

func read_usage() {
	fmt.Println("usage: gitcli read <repo> <#>")
}

func update_usage() {
	fmt.Println("usage: gitcli update <repo> <#> ...")
}

func close_usage() {
	fmt.Println("usage: gitcli close <repo> <#>")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}
	token := os.Getenv("GITHUB_OAUTH")
	if token == "" {
		fmt.Println("You must set GITHUB_OAUTH")
		return
	}
	switch os.Args[1] {
	case "create":
		if len(os.Args) < 4 {
			create_usage()
			return
		}
		createFlag := flag.NewFlagSet("gocli create <repo> <title>", flag.ExitOnError)
		body := createFlag.String("body", "", "Issue body")
		assignee := createFlag.String("assignee", "", "Assignee")
		createFlag.Parse(os.Args[4:])
		opts := mygithub.IssueRequest{
			Title:    os.Args[3],
			Body:     *body,
			Assignee: *assignee,
		}
		issue, err := mygithub.CreateIssue(os.Args[2], opts, token)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Success!")
		fmt.Println(issue)
	case "read":
		if len(os.Args) != 4 {
			read_usage()
			return
		}
		issue, err := mygithub.ReadIssue(os.Args[2], os.Args[3], token)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(issue)
	case "update":
		if len(os.Args) < 4 {
			update_usage()
			return
		}
		updateFlag := flag.NewFlagSet("gocli update <repo> <#>", flag.ExitOnError)
		title := updateFlag.String("title", "", "New title")
		body := updateFlag.String("body", "", "New body")
		state := updateFlag.String("state", "", "New state ('open' or 'closed')")
		updateFlag.Parse(os.Args[4:])
		opts := mygithub.IssueRequest{
			Title: *title,
			Body:  *body,
			State: *state,
		}
		issue, err := mygithub.UpdateIssue(os.Args[2], os.Args[3], opts, token)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Success!")
		fmt.Println(issue)
	case "close":
		if len(os.Args) < 4 {
			close_usage()
			return
		}
		err := mygithub.CloseIssue(os.Args[2], os.Args[3], token)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Success!")
	}
}
