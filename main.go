package main

import (
	"fmt"

	"./sle"
)

func main() {
	matr := [][]float64{
		{4, 7, -3, -11, 7, 36},
		{-2, 11, 3, 21, 0, -11},
		{0, 0, 4, -16, -5, 27},
		{26, 0, -7, -1, 18, -5},
		{5, 8, 21, -5, 3, 13}}
	Sle, err := sle.CreateSle(matr)
	if err == nil {
		Sle.Print()
	} else {
		fmt.Println(err)
	}

	sleRes, err := Sle.Solve()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%f\n", sleRes)
	}

	// var bigTest sle.MatrSlice = [][]float64{
	// 	{4, -1, 11, 5, 2},
	// 	{4, -2, 11, -9, 5},
	// 	{1, 2, 3, 4, 5},
	// 	{5, 4, 3, 2, 1},
	// }

	// bigTest.Print()
}
