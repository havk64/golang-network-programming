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

func createHMAC(message, key []byte) []byte {

	hash := hmac.New(sha256.New, key)

	hash.Write(message)
	return hash.Sum(nil)
}

func checkMAC(message, messageMAC, key []byte) ([]byte, bool) {

	hash := hmac.New(sha256.New, key)

	hash.Write(message)
	expectedMAC := hash.Sum(nil)
	return expectedMAC, hmac.Equal(messageMAC, expectedMAC)
}
