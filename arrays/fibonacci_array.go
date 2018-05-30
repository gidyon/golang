package main

import (
	"fmt"
)

const LIM int = 50

var fibs [LIM]uint64

func main() {
	// for i := 0; i < LIM; i++ {
	// 	fmt.Printf("%d\t", fibonacci(i))
	// }
	fmt.Print(fib(LIM))
}

func fibonacci(n int) (res uint64) {
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return
}

func fib(n int) (res []uint64) {
	for i := 0; i < n; i++ {
		fibonacci(i)
	}
	res = fibs[:]
	return
}
