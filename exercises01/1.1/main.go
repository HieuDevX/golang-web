package main

import (
	"fmt"
	"math"
)

func area(radius float64) float64 {
	return radius * radius * (math.Pi)
}

func perimeter(radius float64) float64 {
	return 2 * radius * (math.Pi)
}

func main() {
	radius := float64(10)

	fmt.Printf("area of circle which has radius %.2f is: %.2f \n", radius, area(radius))
	fmt.Printf("perimeter of circle which has radius %.2f is: %.2f \n", radius, perimeter(radius))
}
