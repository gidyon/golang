package main

import "fmt"

func main() {
	n := 10
	print(n)
}

func print(n int) {
	if n < 1 {
		return
	}
	fmt.Println(n)
	print(n - 1)
}
