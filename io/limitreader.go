package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	// Stops with EOF after reading n bytes
	lr := io.LimitReader(r, 10)

	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
}
