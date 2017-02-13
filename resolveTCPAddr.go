package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <network-type> <address-with-port>\n", os.Args[0])
		os.Exit(1)
	}

	networkType := os.Args[1]
	addr := os.Args[2]

	TCPAddr, err := net.ResolveTCPAddr(networkType, addr)
	if err != nil {
		fmt.Println("Error ", err.Error())
	}
	fmt.Printf("TCPAddr struct: %#v\n", TCPAddr)
}
