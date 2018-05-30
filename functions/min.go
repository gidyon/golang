package main

import "fmt"  

func main() {
	x := Min(6,14,5,8,15,12,10)
	fmt.Printf("The Minimum is %d\n", x)
	arr := []int{7,9,3,5,1}
	x = Min(arr...)
	fmt.Printf("The Minimum in the array arr is %d\n", x)
}

func Min(a...int) int {
	if len(a) <= 1 {
		return a[0]
	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}