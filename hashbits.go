package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"os"
)

func countBinaryHamminDistance(b1, b2 []byte) (count int, err error) {
	if len(b1) != len(b2) {
		return 0, errors.New("Cannot compare strings of unequal length")
	}

	for i, b := range b1 {
		for bi := uint(0); bi < 8; bi++ {
			if (b >> bi & 1) != (b2[i] >> bi & 1) {
				count += 1
			}
		}
	}
	return
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: hashbits STRING1 STRING2")
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]
	h1 := sha256.Sum256([]byte(s1))
	h2 := sha256.Sum256([]byte(s2))
	fmt.Printf("%s: %x\n%s: %x\n", s1, h1, s2, h2)
	count, _ := countBinaryHamminDistance(h1[:], h2[:])
	fmt.Printf("Bits different: %d out of %d\n", count, sha256.Size*8)
}
