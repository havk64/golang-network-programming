package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	// The options for protocol are: tcp, tcp4(ipv4), tcp6(ipv6), upd, udp4(ipv4),
	// upd6(ipv6), ip, ip4(ipv4), ip6(ipv6).
	protocol := "tcp"
	service := ":1200"

	listener, err := net.Listen(protocol, service)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Connection: %s\n", err.Error())
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 512)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			return
		}
		_, err2 := conn.Write(buf)
		if err2 != nil {
			return
		}
	}
}
