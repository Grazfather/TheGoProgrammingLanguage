// Exercise 5.5:
// Implement countWordsAndImages. (See Exercise 4.9 for word-splitting.)
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "usage: %s URL", os.Args[0])
		os.Exit(1)
	}

	words, images, err := CountWordsAndImages(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("Words: %d, images: %d\n", words, images)
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	} else if n.Type == html.TextNode {
		scanner := bufio.NewScanner(strings.NewReader(n.Data))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			words++
		}
	}

	w1, i1 := countWordsAndImages(n.FirstChild)
	w2, i2 := countWordsAndImages(n.NextSibling)
	return words + w1 + w2, images + i1 + i2
}
