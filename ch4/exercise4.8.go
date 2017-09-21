// Exercise 4.8:
// Modify charcount to count letters, digits, and so on in their Unicode
// categories, using functions like unicode.IsLetter.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

var types = [...]struct {
	name string
	test func(rune) bool
}{
	{"Control", unicode.IsControl},
	{"Digit", unicode.IsDigit},
	{"Graphic", unicode.IsGraphic},
	{"Letter", unicode.IsLetter},
	{"Lower", unicode.IsLower},
	{"Mark", unicode.IsMark},
	{"Number", unicode.IsNumber},
	{"Print", unicode.IsPrint},
	{"Punct", unicode.IsPunct},
	{"Space", unicode.IsSpace},
	{"Symbol", unicode.IsSymbol},
	{"Title", unicode.IsTitle},
	{"Upper", unicode.IsUpper},
}

func main() {
	counts := make(map[rune]int)       // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int    // count of lengths of UTF-8 encodings
	invalid := 0                       // count of invalid UTF-8 characters
	typecounts := make(map[string]int) // count of each unicode character type

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		for _, t := range types {
			if t.test(r) {
				typecounts[t.name]++
			}
		}
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\ntype\tcount\n")
	for _, t := range types {
		if n := typecounts[t.name]; n > 0 {
			fmt.Printf("%s\t%d\n", t.name, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
