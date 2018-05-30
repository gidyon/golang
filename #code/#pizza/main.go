package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	ErrPartOfSlice         = errors.New("overlap: point part of another slice")
	ErrColValidationFailed = errors.New("column validation failed")
	ErrRowValidationFailed = errors.New("row validation failed")
)

var pissa = [][]cell{}

type pizza struct {
	rows          int16
	columns       int16
	currRow       int16
	currCol       int16
	maxCells      int16
	minIngredient int8
	nextRow       bool
	cells         [][]cell
	slices        []slice
	newPoint      *point
}

type point struct {
	row int16
	col int16
}

type slice struct {
	row1    int16
	col1    int16
	row2    int16
	col2    int16
	currRow int16
	currCol int16
	H       int16
	M       int16
	T       int16
	D       string
	*pizza
}

func newSlice(startingPoint *point) (*slice, error) {
	cell := pissa[startingPoint.row][startingPoint.col]
	if cell.isSliced {
		return nil, ErrPartOfSlice
	}
	newSlice := &slice{
		row1: startingPoint.row,
		col1: startingPoint.col,
		row2: startingPoint.row,
		col2: startingPoint.col,
	}
	if cell.food == "T" {
		newSlice.T++
	}
	if cell.food == "M" {
		newSlice.M++
	}
	newSlice.H++
	return newSlice, nil
}

func (s *slice) enlargeSlice() {
	cell := s.pizza.cells[s.row1][s.col1]
	if s.pizza.nextRow {
		s.col1, s.col2 = s.pizza.currCol, s.pizza.currCol
		s.row1, s.row2 = s.pizza.currRow, s.pizza.currCol
	}
	if cell.isSliced {
		s.col1++
		s.enlargeSlice()
		s.col1--
		s.row1++
		s.enlargeSlice()
	}
	if s.D == "right" {
		// if s.col2 == piz.columns {
		// 	s.D = "down"
		// 	s.enlargeSlice()
		// }
		cell := s.pizza.cells[s.row1][s.col1]
		if cell.isSliced {
			s.col1++
			s.col2++
			s.D = "down"
			s.enlargeSlice()
		} else {
			if cell.food == "T" {
				s.T++
			}
			if cell.food == "M" {
				s.M++
			}
			s.col2++
			s.H++
			s.D = "down"
			s.enlargeSlice()
		}
	}
	if s.D == "down" {
		cell := pissa[s.row2+1][s.col2]
		if cell.isSliced {
			s.D = "right"
			s.enlargeSlice()
		} else {
			if cell.food == "T" {
				s.T++
			}
			if cell.food == "M" {
				s.M++
			}
			s.row2++
			s.H++
			s.D = "left"
			s.enlargeSlice()
		}
	}
	if s.D == "left" {
		cell := pissa[s.row2+1][s.col1]
		if cell.isSliced {
			s.D = "right"
			s.enlargeSlice()
		} else {
			if cell.food == "T" {
				s.T++
			}
			if cell.food == "M" {
				s.M++
			}
			s.row2++
			s.H++
			s.D = "left"
			s.enlargeSlice()
		}
	}
	if direction == "up" {
		cell := pissa[s.row1][s.col1]
		if cell.isSliced {
			point := &point{
				row: s.row2 + 1,
				col: s.col2 + 1,
			}
			// slice, err := newSlice(point)
			// if err != nil {
			// 	log.Fatalln(err)
			// }
		}
		cell.isSliced = true
		s.enlargeSlice("down")
	}

}

func (s *slice) satisfyMin(minIngredient int16) bool {
	// var M, T int8 = 0, 0
	// for row := s.row1; row > s.row2; row++ {
	// 	for col := s.col1; col > s.col2; col++ {
	// 		cell := pissa[row][col]
	// 		if cell.food == "T" {
	// 			T++
	// 		}
	// 		if cell.food == "M" {
	// 			M++
	// 		}
	// 	}
	// }
	if s.M < minIngredient || s.T < minIngredient {
		return false
	}
	return true
}

func (s *slice) satisfyMax(maxCells int16) bool {
	// width := s.row2 - s.row1
	// height := s.col2 - s.col1
	cellsNumber := s.M * s.T
	if cellsNumber > maxCells {
		return false
	}
	return true
}

type cell struct {
	food     string
	isSliced bool
}

func piza(inputFile string) error {
	fileReader, err := os.Open(inputFile)
	if err == os.ErrNotExist {
		return err
	}
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(fileReader)
	// We want the first line first, the store it in a slice
	scanner.Scan()
	info := make([]uint64, 0)
	strArr := strings.Split(scanner.Text(), " ")
	for i := range strArr {
		val, err := strconv.ParseUint(strArr[i], 10, 8)
		if err != nil {
			return err
		}
		info = append(info, val)
	}
	if len(info) != int(info[0]) {
		return ErrRowValidationFailed
	}
	pizza := &pizza{
		rows:          int16(info[0]),
		columns:       int16(info[1]),
		currCol:       0,
		currRow:       0,
		maxCells:      int16(info[3]),
		minIngredient: int8(info[2]),
		nextRow:       false,
		cells:         make([][]cell, 0, info[0]),
		newPoint:      &point{0, 0},
	}
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), " ")
		if len(row) != int(info[1]) {
			return ErrColValidationFailed
		}
		rowCells := make([]cell, 0)
		for _, item := range row {
			rowCells = append(rowCells, cell{food: item, isSliced: false})
		}
		pizza.cells = append(pizza.cells, rowCells)
	}
	// We want to create slices, now we should have a default start of 4 cells per slice, in case the condition is unfullfilled, we expand our slice
	newSlice := &slice{
		row1:    0,
		col1:    0,
		row2:    0,
		col2:    0,
		currRow: 0,
		currCol: 0,
		M:       0,
		T:       0,
		D:       "right",
	}
	for pizza.currCol <= pizza.columns {
		newSlice.enlargeSlice()
		pizza.slices = append(pizza.slices, *newSlice)
	}
	return nil
}

func main() {
	fmt.Println(piza("small.in"))
}
