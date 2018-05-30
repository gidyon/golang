package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("test.txt", os.O_RDWR, 0)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(data))
	_, err = f.WriteString("hello world")
	if err != nil {
		log.Fatalln(err)
	}
}
