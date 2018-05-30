package main

import "fmt"  

func main() {
	n := 7
	fmt.Printf("Is %d even %v\n", n, even(n))
	fmt.Printf("Is %d odd %v\n", 5, odd(5))
	fmt.Printf("Is %d even %v\n", 6, even(6))
	fmt.Printf("Is %d odd %v\n", 4, odd(4))
}

func even(n int) bool {
	if n == 0 {
		return true
	}
	return odd(n-1)
}

func odd(n int) bool {
	if n == 0 {
		return false
	}
	return even(n-1)
}
