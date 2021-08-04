package main

import "fmt"

type Student struct {
	name  []string
	score []int
}

func (s Student) Avarage() float64 {
	var rata float64
	var jumlah float64
	for _, value := range s.score {
		jumlah = jumlah + float64(value)
	}
	rata = jumlah / float64(len(s.score))
	return float64(rata)
}

func (s Student) Min() (min int, name string) {
	min = s.score[0]
	name = s.name[0]
	for i, value := range s.score {
		if value < min {
			min = value
			name = s.name[i]
		}
	}
	return min, name
}

func (s Student) Max() (max int, name string) {
	max = s.score[0]
	name = s.name[0]
	for i, value := range s.score {
		if value > max {
			max = value
			name = s.name[i]
		}
	}
	return max, name
}

func main() {
	var a = Student{}

	for i := 0; i < 6; i++ {
		var name string
		fmt.Print("Input " + string(i) + " Student’s Name : ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input " + name + " Score : ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAvarage Score Students is", a.Avarage())
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ")")
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ")")
}
