package main

import (
	"fmt"
)

func mistery(n uint) {
	fmt.Println(n)
	if n > 0 {
		mistery(n - 1)
		fmt.Println(n)
	}
}

func main() {
	mistery(5)
}
