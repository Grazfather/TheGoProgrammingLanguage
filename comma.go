package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	if s == "" {
		return ""
	}
	buf := bytes.Buffer{}

	// Handle optional sign
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		s = s[1:]
	}

	// First write everything before the first comma
	dotindex := strings.LastIndexByte(s, '.')
	var n int
	var intpart, decimalpart string
	if dotindex != -1 {
		intpart = s[:dotindex]
		decimalpart = s[dotindex+1:]
	} else {
		intpart = s[:]
		decimalpart = ""
	}
	n = len(intpart) % 3

	if n == 0 && len(intpart) > 0 {
		n = 3
	}
	buf.WriteString(intpart[:n])
	intpart = intpart[n:]

	// Now loop 3 chars at a time, with a comma preceding
	for ; len(intpart) >= 3; intpart = intpart[3:] {
		buf.WriteByte(',')
		buf.WriteString(intpart[:3])
	}

	// If there's a fractional part, write it.
	if len(decimalpart) > 0 {
		buf.WriteByte('.')
		buf.WriteString(decimalpart)
	}

	return buf.String()
}

func main() {
	fmt.Println(comma(""))
	fmt.Println(comma("123"))
	fmt.Println(comma("12345"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("12345678"))
	fmt.Println(comma("1234567890123"))
	// Floats
	fmt.Println(comma("+"))
	fmt.Println(comma("."))
	fmt.Println(comma("1."))
	fmt.Println(comma(".5"))
	fmt.Println(comma("12.3"))
	fmt.Println(comma("123.45"))
	fmt.Println(comma("+123456"))
	fmt.Println(comma("-12345678"))
	fmt.Println(comma("-1234567.890123"))
}
