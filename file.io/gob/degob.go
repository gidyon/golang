package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}
type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	vc := new(VCard)
	file, _ := os.OpenFile("vcard.gob", os.O_RDONLY, 0)
	defer file.Close()
	enc := gob.NewDecoder(file)
	err := enc.Decode(vc)
	if err != nil {
		log.Println("Error in encoding gob")
	}
	fmt.Println(vc)
}
