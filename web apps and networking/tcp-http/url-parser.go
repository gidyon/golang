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
	i, url := 0, ""
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			header := strings.Fields(ln)
			method := header[0]
			url = header[1]
			fmt.Println("METHOD: ", method)
			fmt.Println("URL: ", url)
		} else {
			if ln == "" {
				break
			}
		}
		i++
	}
	conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(url))
	io.WriteString(conn, "\r\n")
	conn.Write([]byte(url + "\n"))
}
