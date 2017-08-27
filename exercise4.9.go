// Exercise 4.9:
// Write a program wordfreq to report the frequency of each word in an input
// text file. Call input.Split(bufio.ScanWords) before the first call to Scan
// to break the input into words instead of lines.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordcounts := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wordcounts[scanner.Text()]++
	}
	fmt.Printf("word\tcount\n")
	for word, n := range wordcounts {
		fmt.Printf("%s\t%d\n", word, n)
	}
}
