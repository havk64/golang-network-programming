package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	tick := time.NewTicker(time.Second)
	go func() {
		for _ = range tick.C {
			fmt.Printf(".")
		}
	}()

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <host>\n", os.Args[0])
		os.Exit(1)
	}

	raddr, err := net.ResolveIPAddr("ip4", os.Args[1])
	checkError("Resolve remote", err)

	laddr, err := net.ResolveIPAddr("ip4", "127.0.0.1")
	checkError("Resolve local", err)

	conn, err := net.DialIP("ip4:icmp", laddr, raddr)
	checkError("DialIP", err)

	var msg [512]byte
	msg[0] = 8  // echo
	msg[1] = 0  // code 0
	msg[2] = 0  // checksum, fix later
	msg[3] = 0  // checksum, fix later
	msg[4] = 0  // identifier[0]
	msg[5] = 13 // identifier[1]
	msg[6] = 0  // sequence[0]
	msg[7] = 37 // sequence[1]
	len := 8

	check := checkSum(msg[0:len])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)

	_, err = conn.Write(msg[0:len])
	checkError("Write", err)

	_, err = conn.Read(msg[0:])
	checkError("Read", err)

	tick.Stop()
	fmt.Println("\nGot response")
	if msg[5] == 13 {
		fmt.Println("identifier matches")
	}
	if msg[7] == 37 {
		fmt.Println("Sequence matches")
	}

	os.Exit(0)
}

func checkSum(msg []byte) uint16 {
	sum := 0

	// assume even for now
	for n := 1; n < len(msg)-1; n += 2 {
		sum += int(msg[n])*256 + int(msg[n+1])
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	var answer = uint16(^sum)
	return answer
}

func checkError(s string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", s, err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	var result bytes.Buffer
	buf := make([]byte, 512)
	for {
		_, err := conn.Read(buf)
		result.Write(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
