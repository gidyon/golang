package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}
	w, err := os.Create("output.txt")
	_, err = w.Write(data)
	if err != nil {
		log.Fatalln(err)
	}
}
