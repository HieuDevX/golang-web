package main

import (
	"fmt"
)

func sumDivisor(number int) int {
	sum := 0
	for i := 1; i < number; i++ {
		if number%i == 0 {
			sum += i
		}
	}
	return sum
}

func isPerfectNumber(number int) bool {
	return sumDivisor(number) == number
}

func main() {
	fmt.Println("Perfect numbers is smaller than 200")

	for i := 0; i < 8200; i++ {
		if isPerfectNumber(i) {
			fmt.Println(i)
		}
	}
}
