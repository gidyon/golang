package main

import (
	"fmt"
	"net"
	"strings"
)

var clientLists []string

func main() {
	fmt.Println("Starting the server ...")
	// create listener:
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return // terminate program
	}
	// listen and accept connections from clients:
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return // terminate program
		}
		go doServerStuff(conn)
	}
}
func isUserActive(name string) byte {
	for _, v := range clientLists {
		if v == name {
			return 1
		}
	}
	return 0
}

func addUser(name string) {
	clientLists = append(clientLists, name)
}

func isUserInList(name string) bool {
	for _, v := range clientLists {
		if v == name {
			return true
		}
	}
	return false
}
func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		msg := string(buf)[:strings.Index(string(buf), "#end")]
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return // terminate program
		}
		name := msg[:strings.Index(msg, " says")]
		if !isUserInList(name) {
			addUser(name)
		}
		if strings.HasSuffix(strings.Trim(msg, " "), " SH") {
			conn.Close()
			return
		}
		if strings.HasSuffix(msg, " WHO") {
			fmt.Println("This is the client list: 1=active, 0=inactive")
			for _, v := range clientLists {
				fmt.Printf("User %v is %v\n", v, isUserActive(v))
			}
			return
		}
		fmt.Printf("Received data: %v\n", msg)
	}
}
