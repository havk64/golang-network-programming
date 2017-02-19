// Daytime Server using asn1 serialization
package main

import (
	"encoding/asn1"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	// The options for protocol are: tcp, tcp4(ipv4), tcp6(ipv6), upd, udp4(ipv4),
	// upd6(ipv6), ip, ip4(ipv4), ip6(ipv6).
	protocol := "tcp"
	service := ":1200"

	listener, err := net.Listen(protocol, service)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Connection: %s", err.Error())
		os.Exit(1)
	}
	// Prints an init message
	fmt.Printf("Listening on %s...\n", listener.Addr())
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		daytime := time.Now()
		// Marshals the message before send it
		mdata, err := asn1.Marshal(daytime)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Marshall: %s\n", err.Error())
		}
		conn.Write(mdata)
		conn.Close()
	}
}
