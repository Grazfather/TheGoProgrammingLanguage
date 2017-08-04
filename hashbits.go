package main

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"os"

	"popcount"
)

// Count the number of different bits between two equal-length byte slices. Their length must be a multiple of 8.
func countBinaryHammingDistance(b1, b2 []byte) (count int, err error) {
	if len(b1) != len(b2) {
		return 0, errors.New("Cannot compare strings of unequal length")
	}

	if len(b1)%8 != 0 {
		return 0, errors.New("Cannot compare strings whose lengths aren't a multiple of 8")
	}

	for len(b1) > 0 {
		s1 := binary.LittleEndian.Uint64(b1[:8])
		s2 := binary.LittleEndian.Uint64(b2[:8])
		count += popcount.PopCountTable(s1 ^ s2)
		b1 = b1[8:]
		b2 = b2[8:]
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
	count, _ := countBinaryHammingDistance(h1[:], h2[:])
	fmt.Printf("Bits different: %d out of %d\n", count, len(h1)*8)
}
