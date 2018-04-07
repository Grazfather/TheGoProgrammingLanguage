// Exercise 5.17
// Write a variadic function ElementsByTagName that, given an HTML node tree
// and zero or more names, returns all the elements that match one of those
// names.

package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	matches := make([]*html.Node, 0)
	tags := make(map[string]bool)
	for _, n := range name {
		tags[n] = true
	}

	var appendMatches = func(n *html.Node) {
		if n.Type != html.ElementNode {
			return
		}

		if _, ok := tags[n.Data]; ok {
			matches = append(matches, n)
		}
	}
	forEachNode(doc, appendMatches, nil)

	return matches
}

// Taken from exercise 5.12
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func main() {
	resp, err := http.Get("http://example.com")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println(ElementsByTagName(doc, "a", "html", "body", "div"))
	fmt.Println(ElementsByTagName(doc, "meta"))
	fmt.Println(ElementsByTagName(doc))
}
