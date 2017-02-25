// Calculates the md5 checksum of a sequence of bytes and prints it in for parts
// of 4 bytes each(16 bytes total).
package main

import (
	"crypto/md5"
	"fmt"
)

func main() {

	hash := md5.New()
	bytes := []byte("Hello World!\n")

	hash.Write(bytes)
	hashSize := hash.Size()
	hashValue := hash.Sum(nil)
	// Print the full hash
	// fmt.Printf("%x\n", hashValue)
	splitted := splitHash(hashValue, hashSize)
	fmt.Printf("%s\n", splitted)
}

func splitHash(hash []byte, size int) string {

	var str string

	for n := 0; n < size; n += 4 {
		var val uint32
		val = uint32(hash[n])<<24 +
			uint32(hash[n+1])<<16 +
			uint32(hash[n+2])<<8 +
			uint32(hash[n+3])
		str += fmt.Sprintf("%x ", val)
	}
	return str
}
