package main

import "fmt"

func BinarySearch(array []int, x int) {
	// your code here
	result := array
	getArr := result[0:x]
	jumlah := getArr[0] + getArr[x-1]
	avg := 0

	avg = jumlah / 2
	fmt.Println(avg)
}

func main() {
	BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3)                  // 2
	BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5)                 // 3
	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53)  // 6
	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100) // -1
}
