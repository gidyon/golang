package main

import (
	"fmt"
	"log"
)

type Test int
type Inter interface {
	Demo()
}

func (t Test) Demo() {
	log.Println("Demo of ", t)
}
func main() {
	var a Inter
	a = Test(12)
	a.Demo()
	fmt.Println(a)
}
