package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	SHA256 = iota
	SHA384
	SHA512
)

func hash(algo int, data []byte) (hash []byte, err error) {
	if algo == SHA256 {
		s := sha256.Sum256(data)
		hash = s[:]
	} else if algo == SHA384 {
		s := sha512.Sum384(data)
		hash = s[:]
	} else if algo == SHA512 {
		s := sha512.Sum512(data)
		hash = s[:]
	} else {
		err = errors.New("Invalid hash algorithm")
	}
	return
}

func main() {
	var algo int
	do384 := flag.Bool("sha384", false, "Use sha384")
	do512 := flag.Bool("sha512", false, "Use sha512")
	flag.Parse()

	if *do384 && *do512 {
		fmt.Println("Error: Only specify one hash algorithm")
		os.Exit(1)
	}

	if *do384 {
		algo = SHA384
	} else if *do512 {
		algo = SHA512
	} else {
		algo = SHA256
	}

	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error: Could not read from STDIN")
		os.Exit(1)
	}

	h, _ := hash(algo, buf)
	fmt.Printf("%x\n", h)
}
