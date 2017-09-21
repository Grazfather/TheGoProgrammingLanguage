// Exercise 4.4
// Write a version of rotate that operates in a single pass.
package main

import "fmt"

// rotate rotates the elements in s by n spots to the left. 0 <= n <= len(s)
func rotate(s []int, n int) {
	f := make([]int, n, n)
	// Save first n elements
	copy(f, s)

	// Shift elements from n to end
	i := 0
	for ; i < len(s)-n; i++ {
		s[i] = s[n+i]
	}

	// Copy the backup to the tail
	copy(s[i:], f)
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(s)
	rotate(s, 3)
	fmt.Println(s)
}
