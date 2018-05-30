package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	Title    string
	Price    float64
	Quantity int64
}

var products = make([]Product, 0)

func main() {
	file, err := os.Open("product1.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	bufioReader := bufio.NewReader(file)
	product := Product{}
	for {
		dataString, err := bufioReader.ReadString('\n')
		if err == io.EOF {
			break
		}
		dataSlice := strings.Split(dataString, ";")
		product.Title = dataSlice[0]
		product.Price, _ = strconv.ParseFloat(dataSlice[1], 32)
		product.Quantity, _ = strconv.ParseInt(dataSlice[2], 10, 32)
		products = append(products, product)
	}
	fmt.Println(products)
}
