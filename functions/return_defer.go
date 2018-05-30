package main

import (
	"fmt"
	"errors"
)

func main() {
	fmt.Println(f())
	val, err := tst()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(val, err)
}

func f() (ret int) {
	defer func(){
		ret++
	}()
	return 1
}

func tst() (res int, err error) {
	defer func(){
			err = errors.New("Bad")
		}()
	return 2, err
}