package main

import "fmt"

func main() {
	ch1 := make(chan int)
	go pump(ch1)       // pump hangs
	fmt.Println(<-ch1) // prints only 0
}

// Also called generator ~ because it supply values to a channel
func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}
