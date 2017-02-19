// Needs to be compiled with the file person.go
package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

func main() {
	// The 'person' variable is defined in the file person.go
	saveGob("person.gob", person)
}

func saveGob(fileName string, key interface{}) {

	outFile, err := os.Create(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File creation: %s\n", err.Error())
		os.Exit(1)
	}

	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Encode: %s\n", err.Error())
		os.Exit(1)
	}

	outFile.Close()
}
