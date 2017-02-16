// Daytime Client
//
// Usage example:
//   $ go run dayTimeClient.go time.nist.gov:13
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <host:port>\n", os.Args[0])
		os.Exit(1)
	}
	// Convert first argument(Args[1]) to binary address and fill a TCPAddr struct
	addr, err := net.ResolveTCPAddr("tcp", os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Address error, %s\n", err.Error())
		os.Exit(1)
	}
	// Optionally prints the IP address:(uncomment the line below)
	// fmt.Printf("IP: %s\n", addr.IP)
	//
	// Create a TCP Socket and connect to address:port from the TCPAddr struct
	// and returns a file descriptor
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Connection error: %s\n", err.Error())
		os.Exit(1)
	}
	// Create a little buffer to read the http response
	buf := make([]byte, 64)
	for { // The conn file descriptor is an 'io.Reader' so it implements the Read method
		_, err := conn.Read(buf)
		if err != nil {
			if err.Error() == "EOF" { // Exit loop when read is done
				break
			} // Otherwise report the read error and exit
			fmt.Fprintf(os.Stderr, "Read error: %s\n", err.Error())
			os.Exit(1)
		} // Prints the response(type []byte) as string
		fmt.Printf("%s", string(buf))
	}
}
