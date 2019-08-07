package main

import (
	"fmt"
	"math"
)

// Rectangle struct
type Rectangle struct {
	length int
	width  int
}

// Circle struct
type Circle struct {
	radius float64
}

// Area func
func (r Rectangle) Area() int {
	return (r.length) * (r.width)
}

// Area func
func (c Circle) Area() float64 {
	return (c.radius) * (c.radius) * (math.Pi)
}

// Student struct
type Student struct {
	toan int
	van  int
	anh  int
}

func (s Student) reviewStudent() {
	average := float64(s.toan+s.van+s.anh) / 3

	switch {
	case average < float64(4):
		fmt.Println("F")
	case average >= 4 && average < 5:
		fmt.Println("D")
	case average >= 5 && average < 7:
		fmt.Println("C")
	case average >= 7 && average <= 8:
		fmt.Println("B")
	case average >= 9:
		fmt.Println("A")
	}
}

func main() {
	// var rectangle = Rectangle{
	// 	width:  20,
	// 	length: 30,
	// }

	// fmt.Printf("Area: %d \n", rectangle.Area())

	// var circle = Circle{
	// 	radius: 10,
	// }

	// fmt.Printf("Area: %f \n", circle.Area())

	var student1 = Student{
		toan: 8,
		van:  5,
		anh:  6,
	}
	student1.reviewStudent()
}
