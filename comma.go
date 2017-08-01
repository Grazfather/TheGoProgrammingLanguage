package main

import "fmt"

// command inserts commands in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	fmt.Println(comma("123"))
	fmt.Println(comma("12345"))
	fmt.Println(comma("12345678"))
}
