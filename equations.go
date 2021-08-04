package main

import "fmt"

func SimpleEquations(a, b, c int) {
	// your code here
	result := []int{}
	cekA := 0
	cekB := 0
	cekC := 0
	if a != 0 {
		for i := 1; i < a; i++ {
			if a%i == 0 {
				result = append(result, i)
				cekA = cekA + i
				cekB = cekB + (1 * i)
				cekC = cekC + (i * i)
			}
		}

		if cekA == a && cekB == b && cekC == c {
			fmt.Println(result)
		} else {
			fmt.Println("No Solution")
		}
	}
}

func main() {
	SimpleEquations(1, 2, 3)  // no solution
	SimpleEquations(6, 6, 14) // 1 2 3
}
