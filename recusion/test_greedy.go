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
	"time"
)

var Ksh = []int{1, 5, 10, 20, 40, 50, 100, 200, 500, 1000}
var length = len(Ksh)

func findBalTest(balance int, coinsList []int) (arrOfCoinsUsed []int) {
	carriedCoins := arrOfCoinsUsed
	if balance <= 0 {
		return
	}
	coin := findClosestToBal(balance, (length - 1))
	carriedCoins = append(carriedCoins, coin)
	findBalTest(balance-coin, carriedCoins)
	arrOfCoinsUsed = carriedCoins
	return
}

func findClosestToBal(bal, i int) int {
	var coin int
	if bal >= Ksh[i] {
		return Ksh[i]
	}
	coin = findClosestToBal(bal, i-1)
	return coin
}

func main() {
	t1 := time.Now()
	var balance = 128
	fmt.Println(findBalTest(balance, []int{}))
	t2 := time.Now()
	fmt.Print(t2.Sub(t1))
}
