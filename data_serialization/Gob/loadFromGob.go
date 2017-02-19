// Needs to be compiled with the file person.go
package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

func main() {

	var person Person
	loadGob("person.gob", &person)
	fmt.Printf("%s\n", person.String())
}

func loadGob(fileName string, key interface{}) {

	inFile, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Opening file: %s\n", err.Error())
		os.Exit(1)
	}

	decoder := gob.NewDecoder(inFile)
	err = decoder.Decode(key)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Encode: %s\n", err.Error())
		os.Exit(1)
	}

	inFile.Close()
}
