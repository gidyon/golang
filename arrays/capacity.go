package main

import "fmt"
func main(){
	s := make([]int, 0)
	fmt.Println("The capacity of s is: ",cap(s))
	s = append(s, 3)
	fmt.Println("The new capacity of s is: ",cap(s))
}