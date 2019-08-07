package main

import (
	"fmt"
	_ "fmt"
)

// Student struct
type Student struct {
	FullName     string
	MidSemGrade  int
	EndSemGrade  int
	OverallGrade float64
}

func (s *Student) calculateOverallGrade() {
	s.OverallGrade = float64(s.EndSemGrade*2+s.MidSemGrade) / 3
}

// BubbleSort func
func BubbleSort(students []Student) []Student {
	for i := len(students); i > 0; i-- {
		for j := 1; j < i; j++ {
			if students[j-1].OverallGrade < students[j].OverallGrade {
				students[j], students[j-1] = students[j-1], students[j]
			}
		}
	}

	return students
}

func main() {
	ListStudent := []Student{
		{"Liquid.Miracle", 9, 10, 0},
		{"Liquid.GH", 6, 5, 0},
		{"VP.NoOne", 7, 8, 0},
		{"VP.Solo", 8, 8, 0},
		{"Navi.Dendi", 4, 6, 0},
	}

	// Not working
	// for _, v := range ListStudent {
	// 	v.calculateOverallGrade()
	// }

	for i := 0; i < len(ListStudent); i++ {
		ListStudent[i].calculateOverallGrade()
	}

	for i, v := range ListStudent {
		fmt.Printf("student %d : %+v \n", i+1, v)
	}

	fmt.Println("After sorting")
	for i, v := range BubbleSort(ListStudent) {
		fmt.Printf("student %d : %+v \n", i+1, v)
	}
}
