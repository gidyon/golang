package main

import (
	"fmt"
	"time"
)

func Registrations() func(string) {
	total := 0
	females := 0 // Female Tickets
	males := 0   // Male Tickets
	return func(s string) {
		if total == 24 {
			fmt.Printf("Space filled: Total: %d, Males: %d, Females: %d\n", total, males, females)
			return
		}
		if s == "female" {
			total++
			females++
		}
		if s == "male" {
			total++
			males++
		}
		fmt.Printf("Total: %d/24, Males: %d, Females: %d\n", total, males, females)
	}
}

func main() {
	registrations := Registrations()
	go func() {
		for {
			registrations("male")
			time.Sleep(time.Millisecond * 2000)
		}
	}()
	go func() {
		for {
			registrations("female")
			time.Sleep(time.Millisecond * 2500)
		}

	}()
	var c string
	fmt.Scan(&c)
}
