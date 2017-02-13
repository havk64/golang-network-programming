package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <dotted-ip-addr>\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]

	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	}
	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()
	fmt.Printf("Address is: %v\n", addr.String())
	fmt.Printf("Default mask length is: %d\n", bits)
	fmt.Printf("Leading ones count is: %d\n", ones)
	fmt.Printf("Mask is (hex): %v\n", mask.String())
	fmt.Printf("Network is: %v\n", network.String())
	os.Exit(0)
}
