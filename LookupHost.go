package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <hostname>\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]

	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}
	len := len(addrs)

	fmt.Printf("Addresses found: {")
	for i, s := range addrs {
		fmt.Printf(s)
		if i == len-1 {
			fmt.Printf("}\n")
		} else {
			fmt.Printf(", ")
		}
	}

	cname, err := net.LookupCNAME(name)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}

	fmt.Printf("CNAME is: %s", cname)

	os.Exit(0)
}
