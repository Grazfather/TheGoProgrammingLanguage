// Exercise 5.13:
// Modify crawl to make local copies of the pages it finds, creating
// directories as necessary. Don't make copies of pages that come from a
// different domain. For example, if the original page comes from golang.org,
// save all files from there, but exclude ones from vimeo.com
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"gopl.io/ch5/links"
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

var host string

func crawl(rawurl string) []string {

	u, err := url.Parse(rawurl)
	if err != nil {
		log.Print(err)
		return []string{}
	}

	// First URL crawled is the domain we care about
	if host == "" {
		host = u.Hostname()
	}

	if u.Hostname() != host {
		return []string{}
	}

	// Download
	err = download(rawurl)
	if err != nil {
		log.Print(err)
	}

	list, err := links.Extract(rawurl)
	if err != nil {
		log.Print(err)
	}
	return list
}

func download(rawurl string) error {
	fmt.Println(rawurl)
	resp, err := http.Get(rawurl)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("unexpected HTTP status code: %d", resp.StatusCode)
	}

	u, err := url.Parse(rawurl)
	if err != nil {
		return err
	}
	dir := path.Dir(u.EscapedPath())
	filename := path.Base(u.EscapedPath())
	if strings.HasSuffix(rawurl, "/") || filename == "." {
		filename = "index.html" // Dumb assumption
	}
	cwd, _ := os.Getwd()
	fmt.Println(dir, filename)
	dir = filepath.Join(cwd, dir)

	content, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, content, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Craw the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}
