package main

import "fmt"

func playWithAsterik(n int) {
	// your code here

	for i := 1; i <= n; i++ {
		for j := n; j >= i; j-- {
			fmt.Print(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func main() {

	//METODE 1 -- AUTO CHECK
	playWithAsterik(10)
	/*
	       *
	      * *
	     * * *
	    * * * *
	   * * * * *
	*/

	//METODE 2 -- INPUT CHECK
	/*
		var input_nilai int
		fmt.Println("Fungsi Segitiga Sama Kaki Asterisk")
		fmt.Print("Masukkan sebuah nilai : ")
		fmt.Scan(&input_nilai)

		playWithAsterik(input_nilai)
	*/
}
