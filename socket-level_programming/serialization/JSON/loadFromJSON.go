package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Person struct
type Person struct {
	Name  Name    `json:"name"`
	Email []Email `json:"emails"`
}

// Name struct
type Name struct {
	Family   string `json:"family"`
	Personal string `json:"personal"`
}

// Email struct
type Email struct {
	Kind    string `json:"kind"`
	Address string `json:"address"`
}

func main() {

	var person Person
	loadJSON("person.json", &person)
	fmt.Printf("%s\n", person.String())
}

func loadJSON(fileName string, key interface{}) {

	inFile, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Opening file: %s\n", err.Error())
		os.Exit(1)
	}

	decoder := json.NewDecoder(inFile)
	err = decoder.Decode(key)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Encode: %s\n", err.Error())
		os.Exit(1)
	}

	inFile.Close()
}

// Implement fmt.Stringer interface
func (p Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family
	for _, v := range p.Email {
		s += "\n" + v.Kind + ": " + v.Address
	}
	return s
}
