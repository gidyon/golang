package main

import "fmt"  

func main() {
	fib := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(fib())
	}
}

func fib(n int) (res int) {
	if n <= 1 {
		res = 1
	}else{
		res = fib(n-1) + fib(n-2)
	}	
	return
}


func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first + second
		return ret
	}
}