package main

import (
	"errors"
	"fmt"
	"math"
)

var errNotFound error = errors.New("Not found error")

func Sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("math - square root of negative numbers")
	}
	return math.Sqrt(num), nil
}

func main() {
	//fmt.Printf("error: %v", errNotFound)

	f, err := Sqrt(-1)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Println(f)
}

// error: Not found error
