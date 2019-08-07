package main

import (
	"fmt"
)

func main() {
	a := [...]int{5, 4, 6, 7, 8, 9}
	m, n := 10, 3

	fmt.Println(a)
	fmt.Printf("m : %d \n", m)
	fmt.Printf("n : %d \n", n)

	result := []int{}
	result = append(result, a[:n]...)
	result = append(result, m)
	result = append(result, a[n:]...)
	fmt.Println(result)

}
