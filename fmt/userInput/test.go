package main

import (
	"fmt"
	"time"
)

func sq() {
	var num float64
	fmt.Print("Enter A Number to find the Square: ")
	fmt.Scanln(&num)
	if num == 0 {
		fmt.Println("Please Enter a Valid Number!!")
		time.Sleep(time.Second)
		sq()
		return
	}
	fmt.Printf("The Square of %v is: %.4f\n\n", num, num*num)
	fmt.Print("Press Enter to exit")
	fmt.Scanln(&num)
	fmt.Print("Exiting...")
	return
}

func main() {
	sq()
}
