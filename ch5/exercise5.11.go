// Exercise 5.11
// The instructor of the linear algebra course decides that calculus is now a
// prerequisite. Extend the topoSort function to report cycles.
package main

import (
	"fmt"
	"os"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"linear algebra":        {"calculus"}, // This is creates a cycle
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	order, err := topoSort(prereqs)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) (order []string, err error) {
	allDone := make(map[string]bool)
	var visitAll func(base string, items []string)
	visitAll = func(base string, items []string) {
		for _, item := range items {
			// seen is true as soon as we start this class chain,
			// but done isn't true until we've exhausted all its
			// prereqs
			done, seen := allDone[item]
			if seen && !done {
				err = fmt.Errorf("Cycle detected: %s was already seen in this chain\n", item)
			}
			if !seen {
				// We aren't done with this course's prereqs yet, but we've seen the course itself
				allDone[item] = false
				visitAll(item, m[item])
				allDone[item] = true
				order = append(order, item)
			}
		}
	}

	for k := range m {
		// Convert the class as a single item slice so it's considered its own prerequisit
		visitAll(k, []string{k})
		if err != nil {
			return nil, err
		}
	}
	return order, err
}
