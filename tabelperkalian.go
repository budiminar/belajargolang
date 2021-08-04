package main

import "fmt"

func cetakTablePerkalian(number int) {

	var hasil int = number

	for i := 1; i <= hasil; i++ {
		var temp = 0
		for j := hasil; j != 0; j-- {
			temp = temp + i
			if temp >= 10 {
				fmt.Print(temp, "  ")
			} else {
				fmt.Print(temp, "   ")
			}
		}
		fmt.Println()
	}

}

func main() {

	//METODE 1 -- AUTO CHECK
	cetakTablePerkalian(10)

	//METODE 2 -- INPUT CHECK
	/*
		var input_nilai int
		fmt.Println("Fungsi Membuat Tabel Perkalian")
		fmt.Print("Masukkan sebuah nilai : ")
		fmt.Scan(&input_nilai)

		cetakTablePerkalian(input_nilai)
	*/

}
