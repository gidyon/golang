package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var inputReader *bufio.Reader
var input string
var err error

func auth(attempts int8) {
	if attempts >= 3 {
		return
	}
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Enter UserName: ")
	input, err = inputReader.ReadString('\n')
	if err == nil {
		if input != "gideon" {
			fmt.Println("Incorrect UserName!")
			auth(attempts + 1)
			return
		}
		time.Sleep(time.Second)
		fmt.Println("Enter Password: ")
		input, err = inputReader.ReadString('\n')
		if input != "hakty11" {
			fmt.Println("Incorrect Password!")
			attempts++
			auth(attempts + 1)
			return
		}
	}
	return
}

func main() {
	auth(0)
}
