package main

import "fmt"

func pangkat(base, pangkat int) int {
	// your code here

	var hasil int = base
	for i := 1; i < pangkat; i++ {
		hasil *= base
	}
	return hasil
}

func main() {

	//METODE 1 -- AUTO CHECK
	fmt.Println(pangkat(2, 3))  // 8
	fmt.Println(pangkat(5, 3))  // 125
	fmt.Println(pangkat(10, 2)) // 100
	fmt.Println(pangkat(2, 5))  // 32
	fmt.Println(pangkat(7, 3))  // 343
	fmt.Println(pangkat(10, 3)) // 1000

	//METODE 2 -- INPUT CHECK
	/*var input_nilai int
	var input_pangkat int
	fmt.Println("Fungsi Cek Hasil Pangkat")
	fmt.Print("Masukkan sebuah nilai : ")
	fmt.Scan(&input_nilai)
	fmt.Print("Masukkan sebuah nilai pangkat : ")
	fmt.Scan(&input_pangkat)

	fmt.Println(pangkat(input_nilai, input_pangkat))
	*/
}
