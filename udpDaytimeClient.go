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

	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError("Address", err)

	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError("Connection", err)

	_, err = conn.Write([]byte("anything"))
	checkError("Write", err)

	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError("Read", err)

	fmt.Println(string(buf[0:n]))

	os.Exit(0)
}

func checkError(s string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", s, err.Error())
		os.Exit(1)
	}
}
