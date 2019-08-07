package main

import (
	"fmt"
)

func isPrimeNumber(number int) bool {
	if number <= 1 {
		return false
	}

	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	n := 11

	if isPrimeNumber((n)) {
		fmt.Printf("%d is prime number \n", n)
	} else {
		fmt.Printf("%d is not prime number \n", n)
	}
}
