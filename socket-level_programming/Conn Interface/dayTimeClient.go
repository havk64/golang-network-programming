// Simplifyied verion of Daytime Server
// It uses Listen method which uses the "conn" interface and
// works with TCP and UDP
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	service := ":1200"

	listener, err := net.Listen("tcp", service)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Connection: %s", err.Error())
		os.Exit(1)
	}
	// Prints an init message
	fmt.Printf("Listening on localhost%s...\n", service)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}
}
