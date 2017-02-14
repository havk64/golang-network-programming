// Make a http HEAD request to an address:port specified in the command line
// Example:
//   $ go run getHeadInfo.go www.golang.org:80
package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <host:port>", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	// Resolve IP and return a TCPAddr struct with IP address and Port
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError("Resolve error", err)
	// Create the socket, connect and return its file descriptor
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError("Socket error", err)
	// Send the HEAD request
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n" +
		"User-Agent: Golang Tcp client\r\n" +
		"\r\n"))
	checkError("Write error", err)
	// Read the response
	result, err := ioutil.ReadAll(conn)
	checkError("Read error", err)

	fmt.Println(string(result))
}

func checkError(s string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s", s, err.Error())
		os.Exit(1)
	}
}
