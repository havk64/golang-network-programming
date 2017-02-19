package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
)

func main() {

	eightBitData := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	var bb bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &bb)
	encoder.Write(eightBitData)
	encoder.Close()
	fmt.Printf("bb: %s\n", bb.String())

	dbuf := make([]byte, 12)
	decoder := base64.NewDecoder(base64.StdEncoding, &bb)
	decoder.Read(dbuf)
	fmt.Printf("%v\n", dbuf)
}
