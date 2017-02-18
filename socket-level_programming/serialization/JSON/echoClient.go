package main

import (
	"encoding/json"
	"fmt"
	"net"
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

// Implement fmt.Stringer interface
func (p Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family
	for _, v := range p.Email {
		s += "\n" + v.Kind + ": " + v.Address
	}
	return s
}

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

	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)

	for n := 0; n < 10; n++ {
		encoder.Encode(person)
		var newPerson Person
		decoder.Decode(&newPerson)
		fmt.Println(newPerson.String())
	}
}
