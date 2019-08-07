package main

import (
	"fmt"
)

func reverseArray(a []int) []int {
	result := []int{}
	for i := len(a) - 1; i >= 0; i-- {
		result = append(result, a[i])
	}

	return result
}

func main() {
	a := []int{5, 6, 7, 8, 1}

	fmt.Printf("a: ")
	fmt.Println(a)

	b := reverseArray(a)
	fmt.Printf("b: ")
	fmt.Println(b)

}
