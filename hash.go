package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"
)

func main() {
	//var algo int
	algo := flag.String("hash", "sha256", "hashtype: sha256, sha384, or sha512")
	flag.Parse()

	var hash hash.Hash
	switch *algo {
	case "sha256":
		hash = sha256.New()
	case "sha384":
		hash = sha512.New384()
	case "sha512":
		hash = sha512.New()
	default:
		fmt.Println("Error: Only sha256, sha384, and sha512 supported")
		os.Exit(1)
	}

	io.Copy(hash, os.Stdin)

	h := hash.Sum(nil)
	fmt.Printf("%x\n", h)
}
