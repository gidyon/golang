package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	chars, words, lines := 0, 0, 0
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		if input == "S\r\n" {
			fmt.Printf("chars: %d, words: %d, lines: %d\n", chars, words, lines)
			break
		}
		lines++
		words += len(strings.Fields(input))
		for range input {
			chars++
		}
	}
}
