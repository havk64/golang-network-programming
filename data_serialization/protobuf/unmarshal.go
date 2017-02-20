package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
)

func writePerson(w io.Writer, p *Person) {
	fmt.Fprintln(w, "Person ID:", p.Id)
	fmt.Fprintf(w, "  Name: %s %s\n", p.Name.Personal, p.Name.Family)

	for _, v := range p.Emails {
		switch v.Kind {
		case Person_SCHOOL:
			fmt.Fprint(w, "  School email: ")
		case Person_HOME:
			fmt.Fprint(w, "  Home email: ")
		case Person_WORK:
			fmt.Fprint(w, "  Work email: ")
		}
		fmt.Fprintln(w, v.Address)
	}
}

func listPeople(w io.Writer, book *AddressBook) {
	for _, p := range book.People {
		writePerson(w, p)
	}
}

func main() {

	if len(os.Args) != 2 {
		log.Fatalf("Usage:  %s ADDRESS_BOOK_FILE\n", os.Args[0])
	}
	fname := os.Args[1]

	// Read the existing address book.
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	book := &AddressBook{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	listPeople(os.Stdout, book)
}
