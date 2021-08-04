package main

import "fmt"

func moneyCoins(money int) []int {
	// your code here
	arr := []int{1, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000}
	result := []int{}
	n := money

	for i := 1; i <= len(arr); i++ {
		for j := len(arr) - 1; j > 0; j-- {
			if n-arr[len(arr)-i] >= 0 {
				n = n - arr[len(arr)-i]
				result = append(result, arr[len(arr)-i])
			}
		}
	}
	return result
}

func main() {
	fmt.Println(moneyCoins(123))   // [100 20 1 1 1]
	fmt.Println(moneyCoins(432))   // [200 200 20 10 1 1]
	fmt.Println(moneyCoins(543))   // [500, 20, 20, 1, 1, 1]
	fmt.Println(moneyCoins(7752))  // [5000, 2000, 500, 200, 50, 1, 1]
	fmt.Println(moneyCoins(15321)) // [10000 5000 200 100 20 1]
}
