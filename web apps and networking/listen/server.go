package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		fmt.Println(conn.LocalAddr(), "------", conn.RemoteAddr())
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return
		}
		io.Copy(conn, conn)
		conn.Close()
	}
}
