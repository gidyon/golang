package main

import (
	//"bufio"
	//"os"
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

func main() {	
	// str1 := []string{"brainfuck", "98", "okie"}
	// str2 := []string{"dojo", "sensei", "@!89xhj"}
	// result := scanRanges(str1, str2,[]string{"trump$", "kanye#@west"})
	// fmt.Print(result)	 
	/*
		[[[97 117] [56 57] [101 111]] [[100 111] [101 115] [33 120]][[36 117] [35 121]]]
	*/
	fmt.Println(scanRanges([]string{"tset", "()tset"}))
}