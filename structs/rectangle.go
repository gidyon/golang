package main

import (
	"fmt"
)

type Rectangle struct {
	length, width int
}

func (r *Rectangle) Area() int {
	return r.length * r.width
}

func (r *Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

func main() {
	rect := &Rectangle{10, 20}
	fmt.Println(rect.Area(), rect.Perimeter())
}
