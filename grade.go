package main

import "fmt"

func main() {

	var nama string
	var nilai int

	fmt.Println("Fungsi Mencari Grade")
	fmt.Print("Masukkan Nama Mahasiswa : ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan Nilai Mahasiswa : ")
	fmt.Scan(&nilai)

	switch {
	case nilai >= 80 && nilai <= 100:
		fmt.Printf("Grade untuk %v adalah : Nilai A", nama)
	case nilai >= 65 && nilai <= 79:
		fmt.Printf("Grade untuk %v adalah : Nilai B", nama)
	case nilai >= 50 && nilai <= 64:
		fmt.Printf("Grade untuk %v adalah : Nilai C", nama)
	case nilai >= 35 && nilai <= 49:
		fmt.Printf("Grade untuk %v adalah : Nilai D", nama)
	case nilai >= 0 && nilai <= 34:
		fmt.Printf("Grade untuk %v adalah : Nilai E", nama)
	default:
		fmt.Printf("Nilai ivalid!")
	}
}
