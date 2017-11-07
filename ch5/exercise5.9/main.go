// Exercise 5.9
// Write a function expand(s string, f func(string) string) string that
// replaces each substring "$foo" within s by the text returned by f("foo").
package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	cap := func(s string) string {
		return "**" + s + "**"
	}
	fmt.Println(expand("this is a $test.", cap))
	fmt.Println(expand("this is a $PATH.", os.Getenv))
}

// expand finds occurrences of '$<s>' and replaces them with '$<f(s)>' and
// returns that new string.
func expand(s string, f func(string) string) string {

	re := regexp.MustCompile("\\$\\w+")
	new := re.ReplaceAllFunc([]byte(s), func(s []byte) []byte {
		return []byte(f(string(s[1:])))
	})
	return string(new)
}
