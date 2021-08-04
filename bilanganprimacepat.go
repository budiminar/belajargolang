package main

import (
	"fmt"
	"time"
)

func main() {

	var bil int
	fmt.Print("Masukkan sebuah bilangan : ")
	fmt.Scan(&bil)

	waktuawal := time.Now()
	if primeNumber(bil) {
		fmt.Printf("%v adalah Bilangan Prima", bil)
		fmt.Println()
	} else {
		fmt.Printf("%v bukan Bilangan Prima", bil)
		fmt.Println()
	}
	waktuakhir := time.Now()
	fmt.Println("Waktu Eksekusi ", (waktuakhir.Second() - waktuawal.Second()), " seconds")

	/*fmt.Println(primeNumber(1000000007))  // true
	fmt.Println(primeNumber(1500450271))  // true
	fmt.Println(primeNumber(1000000000))  // false
	fmt.Println(primeNumber(10000000019)) // true
	fmt.Println(primeNumber(10000000033)) // true*/

}

func primeNumber(number int) bool {

	temp := number - 1

	for i := 2; i < temp; i++ {
		if number%2 == 0 {
			return false
			break
		}
	}
	return true

}
