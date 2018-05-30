package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type MyLocalAddress struct {
	protocol string
	network  string
}

func (addr *MyLocalAddress) String() string {
	return addr.network
}
func (addr *MyLocalAddress) Network() string {
	return addr.protocol
}
func main() {
	// open connection:
	ipAddr, _ := net.ResolveIPAddr("udp", "193.0.2.1/8")
	fmt.Println(ipAddr.String())
	conn, err := net.Dial("tcp", "localhost:5000")
	fmt.Println(conn.LocalAddr(), "------", conn.RemoteAddr())
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\r\n")
	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimmedClient + " says: " +
			trimmedInput + "#end"))
	}
}
