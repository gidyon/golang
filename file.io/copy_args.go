package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("Usage: cc <src> <dst>")
	}
	src := os.Args[1]
	dst := os.Args[2]
	f, err := os.Open(src)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	w, err := os.Create(dst)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = io.Copy(w, f)
	if err != nil {
		log.Fatalln(err)
	}
}
