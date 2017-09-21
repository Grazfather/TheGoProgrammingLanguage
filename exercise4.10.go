// Exercise 4.10:
// Modify issues to report the results in age categories, say less than a month
// old, less than a year old, and more than a year old.

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"gopl.io/ch4/github"
)

var filter = flag.String("age", "", "An age filter in the form [>,<][0-9]+{hdmy} e.g. <3m")

func getFilterFunc(filter string) (func(*github.Issue) bool, error) {
	if filter == "" {
		return func(*github.Issue) bool { return true }, nil
	}
	var lt bool
	if filter[0] == '<' {
		lt = true
	} else if filter[0] == '>' {
		lt = false
	} else {
		return nil, errors.New("Cannot parse age filter (>/< missing)")
	}

	num, err := strconv.Atoi(filter[1 : len(filter)-1])
	if err != nil {
		return nil, errors.New("Cannot parse age filter (number)")
	}

	var threshold time.Duration
	switch filter[len(filter)-1] {
	case 'h':
		threshold = time.Hour * time.Duration(num)
	case 'd':
		threshold = time.Hour * time.Duration(num) * 24
	case 'm':
		threshold = time.Hour * time.Duration(num) * 24 * 30
	case 'y':
		threshold = time.Hour * time.Duration(num) * 24 * 365
	default:
		return nil, errors.New("Cannot parse age filter (hdmy missing)")
	}

	return func(issue *github.Issue) bool {
		if lt {
			if time.Since(issue.CreatedAt) < threshold {
				return true
			}
			return false
		}
		if time.Since(issue.CreatedAt) > threshold {
			return true
		}
		return false
	}, nil
}

func main() {
	flag.Parse()
	filterFunc, err := getFilterFunc(*filter)
	if err != nil {
		log.Fatal(err)
	}

	result, err := github.SearchIssues(flag.Args())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Issues:")
	for _, item := range result.Items {
		if filterFunc(item) == true {
			fmt.Printf("#%-5d %s %9.9s %.55s\n",
				item.Number, item.CreatedAt.Format("2006 Jan _2"), item.User.Login, item.Title)
		}
	}
}
