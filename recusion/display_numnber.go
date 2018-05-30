package main

import (
	"fmt"
)

func displayNumber(n int) {
	fmt.Println(n)
	if n > 0 {
		fmt.Println(n)
		displayNumber(n - 1)
		fmt.Println(n)
	} else {
		fmt.Println(n)
	}
}

func main() {
	displayNumber(3)
}
