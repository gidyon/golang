package main

import (
	"fmt"
)

type Base struct {
	id int
}

func (b *Base) Id() int {
	return b.id
}

func (b *Base) SetId(id int) {
	b.id = id
}

type Person struct {
	Base
	FirstName string
	LastName string
}

type Employee struct {
	Person
	salary float64
}

func (s *Employee) Salary() float64 {
	return s.salary
}

func (s *Employee) SetSalary(amount float64) {
	s.salary = amount
}

func main(){
	emp := new(Employee)
	emp.SetSalary(1000.5)
	emp.FirstName = "Gideon"
	emp.LastName = "Kamau"
	emp.SetId(10)
	fmt.Println(emp.Id())
}
