// Exercise 5.14:
// Use the breadthFirst function to explor a different structure.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func walk(dir string) []string {
	fmt.Println(dir)
	subdirs := make([]string, 0)
	dirs, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for _, subdir := range dirs {
		if subdir.IsDir() {
			subdirs = append(subdirs, path.Join(dir, subdir.Name()))
		}
	}

	return subdirs
}

func main() {
	// Walk the directory breadth-first starting from the command-line
	// arguments.
	breadthFirst(walk, os.Args[1:])
}
