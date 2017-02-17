// Simplifyied verion of Daytime Client
// It uses Dial method which uses the "conn" interface and
// works both with TCP and UDP
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <host:port>", os.Args[0])
		os.Exit(1)
	}
	// The options for protocol are: tcp, tcp4(ipv4), tcp6(ipv6), upd, udp4(ipv4),
	// upd6(ipv6), ip, ip4(ipv4), ip6(ipv6).
	protocol := "tcp"
	conn, err := net.Dial(protocol, os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Connection: %s\n", err.Error())
		os.Exit(1)
	}
	counter := 0
	for { // The conn file descriptor is an 'io.Reader' so it implements the Read method
		buf := make([]byte, 64)
		if _, err := conn.Read(buf); err != nil {
			if err.Error() == "EOF" { // Exit loop when read is done
				break
			} // Otherwise report the read error
			fmt.Fprintf(os.Stderr, "Read: %s\n", err.Error())
		} // Prints the response(type []byte) as string
		fmt.Printf("%s", string(buf))
		counter++
	}
	fmt.Printf("\nCounter: %d\n", counter)
}
