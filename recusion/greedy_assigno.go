/*An algorithm that solves the balance the problem simply lists the bills/coins that make up the change when its given a value of the balance
i)  Write a recursive version of this algorithm
ii) Write an iterative version of the algorithm

Kenyan Currency coins/notes [1,5,10,20,40,50,100,200,200,1000]

We want a change of 896

Greedy approach 500, 200, 100, 50, 40, 5, 1

*/
package main

import (
	"fmt"
)

var CoinsList = []int{1, 5, 10, 20, 40, 50, 100, 200, 500, 1000}
var length = len(CoinsList)
var arrCoins []int

func Balance(balance int) {
	if balance <= 0 {
		return
	}
	// I wanted to define findClosestToBal(bal int, i int) inside Balance (inline)
	coin := findClosestToBal(balance, (length - 1))
	arrCoins = append(arrCoins, coin)
	Balance(balance - coin)
}

func findClosestToBal(bal, i int) int {
	var coin int
	if bal >= CoinsList[i] {
		return CoinsList[i]
	}
	coin = findClosestToBal(bal, i-1)
	return coin
}

func main() {
	var balance = 128
	Balance(balance)
	fmt.Println(arrCoins)
	// Output: [100 20 5 1 1 1]
}
