// Exercise 5.7:
// Develop startElement and endElement into a general HTML prett-printer. Print comment nodes, text nodes, and the attributes of each element (<a href='...'>). Use short forms like <img/> instead of <img></img> when an element has no children. Write a test to ensure that the output can be parsed successfully. (See Chapter 11.)
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"unicode"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

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

var depth int

func startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		for _, attr := range n.Attr {
			fmt.Printf(" %s='%s'", attr.Key, attr.Val)
		}
		if n.FirstChild == nil {
			fmt.Printf("/>\n")
		} else {
			fmt.Printf(">\n")
		}
		depth++
	case html.TextNode:
		s := strings.TrimFunc(n.Data, unicode.IsSpace)
		if len(s) != 0 {
			depth++
			fmt.Printf("%*s%s\n", depth*2, "", s)
			depth--
		}
	case html.CommentNode:
		fmt.Printf("%*s<!--%s-->", depth*2, "", n.Data)
	case html.DoctypeNode:
		fmt.Printf("<!doctype %s>\n", n.Data)
	}
}
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild == nil {
		} else {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}
