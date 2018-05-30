package main

import (
	"fmt"
	"time"
	"unsafe"
)

type VCard struct {
	Name      string
	DOB       time.Time
	Photo     string
	Addresses *[]Address
}

type Address struct {
	Estate string
	Line   int
}

func main() {
	myVcard := VCard{
		Name:  "Gideon",
		DOB:   time.Now(),
		Photo: "/photos/gideon.jpg",
	}
	fmt.Println(myVcard, unsafe.Sizeof(Address{}))
}
