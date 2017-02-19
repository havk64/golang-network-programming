package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatalf("Usage:  %s ADDRESS_BOOK_FILE\n", os.Args[0])
	}
	fname := os.Args[1]
	//Define a person using struct literal
	person := &Person{
		Id: 1234,
		Name: &Person_Name{
			Family:   "de Oliveira",
			Personal: "Alexandro",
		},
		Emails: []*Person_Email{
			{Kind: Person_HOME,
				Address: "alexandro.deoliveira@icloud.com"},
			{Kind: Person_SCHOOL,
				Address: "alexandro.oliveira@holbertonschool.com"},
		},
	}
	// Add person to the address book
	book := &AddressBook{People: []*Person{person}}
	// Marshalling the address book
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	} // Writing the result to the file specified as first parameter in command line
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	} else {
		fmt.Printf("File %s created successfully\n", os.Args[1])
	}
}
