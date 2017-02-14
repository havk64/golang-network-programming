// Daytime Client
//
// Usage example:
//   $ go run dayTimeClient.go time.nist.gov:13
package main

import (
	"fmt"
	"io/ioutil"
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

	response, err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Read error %s\n", err.Error())
	}
	fmt.Printf("%s", string(response))

	os.Exit(0)
}
