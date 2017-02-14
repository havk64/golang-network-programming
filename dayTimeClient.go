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

	addr, err := net.ResolveTCPAddr("tcp", os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Address error, %s\n", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Socket error %s\n", err.Error())
		os.Exit(1)
	}

	buf := make([]byte, 64)
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		fmt.Printf("%s", string(buf[0:n]))
	}
}
