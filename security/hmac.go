// This program implements a Keyed-Hash Message Authentication Code (HMAC).
// An HMAC is a cryptographic hash that uses a key to sign a message.
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func main() {
	message := []byte("Hello World!")
	key := []byte("My $3cr3t k3y")
	messageMAC := createHMAC(message, key)

	if checksum, ok := checkMAC(message, messageMAC, key); ok {
		fmt.Printf("Ok!\nMessage: %s\nChecksum: %x\n", string(message), checksum)
	} else {
		fmt.Println("Checksum error!")
	}
}

// createHMAC function return a SHA256 hash using a message and a key.
func createHMAC(message, key []byte) []byte {

	hash := hmac.New(sha256.New, key)

	hash.Write(message)
	return hash.Sum(nil)
}

// checkMAC function verifies the hash by recomputing it using the same key.
func checkMAC(message, messageMAC, key []byte) ([]byte, bool) {

	hash := hmac.New(sha256.New, key)

	hash.Write(message)
	expectedMAC := hash.Sum(nil)
	return expectedMAC, hmac.Equal(messageMAC, expectedMAC)
}
