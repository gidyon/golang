package main

import (
	"fmt"
)

func main() {
	var dataSlice []string = []string{"hello", "cow", "boy", "or", "world"}
	var interfaceSlice []interface{} = make([]interface{}, len(dataSlice))
	for i, v := range dataSlice {
		interfaceSlice[i] = v
	}
	fmt.Println(dataSlice, interfaceSlice)
}
