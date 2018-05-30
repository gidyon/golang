package main

import "fmt"

func main() {
	p := func () {
		fmt.Println("Hello World")
	}
	p()
	fmt.Printf("Type of %v is %T", p, p)
}

