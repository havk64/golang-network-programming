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
	// The options for protocol are: tcp, tcp4(ipv4), tcp6(ipv6), upd, udp4(ipv4),
	// upd6(ipv6), ip, ip4(ipv4), ip6(ipv6).
	protocol := "tcp"
	service := ":1200"

	listener, err := net.Listen(protocol, service)
	if err != nil {
		fmt.Printf("Listening %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Listening on %s\n", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Accept error: %s\n", err.Error())
		}

		encoder := json.NewEncoder(conn)
		decoder := json.NewDecoder(conn)

		for n := 0; n < 10; n++ {
			var person Person
			decoder.Decode(&person)
			fmt.Println(person.String())
			encoder.Encode(person)
		}
		conn.Close()
	}
}
