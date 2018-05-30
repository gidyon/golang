package main

import (
	"fmt"
)

type Stack []int

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	*s = (*s)[:len(*s)-1]
}

func (s *Stack) Push(num int) {
	*s = append(*s, num)
}

func (s *Stack) IsEmpty() bool {
	if len(*s) == 0 {
		return true
	}
	return false
}

func (s *Stack) String() string {
	return fmt.Sprint(s)
}

func main() {
	var s Stack = []int{}
	s.Push(12)
	s.Push(10)
	s.Push(15)
	s.Push(8)
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
}
