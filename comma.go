package main

import (
	"bytes"
	"fmt"
)

// command inserts commands in a non-negative decimal integer string.
func comma_recursive(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma(s string) string {
	size := (len(s)-1)/3 + len(s)
	buf := bytes.Buffer{}
	nums := s[:]
	buf.Grow(size)

	// First write everything before the first comma
	n := len(s) % 3
	if n == 0 && len(s) > 0 {
		n = 3
	}
	buf.WriteString(nums[:n])
	nums = nums[n:]

	// Now loop 3 chars at a time, with a comma preceding
	for ; len(nums) >= 3; nums = nums[3:] {
		buf.WriteByte(',')
		buf.WriteString(nums[:3])
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
}
