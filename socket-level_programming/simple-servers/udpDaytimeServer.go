package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	service := ":1200"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError("Address", err)

	conn, err := net.ListenUDP("udp", udpAddr)
	checkError("Connection", err)
	fmt.Printf("Listening on localhost:%d\n", udpAddr.Port)

	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {

	var buf [512]byte

	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}

	daytime := time.Now().String()

	conn.WriteToUDP([]byte(daytime), addr)
}

func checkError(s string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", s, err.Error())
		os.Exit(1)
	}
}
