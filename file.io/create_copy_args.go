package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	src := os.Args[1]
	dst := os.Args[2]
	if len(os.Args) < 3 {
		log.Fatalln("Usage: cc <src> <dst>")
	}
	f, err := os.Open(src)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}
	w, err := os.Create(dst)
	_, err = w.WriteString(string(data))
	if err != nil {
		log.Fatalln(err)
	}
}
