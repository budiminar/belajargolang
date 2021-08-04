package main

import "fmt"

func main() {

	var input_string string
	fmt.Println("'Hello, world'")
	fmt.Scan(&input_string)
	//count_bilangan_prima()
	//show_bilangan_prima()
	//bil_ganjil()
}

func count_bilangan_prima() {

	var input_numb int
	fmt.Print("Enter number : ")
	fmt.Scan(&input_numb)

	for i := 1; i <= input_numb; i++ {
		x := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				x++
			}
		}
		if x == 2 {
			fmt.Println(i)
		}
	}

}

func show_bilangan_prima() {

	var input_numb int
	fmt.Print("Enter number : ")
	fmt.Scan(&input_numb)

	for i := 1; i <= input_numb; i++ {
		x := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				x++
			}
		}
		if x == 2 && i == input_numb {
			fmt.Println("bilangan prima")
		} else if x > 2 && i == input_numb {
			fmt.Println("bukan bilangan prima")
		}
	}

}

func bil_ganjil() {

	var input_numb int
	fmt.Print("Enter Number : ")
	fmt.Scan(&input_numb)

	if input_numb%2 != 0 {
		fmt.Println("lampu menyala")
	} else {
		fmt.Println("lampu mati")
	}

}
