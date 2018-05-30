package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	ls, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}
	defer ls.Close()
	for {
		conn, err := ls.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			method := strings.Fields(ln)[0]
			fmt.Println("METHOD: ", method)
		} else {
			if ln == "" {
				break
			}
		}
		i++
	}
	body := "Hello world\n"
	conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	io.WriteString(conn, "\r\n")
	conn.Write([]byte(body))
}
