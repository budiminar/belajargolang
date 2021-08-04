package main

import "fmt"

func DragonOfLoowater(dragonHead, knightHeight []int) {
	// your code here
	result := []int{}
	jumlah := 0
	for i, value := range dragonHead {
		for j, _ := range knightHeight {
			if dragonHead[i] >= knightHeight[j] {
				if i == 0 {
					result = append(result, value)
					jumlah = jumlah + value
					i++
				} else {
					if value != dragonHead[i-1] {
						result = append(result, value)
						jumlah = jumlah + value
					} else {

					}
				}
			}
		}
	}
	fmt.Println(jumlah)
	fmt.Println(result[:])
}

func main() {
	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4})    // 11
	DragonOfLoowater([]int{5, 10}, []int{5})         // knight fall
	DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2}) // knight fall
	DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5}) // 10
}
