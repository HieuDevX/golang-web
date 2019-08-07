package main

import (
	"fmt"
)

func main() {
	number := 6
	result := 1

	for i := number; i > 0; i-- {
		result *= i
	}

	fmt.Printf("factorial of %d is %d", number, result)
}
