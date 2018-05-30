package main

import "fmt"

func main() {
	n := 10
	fmt.Print(factorial(n))
}

func factorial(n int) (fact int) {
	if n == 0 {
		return 1
	}
	fact = n * factorial(n-1)
	return
}
