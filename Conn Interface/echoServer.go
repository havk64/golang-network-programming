package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	service := ":1200"
	listener, err := net.Listen("tcp", service)
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
