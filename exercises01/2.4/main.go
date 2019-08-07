package main

import (
	"fmt"
)

func isLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}

	if year%100 == 0 {
		return false
	}

	if year%4 == 0 {
		return true
	}

	return false
}

func main() {
	month := 2
	year := 2019
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("31 ngay")
	case 4, 6, 9, 11:
		fmt.Println("30 ngay")
	case 2:
		if isLeapYear(year) {
			fmt.Println("29 ngay")
		} else {
			fmt.Println("28 ngay")
		}
	}
}
