// Exercise 5.15: Write variadic functions max and min, analogous to sum. What
// should these functions do when called with no arguments? Write variants that
// require at least one argument.

package main

import "fmt"

func max(vals ...int32) int32 {
	max := int32(-0x80000000)
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}

func max2(first int32, vals ...int32) int32 {
	max := first
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}

func min(vals ...int32) int32 {
	min := int32(0x7fffffff)
	for _, v := range vals {
		if v < min {
			min = v
		}
	}
	return min
}

func min2(first int32, vals ...int32) int32 {
	min := first
	for _, v := range vals {
		if v < min {
			min = v
		}
	}
	return min
}

func main() {
	fmt.Println(max(-0x80000000, -20000, 30000, 0x80000000-1))
	fmt.Println(max())
	fmt.Println(min(-0x80000000, -20000, 30000, 0x80000000-1))
	fmt.Println(min())
	fmt.Println(max2(-0x80000000, -20000, 30000, 0x80000000-1))
	fmt.Println(min2(-0x80000000, -20000, 30000, 0x80000000-1))
}
