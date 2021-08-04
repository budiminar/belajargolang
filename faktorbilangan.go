package main

import "fmt"

func main() {

	var bilangan int

	fmt.Println("Fungsi Mencetak Faktor Bilangan")
	fmt.Print("Masukkan Bilangan : ")
	fmt.Scan(&bilangan)

	fmt.Printf("Faktor bilangan dari %v adalah : ", bilangan)
	fmt.Println("")
	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Println(i)
		}
	}

}
