package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
)

type User struct {
	Name   string
	Id     int
	UName  string
	Output chan Message
}

func (u *User) SetUName() {
	u.UName = fmt.Sprintf("%v%v", u.Name, u.Id)
}

type Message struct {
	*User
	Text string
}

type ChatServer struct {
	Users map[string]User
	Join  chan User
	Leave chan User
	Input chan Message
}

func (cs *ChatServer) Run() {
	for {
		select {
		case user := <-cs.Join:
			cs.Users[user.UName] = user
			go func() {
				cs.Input <- Message{
					&User{Name: "SYSTEM"},
					fmt.Sprintf("%s joined", user.Name),
				}
			}()
		case user := <-cs.Leave:
			delete(cs.Users, user.UName)
			go func() {
				cs.Input <- Message{
					&User{Name: "SYSTEM"},
					fmt.Sprintf("%s left", user.Name),
				}
			}()
		case msg := <-cs.Input:
			for _, user := range cs.Users {
				select {
				case user.Output <- msg:
				default:
				}
			}
		}
	}
}
func handleConn(chatServer *ChatServer, conn net.Conn) {
	defer conn.Close()
	io.WriteString(conn, "Enter your UserName: ")
	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	user := User{
		Name:   scanner.Text(),
		Id:     rand.Int(),
		Output: make(chan Message, 10),
	}
	user.SetUName()
	chatServer.Join <- user
	defer func() {
		chatServer.Leave <- user
	}()

	// Read from conn
	go func() {
		for scanner.Scan() {
			ln := scanner.Text()
			chatServer.Input <- Message{&user, ln}
		}
	}()

	// Write to conn
	for msg := range user.Output {
		_, err := io.WriteString(conn, msg.Name+": "+msg.Text+"\n")
		if err != nil {
			break
		}
	}
}

func main() {
	addr, err := net.LookupHost("127.0.0.1")
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range addr {
		fmt.Println(v)
	}
	ls, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err)
	}
	defer ls.Close()

	chatServer := &ChatServer{
		Users: make(map[string]User, 0),
		Join:  make(chan User, 0),
		Leave: make(chan User, 0),
		Input: make(chan Message, 0),
	}
	go chatServer.Run()
	for {
		conn, err := ls.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handleConn(chatServer, conn)
	}
}
