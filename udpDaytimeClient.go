// UDP Daytime client
//
// Usage example:
//   $ go run udpDaytimeClient.go localhost:1200
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
	service := os.Args[1]
	// Convert first argument(Args[1]) to binary address and fill a UDPAddr struct
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError("Address", err)
	// Optionally prints the IP address:(uncomment the line below)
	// fmt.Printf("IP: %s\n", udpaddr.IP)
	//
	// Create a UDP Socket and connect to address:port from the UDPAddr struct
	// and returns a file descriptor
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError("Connection", err)

	_, err = conn.Write([]byte("anything"))
	checkError("Write", err)
	// Create a buffer to read the response
	var buf [512]byte
	n, err := conn.Read(buf[0:]) // Read the response
	checkError("Read", err)
	// Prints the response(type []byte) as string
	fmt.Println(string(buf[0:n]))

	os.Exit(0)
}

func checkError(s string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", s, err.Error())
		os.Exit(1)
	}
}
