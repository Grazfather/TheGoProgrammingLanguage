// Exercise 4.6:
// Write an in-place function that squashes each run of adjacent Unicode spaces
// (see unicode.IsSpace) in UTF-8-encoded []byte slice into a single ASCII space.
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// dedupeSpaces replaces all runs of any unicode space with a single ASCII space
func dedupeSpaces(s []byte) []byte {
	w := 0 // Write cursor
	size := 0
	for r := 0; r < len(s); {
		c, size := utf8.DecodeRune(s[r:])
		// If it's a space, zoom ahead
		if unicode.IsSpace(c) {
			// Replace it with an ascii space
			s[w] = ' '
			// Zoom forward past any spaces
			r += size
			c, size = utf8.DecodeRune(s[r:])
			for unicode.IsSpace(c) {
				r += size
				c, size = utf8.DecodeRune(s[r:])
			}
			w++
			// If we're reached the end, bail
			if r >= len(s) {
				break
			}
		} else { // Otherwise, copy the whole rune
			if w != r {
				copy(s[w:], s[r:r+size])
			}
			w += size
			r += size
		}
	}
	return s[:w+size]
}

func main() {
	s := []byte(`	 	且  界 	`)
	fmt.Println("Before:")
	fmt.Printf("% x\n", s)
	fmt.Println(string(s))
	s = dedupeSpaces(s)
	fmt.Println("After:")
	fmt.Printf("% x\n", s)
	fmt.Println(string(s))
}
