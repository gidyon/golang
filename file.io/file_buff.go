package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, inputError := os.Open("input.dat")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	for {
		buf := make([]byte, 1024)
		n, err := inputReader.Read(buf)
		if err == io.EOF {
			break
		}
		if n == 0 {
			break
		}
		fmt.Printf("The input was: %s", string(buf[:n]))
	}
}
