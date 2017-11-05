// Exercise 5.2:
// Write a function to populate a mapping from elelment names--p, div, span,
// and so on-- to the number of elements with that name in an HTML document
// tree.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(1)
	}
	counts := make(map[string]int)
	visit(counts, doc)

	fmt.Println(counts)
}

// visit keeps a count of all visited node by the tag name
func visit(counts map[string]int, n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		counts[n.Data]++
	}

	visit(counts, n.NextSibling)
	visit(counts, n.FirstChild)
}
