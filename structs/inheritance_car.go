package main

import "fmt"
type Engine interface {
	Start()
	Stop()
}

type Car struct {
	wheelCount int
}

func (c *Car) numberOfWheels() int {
	return c.wheelCount
}

func (*Car) Start() {
	fmt.Println("Starting the Engine")
}

func (*Car) Stop() {
	fmt.Println("Stopping the Engine")
}
 
type Mercedes struct {
	Car
}

func (m *Mercedes) sayHiToMerkel() {
	m.Start()
	fmt.Println("Hello Merkel")
	m.Stop()
}

func main() {
	m := Mercedes{Car{wheelCount: 4}}
	fmt.Println(m)
	m.sayHiToMerkel()
	fmt.Println(m.numberOfWheels())
}