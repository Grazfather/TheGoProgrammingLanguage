package main

import (
	"io/ioutil"
	"os"
	"testing"

	"golang.org/x/net/html"
)

// TestOutput tests to make sure that the output of outline is itself parseable HTML
func TestOutput(t *testing.T) {
	urls := []string{
		"http://example.com",
		"http://grazfather.github.io",
	}

	origStdout := os.Stdout
	for _, url := range urls {
		f, err := ioutil.TempFile("", "parsertest")
		if err != nil {
			t.Fatal("could not open temp file:", err)
		}
		defer f.Close()
		defer os.Remove(f.Name())

		os.Stdout = f
		outline(url)

		_, err = html.Parse(os.Stdout)
		if err != nil {
			t.Fatal("error in parser:", err)
		}
	}
	os.Stdout = origStdout
}
