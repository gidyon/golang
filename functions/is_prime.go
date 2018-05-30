package main

import "fmt"

func isPrime(num int) bool {
	if num%2 == 0 && num != 2 {
		return false
	}
	for i := 3; i < (num+1)/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func isprime(n int, fn func (n, i int) int) func() bool {
	if n%2 == 0 {
		return func() {
			return false
		}
	}
	index := (n + 1) / 2
	return func() {
		fn(n,)
	}
}

chech(n, i int) int {
	if num%i == 0 {
		return false
	}
}
func isp(n int) bool {
	isPrime := num%2 == 0
	if isPrime && num != 2 {
		return false
	}
	if n%index == 0 {
		return false
	}
	isp()
}
func main() {
	fmt.Println(isPrime(101))
}
