package main

import (
	"fmt"
)

var (
	msc    = 0
	fsc    = 0
	total  = 0
	status = ""
)

func main() {
	for total < 5 {
		fmt.Println("enter status")
		fmt.Scanln(&status)
		if status == "female" {
			fsc++
			total++
			fmt.Printf("Total: %d, Males: %d, Females: %d", total, msc, fsc)
		} else if status == "male" {
			msc++
			total++
			fmt.Printf("Total: %d, Males: %d, Females: %d\n", total, msc, fsc)
		} else {
			fmt.Println("Invalid status")
		}
	}
	fmt.Println("\nTotal exceeded")
}
