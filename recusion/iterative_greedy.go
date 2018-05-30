package main

import "fmt"

func greedyChange() {
	var coins = []int{10, 5, 1}
	var amount = 28
	var num int
	for i := 0; i < len(coins); i++ {
		if coins[i] >= amount {
			num = amount / coins[i]
			amount -= num * coins[i]
			fmt.Println(num, " Ksh", coins[i])
		}
		fmt.Println(num, " Ksh", coins[i])
	}
}

func main() {
	greedyChange()
}
