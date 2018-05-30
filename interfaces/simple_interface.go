package main

import "fmt"
type Simpler interface {
	Get() int
	Set(n int)
}

type Simple struct {
	value int
}

func (s Simple) Get() int {
	return s.value
}

func (s Simple) Set(n int) {
	s.value = n
}

type RSimple struct {
	value int
	min int
	max int
}

func (s RSimple) Get() int {
	return s.value
}

func (s RSimple) Set(n int) {
	s.value = n
}

func (s RSimple) SetMin() int {
	s.min = n
}

func (s RSimple) SetMax(n int) {
	s.max = n
}

func Smpl(s Simpler, n int) {
	switch s.(type) {
	case *Simple:
		fmt.Println("The type of s is Simple")
	case *RSimple:
		fmt.Println("The type of s is RSimple")
	}
	s.Set(n)
	fmt.Println(s.Get())
}

func f1(s Simpler) {
	switch s.(type) {
	case *Simple:
		fmt.Println("The type of s is Simple")
	case *RSimple:
		fmt.Println("The type of s is RSimple")
	}
}

func main() {
	var s Simple
	s.Set(10)
	fmt.Println(s)
	Smpl(s, 10)
}