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

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError("Resolve error", err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError("Socket error", err)
	fmt.Printf("conn: %#v\n", conn)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError("Write error", err)

	result, err := ioutil.ReadAll(conn)
	checkError("Read error", err)

	fmt.Println(string(result))

	os.Exit(0)
}

func checkError(s string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s", s, err.Error())
		os.Exit(1)
	}
}
