// Exercise 4.5
// Write an in-place function to eliminate adjacent duplicates in a []string slice.
package main

import "fmt"

// dedpupe removes all ajacent duplicates from the provided slice and returns the next slice
func dedupe(s []string) []string {
	w := 0 // Write cursor
	c := 0 // Compare cursor
	r := 1 // Read cursor
	for ; r < len(s); r++ {
		// Zoom forward past all dupes
		for ; r < len(s) && s[c] == s[r]; r++ {
		}
		if r >= len(s) {
			break
		}
		// Write cursor only moves forward one
		w++
		// Copy the new string
		s[w] = s[r]
		// Now we'll compare to what we just wrote
		c = r
	}
	return s[:w+1]
}

func main() {
	s := []string{"a", "b", "a", "a", "b", "b", "a", "a", "a", "c", "c"}
	fmt.Println(s)
	s = dedupe(s)
	fmt.Println(s)
}
