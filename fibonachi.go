package main

import "fmt"

func fibo(n int) int {
	// your code here

	bil0 := 0
	bil1 := 1
	berikutnya := 0
	result := []int{}
	var results int = 0

	if n == 0 {
		return results
	} else {
		for i := 1; i <= n; i++ {
			if i == 1 {
				result = append(result, bil0)
			}
			if i == 2 {
				result = append(result, bil1)
			}
			berikutnya = bil0 + bil1
			bil0 = bil1
			bil1 = berikutnya

			result = append(result, berikutnya)
		}
	}
	//fmt.Println(result[n])
	results = result[n]
	return results

}

func main() {
	fmt.Println(fibo(0))  // 0
	fmt.Println(fibo(1))  // 1
	fmt.Println(fibo(2))  // 1
	fmt.Println(fibo(3))  // 2
	fmt.Println(fibo(5))  // 5
	fmt.Println(fibo(6))  // 8
	fmt.Println(fibo(7))  // 13
	fmt.Println(fibo(9))  // 13
	fmt.Println(fibo(10)) // 55
}
