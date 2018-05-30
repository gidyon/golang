package main

import "fmt"

func swap(arr []int) []int {
	arrLength := len(arr)
	var pivot int
	if arrLength%2 == 0 {
		pivot = arrLength / 2
	} else {
		pivot = (arrLength - 1) / 2
	}
	for i := 0; i < arrLength; i++ {
		if i == pivot {
			break
		}
		temp := arr[i]
		arr[i] = arr[(arrLength-1)-i]
		arr[(arrLength-1)-i] = temp
	}
	return arr
}

func main() {
	arr := []int{1, 2, 6, 7, 9, 12, 90}
	fmt.Println(swap(arr))
}
