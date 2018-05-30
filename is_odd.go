package main

import "fmt"

func IsOdd(n int) bool {
	if n == 0 {
		return false
	}
	return IsNotOdd(n - 1)
}

func IsNotOdd(n int) bool {
	if n == 0 {
		return true
	}
	return IsOdd(n - 1)
}

func Odd(n int) bool {
	if n == 0 {
		return false
	}
	return Odd(n - 1)
}
func main() {
	fmt.Println(IsOdd(4))
	fmt.Println(Odd(5))
}
