// Make a http HEAD request to an host:port specified in the command line
// Example:
//   $ go run getHeadInfo.go www.golang.org:80
package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	start := time.Now()
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <host:port>", os.Args[0])
		os.Exit(1)
	}
	// The options for protocol are: tcp, tcp4(ipv4), tcp6(ipv6), upd, udp4(ipv4),
	// upd6(ipv6), ip, ip4(ipv4), ip6(ipv6).
	protocol := "tcp"
	// Create the socket, connect and return its file descriptor
	// The Dial method works for create for TCP or UPD sockets
	conn, err := net.Dial(protocol, os.Args[1])
	checkError("Socket error", err)
	// Send the HEAD request
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n" +
		"User-Agent: Golang Tcp client\r\n" +
		"\r\n"))
	checkError("Write error", err)
	// Read the response using a buffer
	var response bytes.Buffer
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		response.Write(buf)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			checkError("Read error", err)
		}
	}
	// Writes directly to stdout
	response.WriteTo(os.Stdout)
	fmt.Printf("Time: %s\n", time.Since(start))
}

func checkError(s string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", s, err.Error())
		os.Exit(1)
	}
}
