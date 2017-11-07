// Exercise 5.8:
// Modify forEachNode so that the pre and post functions return a boolean
// result indicating whether to continue the traversal. Use it to write a
// function ElementByID with the following signature that finds the first HTML
// element with the specified id attribute. The function should stop the
// traversal as soon as a match is found.
// func ElementByID(doc *html.Node, id string) *html.Node
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: %s URL ID", os.Args[0])
		os.Exit(1)
	}
	fetchElement(os.Args[1], os.Args[2])
}

func ElementByID(doc *html.Node, id string) *html.Node {
	finder := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == "id" && attr.Val == id {
					return false
				}
			}
		}
		return true
	}
	return forEachNode(doc, finder, nil)
}

func fetchElement(url, id string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	n := ElementByID(doc, id)
	if n == nil {
		return fmt.Errorf("no element with id %s found\n", id)
	}
	fmt.Printf("Found node:\n%+v\n", n)
	return nil
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
// If either function returns false, then stop traversion and return the
// current node.
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	if pre != nil {
		if pre(n) != true {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if n := forEachNode(c, pre, post); n != nil {
			return n
		}
	}

	if post != nil {
		if post(n) != true {
			return n
		}
	}
	return nil
}
