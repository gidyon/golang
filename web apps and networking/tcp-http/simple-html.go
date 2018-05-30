package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
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
	i, url, method := 0, "", ""
	headers := map[string]string{}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			header := strings.Fields(ln)
			method = header[0]
			url = header[1]
			fmt.Println("METHOD: ", method)
			fmt.Println("URL: ", url)
		} else {
			if ln == "" {
				break
			}
			subs := strings.SplitN(ln, ": ", 2)
			headers[subs[0]] = subs[1]
		}
		i++
	}
	if method == "POST" || method == "PUT" {
		amt, _ := strconv.Atoi(headers["Content-Length"])
		buf := make([]byte, amt)
		io.ReadFull(conn, buf)

		fmt.Println("BODY: ", string(buf))
	}
	body := `<html>
			<head></head>
			<body>
				<form method="POST">
					<input type="text" name="key" value=""></input>
					<input type="submit" value="submit"></input>
				<form>
			</body>
			</html>`
	conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	conn.Write([]byte(body))
}
