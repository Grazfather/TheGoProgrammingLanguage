// Exercise 5.3:
// Write a function to print the contents of all text nodes in an HTML document
// tree. Do not descend into <script> or <style> elements, since their contents
// are not visible in a web browser.
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
	visit(doc)
}

// visit prints the content of all non-script, non-style text nodes.
func visit(n *html.Node) {
	if n == nil || (n.Type == html.ElementNode && (n.Data == "style" || n.Data == "script")) {
		return
	}
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}

	visit(n.FirstChild)
	visit(n.NextSibling)
}
