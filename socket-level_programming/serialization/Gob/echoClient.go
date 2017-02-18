// Needs to be compiled with the file person.go
package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

func main() {
	person = Person{
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

	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}
	// The options for protocol are: tcp, tcp4(ipv4), tcp6(ipv6), upd, udp4(ipv4),
	// upd6(ipv6), ip, ip4(ipv4), ip6(ipv6).
	protocol := "tcp"
	service := os.Args[1]

	conn, err := net.Dial(protocol, service)
	if err != nil {
		fmt.Printf("Dial %s\n", err.Error())
		os.Exit(1)
	}

	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)

	for n := 0; n < 10; n++ {
		encoder.Encode(person)
		var newPerson Person
		decoder.Decode(&newPerson)
		fmt.Println(newPerson.String())
	}
}
