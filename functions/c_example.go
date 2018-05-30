package main

import (
	"fmt"
)

func fn() {
	var x, k int
	fmt.Printf("Enter x =>")
	fmt.Scan(&x)
	k = x*x + 2*x + 5
	fmt.Printf("\nThe function value is %v", k)
}

func main() {
	fn()
}
