package main

import (
	"fmt"
	"time"
)

func main() {

	/*var input_numb int
	fmt.Print("Masukkan bilangan yang akan di cek : ")
	fmt.Scan(&input_numb)*/

	waktuawal := time.Now()
	fmt.Println(primeNumber(1000000007))  // true
	fmt.Println(primeNumber(1500450271))  // true
	fmt.Println(primeNumber(1000000000))  // false
	fmt.Println(primeNumber(10000000019)) // true
	fmt.Println(primeNumber(10000000033)) // true
	waktuakhir := time.Now()
	fmt.Println("Waktu Eksekusi ", (waktuakhir.Second()-waktuawal.Second())/1_000_000, " seconds")

}

func primeNumber(input_numb int) bool {

	for i := 1; i <= input_numb; i++ {
		x := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				x++
			}
		}
		if x == 2 && i == input_numb {
			fmt.Printf("%v adalah bilangan prima ", i)
			//return true
		} else if x > 2 && i == input_numb {
			fmt.Printf("%v adalah bukan bilangan prima ", i)
			return false
		}
	}
	return true

}
