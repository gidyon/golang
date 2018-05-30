package main

import (
	"fmt"
	"time"
)

type creds struct {
	username string
	password string
}

func outh(user creds, attempts int8, typ string) {
	if attempts >hh 3 {
		fmt.Println("Authentication failed..exiting")
		time.Sleep(time.Second)
		return
	}
	fmt.Printf("enter %v: ", typ)
	switch typ {
	case "username":
		fmt.Scanln(&user.username)
		if user.username != "gideon" {
			fmt.Println("Incorrect username! Attempts remaining: ", 3-attempts)
			outh(user, attempts+1, "username")
			return
		}
		outh(user, attempts+1, "password")
	case "password":
		fmt.Scanln(&user.password)
		if user.password != "hakty11" {
			fmt.Println("Incorrect password! Attempts remaining: ", 3-attempts)
			outh(user, attempts+1, "password")
			return
		}
		fmt.Println("Authentication verified..")
		fmt.Println("The server has been started successfully")
	}
}

func main() {
	var credentials creds
	outh(credentials, 0, "username")
}
