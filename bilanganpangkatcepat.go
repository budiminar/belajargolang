package main

import "fmt"

func main() {

	fmt.Println(pow(2, 3))

}

func pow(x, n int) int {

	n = n / 2

	return n

}
