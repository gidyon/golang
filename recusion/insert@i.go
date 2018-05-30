package main

import (
	"fmt"
)

func InsertAtIndexI(arr []int, item, K int) []int {
	N := len(arr)
	J := N
	N = N + 1
	for J >= K {
		arr[J+1] = arr[J]
		J = J - 1
	}
	arr[K] = item
	return arr
}

func main() {
	arr := []int{12, 34, 89, 11, 1, 90}
	fmt.Println(arr)
	fmt.Println("After Insertion at index 3, value 12, the array will be")
	fmt.Println(InsertAtIndexI(arr, 12, 3))
}
