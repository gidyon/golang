package main

import (
	"fmt"
	"time"
)
func main(){
	var a = [10]int{1,2,3,4,5,6,7,8,9,10} 
	t1 := time.Now()
	sum1, sum2 := 0,0
	for i := 0; i < 20; i++ {
		sum1 += Sum(a[:])
	}
	t2 := time.Now()
	fmt.Println(sum1, t2.Sub(t1))
	t3 := time.Now()
	for i := 0; i < 20; i++ {
		sum2 += Pum(&a)
	}
	t4 := time.Now()
	fmt.Println(sum2, t4.Sub(t3))
}

func Sum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func Pum(a *[10]int) int {
	sum := 0
	for _, v := range *a {
		sum += v
	}
	return sum
}