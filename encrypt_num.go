package main

import (
	"fmt"
	"time"
)
func main() {
	c := time.Now()
	y := time.Millisecond*1000
	for i := 0; i < int(y); i++ {
		x := i+1
		x++
	}
	d := time.Now()
	fmt.Println(1000/(y/(d.Sub(c)*1000)))
}