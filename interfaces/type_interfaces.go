package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}
type Circle struct {
	radius float32
}
type Shaper interface {
	Area() float32
}

func main() {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5
	areaIntf = sq1
	// Is Square the type of areaIntf ?
	// Shadowing
	if areaIntf, ok := areaIntf.(*Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", areaIntf)
	}
	sq2 := new(Circle)
	sq2.radius = 5
	areaIntf = sq2
	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}
}
func (sq *Square) Area() float32 {
	return sq.side * sq.side
}
func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}
