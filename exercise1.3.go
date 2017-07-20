package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var s, sep string
	now := time.Now()
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println("Elapsed: ", time.Since(now).Nanoseconds())

	now = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println("Elapsed: ", time.Since(now).Nanoseconds())
}
