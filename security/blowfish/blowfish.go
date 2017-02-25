package main

import (
	"fmt"

	"golang.org/x/crypto/blowfish"
)

func main() {
	key := []byte("my key")
	message := []byte("hello\n\n\n")

	enc, _ := encrypt(message, key)

	msg, _ := decrypt(enc, key)
	fmt.Printf("%s\n", msg)
}

func encrypt(message, key []byte) ([]byte, error) {
	cipher, err := blowfish.NewCipher(key)
	if err != nil {
		return nil, err
	}

	enc := make([]byte, 512)

	cipher.Encrypt(enc, message)
	return enc, nil
}

func decrypt(message, key []byte) ([]byte, error) {
	cipher, err := blowfish.NewCipher(key)
	if err != nil {
		return nil, err
	}

	decrypt := make([]byte, 8)
	cipher.Decrypt(decrypt, message)
	return decrypt, nil
}
