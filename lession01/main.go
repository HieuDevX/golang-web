package main

import (
	"fmt"

	"github.com/hieudevx/golang-web/lession01/area"
	"github.com/hieudevx/golang-web/lession01/perimeter"
)

func main() {
	a, b := 5, 7
	fmt.Printf("Perimeter is : %v \n", perimeter.Perimeter(a, b))
	fmt.Printf("Area is : %v \n", area.Area(a, b))
}
