package main

import (
	"fmt"
	"html/template"
	"os"

	"github.com/Grazfather/TheGoProgrammingLanguage/ch4/exercise4.11/mygithub"
)

const issuesTemplate = `
<html>
<body>
<title>Issues lol</title>
<ol>
{{range .}}
<li>
Number: {{.Number}}<br/>
User: {{.User.Login}}<br/>
Title: <a href="{{.HTMLURL}}">{{.Title | printf "%.64s"}}</a><br/>
Created: {{.CreatedAt}}<br/>
</li>
{{end}}
</ol>
</body>
</html>`

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s URL\n", os.Args[0])
		return
	}
	token := os.Getenv("GITHUB_OAUTH")
	if token == "" {
		fmt.Println("You must set GITHUB_OAUTH")
		return
	}

	issues, err := mygithub.ListIssues(os.Args[1], "all", token)
	if err != nil {
		fmt.Println(err)
	}

	report, err := template.New("report").Parse(issuesTemplate)
	if err != nil {
		fmt.Println(err)
	}

	if err := report.Execute(os.Stdout, issues); err != nil {
		fmt.Println(err)
	}
}
