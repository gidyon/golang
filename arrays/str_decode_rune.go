package main

import (
	"bufio"
	"os"
	"fmt"
	"unicode/utf8"
)

var arraysOfStrings [][]string

var arrayOfStrings []string

func scanRanges(str ...[]string) [][][]rune {			
	minimum := func(a []rune) (min rune) {
			if len(a) <= 1 {
				min = a[0]
				return
			}
			min = a[0]
			for _, v := range a {
				if v < min {
					min = v
				}
			}
			return
		}
	maximum := func(a []rune) (max rune) {
			if len(a) <= 1 {
				max = a[0]
				return
			}
			max = a[0]
			for _, v := range a {
				if v > max {
					max = v
				}
			}
			return
		}
	returnedArray := make([][][]rune, 0)
	for _, v := range str {
		arrayOfStrings := make([][]rune, 0)
		for _, strs := range v {
			arrayChars := make([]rune, 0)
			for len(strs) > 0 {
				r, size := utf8.DecodeRuneInString(strs)				
				arrayChars = append(arrayChars, r)
				strs = strs[size:]
			}
			charCodeMaxMin := []rune{minimum(arrayChars), maximum(arrayChars)}
			arrayOfStrings = append(arrayOfStrings, charCodeMaxMin)
		}
		returnedArray = append(returnedArray, arrayOfStrings)
	}
	return returnedArray
}

func readString2(currArray, currString int8) {	
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Printf("Array %d - Enter string %d: ", currArray, currString)
	str, _ := inputReader.ReadString('\n')
	if str == "exit" {		
		fmt.Printf("(Strings in array %d are %v)\n\n\n", currArray, arrayOfStrings)
		fmt.Printf("Applying the Algorithm....\n\n")
		result := scanRanges(arraysOfStrings...)
		fmt.Printf("Results is\n %v", result)
		os.Exit(1)
	}
	if str == "0n" {
		fmt.Printf("(Strings in array %d are %v)\n\n", currArray, arrayOfStrings)
		arraysOfStrings = append(arraysOfStrings, arrayOfStrings)
		currArray++;currString=1;arrayOfStrings = []string{}
		inputReader.Reset(os.Stdin)
		readString(currArray, currString)
	}	
	arrayOfStrings = append(arrayOfStrings, str)
	currString++
	readString(currArray, currString)
	return
}
func readString(currArray, currString int8) {
	var str string	
	fmt.Printf("Array %d - Enter string %d: ", currArray, currString)
	fmt.Scanln(&str)
	if str == "exit" {		
		fmt.Printf("(Strings in array %d are %v)\n\n\n", currArray, arrayOfStrings)
		fmt.Printf("Applying the Algorithm....\n\n")
		result := scanRanges(arraysOfStrings...)
		fmt.Printf("Results is\n %v", result)
		os.Exit(1)
	}
	if str == "0n" {
		fmt.Printf("(Strings in array %d are %v)\n\n", currArray, arrayOfStrings)
		arraysOfStrings = append(arraysOfStrings, arrayOfStrings)
		currArray++;currString=1;arrayOfStrings = []string{}
		readString(currArray, currString)
	}	
	arrayOfStrings = append(arrayOfStrings, str)
	currString++
	readString(currArray, currString)
	return
}
func readString3(currArray, currString int8) {
	// An artificial input source.
	scanner := bufio.NewScanner(os.Stdin)
	// Set the split function for the scanning operation.
	fmt.Printf("Array %d - Enter string %d: ", currArray, currString)
	scanner.Split(bufio.ScanWords)
	// Count the words.
	// count := 0
	// for scanner.Scan() {
	// 	count++
	// 	fmt.Println(scanner.Text())
	// }
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	// fmt.Printf("%d\n", count)
	str := scanner.Text()
	if str == "exit" {		
		fmt.Printf("(Strings in array %d are %v)\n\n\n", currArray, arrayOfStrings)
		fmt.Printf("Applying the Algorithm....\n\n")
		result := scanRanges(arraysOfStrings...)
		fmt.Printf("Results is\n %v", result)
		os.Exit(1)
	}
	if str == "0n" {
		fmt.Printf("(Strings in array %d are %v)\n\n", currArray, arrayOfStrings)
		arraysOfStrings = append(arraysOfStrings, arrayOfStrings)
		currArray++;currString=1;arrayOfStrings = []string{}
		readString(currArray, currString)
	}	
	arrayOfStrings = append(arrayOfStrings, str)
	currString++
	readString(currArray, currString)
	return
}

func main() {			
	fmt.Println("")
	readString3(1,1)
}