package main

import "fmt"

func swap(a, b *int) {

	var valA, valB, temp int

	valA = *a
	valB = *b

	temp = valA
	valA = *b
	valB = temp
	fmt.Printf("a : %v   b : %v", valA, valB)
}

func main() {
	a := 10
	b := 20

	swap(&a, &b)
	//fmt.Println(a, b)
}
