// Exercise 4.7:
// Modify reverse to reverse the characters of a []byte slice that represents
// a UTF-8-encoding string, in place. Can you do it without allocating new
// memory?
package main

import (
	"fmt"
	"unicode/utf8"
)

func reverseUtf8(s []byte) {
	for r := 0; r < len(s); {
		_, size := utf8.DecodeRune(s[r:])
		// Reverse the bytes of this rune in place
		reverse(s[r : r+size])
		r += size
	}
	// Reverse the whole slice
	reverse(s)
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	s := []byte("啂哃哋hello！我菲爱宝实")
	fmt.Println(s)
	fmt.Println(string(s))
	reverseUtf8(s)
	fmt.Println(s)
	fmt.Println(string(s))
}
