package main

import (
	"fmt"
)

func main() {
	a := [...]int{1, 5, 3, 4, 1, 5, 5, 6, 7, 4}

	count := map[int]int{}

	for _, v := range a {
		count[v]++
	}

	fmt.Println(a)
	fmt.Println(count)

	for key, value := range count {
		if value > 1 {
			fmt.Println(key)
		}
	}
}
