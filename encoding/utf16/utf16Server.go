package main

import (
	"fmt"
	"net"
	"os"
	"unicode/utf16"
)

// BOM (Byte Order Markder) contant defines the endianess
const BOM = '\ufffe'

func main() {

	service := ":1210"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Resolve: %s\n", err.Error())
		os.Exit(1)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Socket: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Listening on %s\n", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		str := "j'ai arrêté"
		shorts := utf16.Encode([]rune(str))
		writeShorts(conn, shorts)

		conn.Close()
	}
}

func writeShorts(conn net.Conn, shorts []uint16) {
	var bytes [2]byte

	// send the BOM as first two bytes
	bytes[0] = BOM >> 8  // = 255 or 0xff
	bytes[1] = BOM & 255 // = 254 or 0xfe
	_, err := conn.Write(bytes[0:])
	if err != nil {
		return
	}

	for _, v := range shorts {
		bytes[0] = byte(v >> 8)
		bytes[1] = byte(v & 255)

		_, err = conn.Write(bytes[0:])
		if err != nil {
			return
		}
	}
}
