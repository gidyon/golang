package main

import (
	"fmt"
)

/* Balance is a closure (higher order function) that will take a slice of coins/note (Like [1,5,10,20,40,50....1000]). It will return a function that takes the balance (to get coins list of the balance) as parameter which in turn will return an array of contructed coins (balance but in coins) to that when summed will be equal to balance
----------------Assumption made is that the list of coins are sorted in ascending order---------------------*/

func Balance(listOfCoinsAndNotes []int) func(int) string {
	// Create a slice of coins that will be returned (used)
	coinsOrNotesUsed := make([]int, 0)
	// A function value that we will call recursively. It will take balance and index of coins that will be used ( In this case, in the array listOfCoins)
	var findClosestToBal func(int, int) int
	// Assign the fincClosestToBal a value. In this case function value. This will help us to call it recursively
	findClosestToBal = func(bal, i int) int {
		// Create a coin variable
		var coin int
		// Check if bal is less than or equal to value at index i in listOfCoins
		if bal >= listOfCoinsAndNotes[i] {
			// Return the coin at that index
			return listOfCoinsAndNotes[i]
		}
		// We recursively call it again, but it index will be -1 the last index. i.e. listOfCoins[i-1]
		coin = findClosestToBal(bal, i-1)
		// Return the coin
		return coin
	}
	// A function value that will also be called recursively. It will take balance as parameter
	var findBalance func(int)
	findBalance = func(balance int) {
		if balance <= 0 {
			return
		}
		coin := findClosestToBal(balance, (len(listOfCoinsAndNotes) - 1))
		coinsOrNotesUsed = append(coinsOrNotesUsed, coin)
		findBalance(balance - coin)
	}
	return func(balance int) string {
		defer func() {
			coinsOrNotesUsed = []int{}
		}()
		findBalance(balance)
		return fmt.Sprintf("The balance of %v is : %v", balance, coinsOrNotesUsed)
	}
}

func main() {
	var Ksh = []int{1, 5, 10, 20, 40, 50, 100, 200, 500, 1000}
	findBalance := Balance(Ksh)
	fmt.Println(findBalance(829))
	fmt.Println(findBalance(378))
	// Output [100 20 5 1 1 1 1]
}
