// Needs to be compiled with the file person.go
package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

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

		encoder := gob.NewEncoder(conn)
		decoder := gob.NewDecoder(conn)

		for n := 0; n < 10; n++ {
			var person Person
			decoder.Decode(&person)
			fmt.Println(person.String())
			encoder.Encode(person)
		}
		conn.Close()
	}
}
