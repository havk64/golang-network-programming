// Needs to be compiled with the file person.go
package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

func main() {
	person := Person{
		Name: Name{
			Family:   "de Oliveira",
			Personal: "Alexandro"},
		Email: []Email{
			Email{
				Kind:    "home",
				Address: "alexandro.deoliveira@icloud.com"},
			Email{
				Kind:    "school",
				Address: "alexandro.oliveira@holbertonschool.com"}}}

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
