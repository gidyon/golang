package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("-----------------------------------------------------------------")
	inputReader := bufio.NewReader(os.Stdin)
	for {
		number1, number2 := 0, 0
		input, err := inputReader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		if input == "q\r\n" {
			fmt.Println("Exiting our calculator program....")
			break
		}
		number1, err = strconv.Atoi(strings.TrimSuffix(input, "\r\n"))
		if err != nil {
			log.Fatalln(err)
		}

		input, err = inputReader.ReadString('\n')
		number2, err = strconv.Atoi(strings.TrimSuffix(input, "\r\n"))
		if err != nil {
			log.Fatalln(err)
		}

		input, err = inputReader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		switch input {
		case "+\r\n":
			res := number1 + number2
			fmt.Printf("Addition of %d + %d = %d\n\n", number1, number2, res)
			continue
		case "-\r\n":
			res := number1 - number2
			fmt.Printf("Subtraction of %d - %d = %d\n\n", number1, number2, res)
			continue
		case "*\r\n":
			res := number1 * number2
			fmt.Printf("Multiplication of %d * %d = %d\n\n", number1, number2, res)
			continue
		case "/\r\n":
			res := number1 / number2
			fmt.Printf("Division of %d / %d = %d\n", number1, number2, res)
			continue
		}
	}

}
