package main

import (
	"fmt"

	_ "github.com/satori/go.uuid"

	chuvi "./chuvi"

	dientich "./dientich"
)

func main() {
	a, b := 5, 7
	fmt.Printf("Chu vi la : %v \n", chuvi.ChuVi(a, b))
	fmt.Printf("Dien tich la : %v \n", dientich.DienTich(a, b))
}
