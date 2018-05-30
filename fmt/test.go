package main

import (
	"fmt"
)

func test(a string, b int) (strs []string, err error) {
	for i := 0; i < b; i++ {
		strs = append(strs, a)
	}
	err = nil
	return
}
func ret(err error) (interface{}, error) {
	if err == nil {
		return nil, err
	}
	return nil, nil
}
func main() {
	res, err := test("hello", 5)
	defer ret(err)
	fmt.Print(res)
}
