// Exercise 5.16
// Write a variadic version of strings.Join

package main

import (
	"bytes"
	"fmt"
)

func join(sep string, a ...string) string {
	if len(a) == 0 {
		return ""
	}

	var buf bytes.Buffer
	for _, s := range a[0 : len(a)-1] {
		buf.WriteString(s)
		buf.WriteString(sep)
	}
	buf.WriteString(a[len(a)-1])

	return buf.String()
}

func main() {
	fmt.Println(join("x"))
	fmt.Println(join("x", "a"))
	fmt.Println(join("x", "a", "b"))
	fmt.Println(join("x", "a", "b", "c"))
}
