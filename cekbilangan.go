package main

import (
	"fmt"
)

func munculSekali(angka string) []int {
	// your code here
	for _, v := range angka {
		for _, val := range angka {
			if v == val {
				fmt.Print("a")
			}
		}
	}
	return angka
}

func main() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6, 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
}
