package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go suck(ch1)
	ch1 <- 10
	time.Sleep(1e9)
}

func suck(ch chan int) {
	fmt.Println(<-ch)
}
