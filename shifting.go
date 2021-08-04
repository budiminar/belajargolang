package main

import (
	"fmt"
)

func caesar(offset int, input string) string {
	// your code here

	convertCharRune := []rune(input)
	var result []int
	var convertchar []string
	var getchar string

	for i := 0; i < len(input); i++ {
		result = append(result, int(convertCharRune[i]+rune(offset)))
	}
	for i, v := range result {
		convertchar = append(convertchar, string(v))
		getchar = getchar + convertchar[i]
	}
	return getchar
}

func main() {
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // cnvc
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
