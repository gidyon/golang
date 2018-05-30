package main

import (
	"fmt"
)

func main() {
	fmt.Println(len(fibonacci3(19)))
}

// Store the fibonacci series in a slice starting from 0 to limit
func fibonacci(limit int) (series []int) {
	x, y := 0, 1
	series = append(series, x, y)
	for {
		x = x + y
		if x > limit {
			break
		}
		series = append(series, x)
		y = x + y
		if y > limit {
			break
		}
		series = append(series, y)
	}
	return
}

// store fibonacci series in a slice starting from low limit to high limit
func fibonacci2(start, second, limit int) (series []int) {
	series = append(series, start, second)
	for {
		start = start + second
		if start > limit {
			break
		}
		series = append(series, start)
		second = start + second
		if second > limit {
			break
		}
		series = append(series, second)
	}
	return
}

// fibonacci of n numbers
func fibonacci3(nums int) (series []int) {
	x, y := 0, 1
	series = append(series, x, y)
	z := nums
	if nums%2 == 1 {
		z = nums + 1
	}
	for i := 0; i < (z / 2); i++ {
		x = x + y
		y = x + y
		series = append(series, x, y)
	}
	if nums%2 == 1 {
		series = series[:len(series)-3]
		return
	}
	series = series[:len(series)-2]
	return
}
