// This program finds the SHA256 checksum of a given file specified as
// the first command line argument
package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file>\n", os.Args[0])
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Opening file error: %s\n", err.Error())
		os.Exit(1)
	}

	defer f.Close()

	hash := sha256.New()
	n, err := io.Copy(hash, f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Read file error: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("File %s has %d bytes\n", os.Args[1], n)
	fmt.Printf("%x\n", hash.Sum(nil))
}
