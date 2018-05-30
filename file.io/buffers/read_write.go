package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, _ := os.Open("test.dat")
	outputFile, _ := os.OpenFile("tested.dat", os.O_WRONLY|os.O_CREATE,
		0666)
	defer inputFile.Close()
	defer outputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			fmt.Println("EOF")
			break
		}
		outputString := inputString[2:5] + "\r\n"
		_, err := outputWriter.WriteString(outputString)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	outputWriter.Flush()
	fmt.Println("Conversion done")
}
