// Exercise 3.13:
// Write const declarations for KB, MB, up through YB as compactly as you can.
package main

import (
	"fmt"
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println("KB:\t", KB)
	fmt.Println("MB:\t", MB)
	fmt.Println("GB:\t", GB)
	fmt.Println("TB:\t", TB)
	fmt.Println("PB:\t", PB)
	fmt.Println("EB:\t", EB)
	fmt.Println("YB and ZB both overflow 64 bit, but we can do constant math on them at compile time:")
	fmt.Println("YB/ZB:\t", YB/ZB)
}
