package main

import (
	"fmt"
)

func palindrome(input string) bool {

	var kata string = input
	var kata_awal string
	var kata_reverse string
	var palindrome bool

	for i := 0; i <= len(kata)-1; i++ {
		kata_awal = string(kata[i])
		//fmt.Print(string(kata[i]), " ")
	}
	for j := len(kata) - 1; j >= 0; j-- {
		kata_reverse = string(kata[j])
		//fmt.Print(string(kata[j]), " ")
	}

	if kata_awal == kata_reverse {
		fmt.Printf("Kata %v adalah Palindrome / ", input)
		palindrome = true
	} else {
		fmt.Printf("Kata %v adalah Non Palindrome / ", input)
		palindrome = false
	}
	return palindrome
}

func main() {

	//METODE 1 -- AUTO CHECK

	fmt.Println(palindrome("civic"))       // true
	fmt.Println(palindrome("katak"))       // true
	fmt.Println(palindrome("kasur rusak")) // true
	fmt.Println(palindrome("mistar"))      // false
	fmt.Println(palindrome("lion"))        // false

	//METODE 2 -- INPUT CHECK
	/*
		var input_kata string
		fmt.Println("Fungsi Cek Kata Palindrome")
		fmt.Print("Masukkan sebuah kata : ")
		fmt.Scan(&input_kata)

		fmt.Println(palindrome(input_kata))
	*/

}
