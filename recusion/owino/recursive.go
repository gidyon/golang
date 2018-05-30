package main

import (
	"fmt"
)

/* Balance is a closure (higher order function) that will take a slice of coins/note (Like [1,5,10,20,40,50....1000]). It will return a function that takes the balance (to get coins list of the balance) as parameter which in turn will return an array of contructed coins (balance but in coins) to that when summed will be equal to balance
----------------Assumption made is that the list of coins are sorted in ascending order---------------------*/

func BalanceCalculator(denominations []int) func(int) string {
	// denominations = func() []int {

	// }()
	var findOptimalCoinOrBill func(int, int) int
	findOptimalCoinOrBill = func(bal, i int) int {
		coinOrBill := 1
		if bal >= denominations[i] {
			return denominations[i]
		}
		coinOrBill = findOptimalCoinOrBill(bal, i-1)
		return coinOrBill
	}
	var findBalanceRecursive func(int)
	return func(balance int) string {
		coinsAndBillsUsed := make([]int, 0)
		findBalanceRecursive = func(balance int) {
			if balance <= 0 {
				return
			}
			coinOrBill := findOptimalCoinOrBill(balance, (len(denominations) - 1))
			coinsAndBillsUsed = append(coinsAndBillsUsed, coinOrBill)
			findBalanceRecursive(balance - coinOrBill)
		}
		findBalanceRecursive(balance)
		return fmt.Sprintf("\nThe balance of %v is : %v\n", balance, coinsAndBillsUsed)
	}
}

func main() {
	var Ksh = []int{1, 5, 10, 20, 40, 50, 100, 200, 500, 1000}

	findBalance := BalanceCalculator(Ksh)
	fmt.Println(findBalance(829))
	fmt.Println(findBalance(378))
	fmt.Println(findBalance(540))
	fmt.Println(findBalance(214))
}
