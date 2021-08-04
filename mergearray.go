package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	// your code here
	//var arr []string = []string{}
	var merge []string
	for i, v := range arrayB {
		//fmt.Println(i, "==> ", v)
		var arr [6]string
		for j, x := range arrayA {
			//fmt.Println(j, "==> ", x)
			if v != x {
				arr[i] = v
				merge = append(arrayA, arr[:]...)
			}
		}
	}
	return merge
}

func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"]

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa", "yoshimitsu", "devil jin", "law"]

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// ["devil jin", "sergei"]

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// ["hwoarang"]

	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []
}
