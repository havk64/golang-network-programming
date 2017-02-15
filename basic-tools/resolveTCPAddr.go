// Usage examples:
//  $ go run resolveTCPAddr.go tcp 192.168.1.1:22
//  $ go run resolveTCPAddr.go tcp [www.golang.org]:53
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <network-type> <address:port>\n", os.Args[0])
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
