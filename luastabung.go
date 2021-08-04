package main

import "fmt"

func main() {

	var L float32
	var r int
	var T int
	var phi float32 = 3.14
	//var phi float64 = 22.0 / 7.0

	fmt.Println("Fungsi hitung Luas Permukaan Tabung")
	fmt.Print("Masukkan Nilai Tinggi Tabung : ")
	fmt.Scan(&T)
	fmt.Print("Masukkan Nilai Jari-jari Tabung : ")
	fmt.Scan(&r)

	L = (2 * phi * float32(r)) * (float32(r) + float32(T))
	fmt.Println("Luas Permukaan Tabung Adalah : ", L)

}

func auto_hitungluas() {

	var L float32
	var r float32 = 4.0
	var T float32 = 20.0
	var phi float32 = 3.14
	//var phi float64 = 22.0 / 7.0

	L = (2 * phi * r) * (r + T)
	fmt.Println(L)

}
