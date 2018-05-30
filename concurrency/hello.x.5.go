package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go func() {
			time.Sleep(time.Second * 1)
			fmt.Println("hello world")
		}()
	}

	var input string
	fmt.Scanln(&input)
	if input == "q" {
		os.Exit(1)
	} else {
		main()
	}
}
