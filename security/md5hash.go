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
	// fmt.Printf("hashvalue: %x\n", hashValue)
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

	str += "\n"
	return str
}
