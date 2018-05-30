/*
	Program empty_switch.go shows an example of usage of the empty interface in a type switch
	combined with a lambda function:
*/
package main

import "fmt"

type specialString string

var whatIsThis specialString = "hello"

func TypeSwitch(any interface{}) {
	testFunc := func() {
		switch v := any.(type) {
		case bool:
			fmt.Printf("any %v is a bool type", v)
		case int:
			fmt.Printf("any %v is an int type", v)
		case float32:
			fmt.Printf("any %v is a float32 type", v)
		case string:
			fmt.Printf("any %v is a string type", v)
		case specialString:
			fmt.Printf("any %v is a special String!", v)
		default:
			fmt.Println("unknown type!")
		}
	}
	testFunc()
}
func main() {
	TypeSwitch(whatIsThis)
}
