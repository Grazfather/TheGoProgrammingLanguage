// Exercise 4.11:
// Build a tool that lets users create, read, update, and close GitHub issues
// from the command line, invoking their preferred text editor when substantian
// text input is required.
package main

import (
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
	fmt.Println("usage: gitcli create <repo> <title> <body>")
}

func read_usage() {
	fmt.Println("usage: gitcli read <repo> <#>")
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
		if len(os.Args) != 5 {
			create_usage()
			return
		}
		issue, err := mygithub.CreateIssue(os.Args[2], os.Args[3], os.Args[4], token)
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
	}
}
