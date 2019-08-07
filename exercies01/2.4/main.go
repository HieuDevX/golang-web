package main

import (
	"fmt"
)

func main() {
	month := 8
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("31 ngay")
	case 4, 6, 9, 11:
		fmt.Println("30 ngay")
	case 2:
		fmt.Println("28 hoac 29 ngay")
	}
}
