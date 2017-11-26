package main

import (
	"fmt"

	"./sle"
)

func getMediumMatrix() sle.MatrSlice {
	return [][]float64{
		{12, 15, -3, 56, 46, 3, 56, -97, 6},
		{34, -8, 87, 5, 45, 6, 34, 6, -87},
		{43, 24, -9, 8, 36, 21, 34, 67, -9},
		{45, 6, 56, -76, 65, 42, 54, -8, -76},
		{56, 43, 76, 5, -8, 87, 53, 24, -87},
		{-56, 20, 43, 43, 34, -17, 72, 43, 23},
		{76, 46, -73, 35, 85, 84, 35, 89, -65},
		{45, 76, -73, 23, 16, 8, -43, 24, 75},
	}
}

func getBigMatrix() sle.MatrSlice {
	return [][]float64{
		{12, 15, -3, 56, 46, 3, 56, -97, 6, 87, 13, -9, 98},
		{34, -8, 87, 5, 45, 6, 34, 6, -87, 6, 54, 35, -7},
		{43, 24, -9, 8, 36, 21, 34, 67, -9, 8, 6, 4, 8},
		{45, 6, 56, -76, 65, 42, 54, -8, -76, 53, 5, 6, -43},
		{56, 43, 76, 5, -8, 87, 53, 24, -87, 54, 34, 3, 16},
		{-56, 20, 43, 43, 34, -17, 72, 43, 23, -6, 13, -32, 2},
		{76, 46, -73, 35, 85, 84, 35, 89, -65, 42, 56, 73, 91},
		{45, 76, -73, 23, 16, 8, -43, 24, 75, 27, -81, 30, 73},
		{4, 26, 34, 95, 27, 96, -7, -36, 75, 82, 93, 26, 75},
		{36, 8, 35, 75, -6, 49, 76, 56, 45, 75, -65, 65, 86},
		{56, 86, 45, 14, -63, 16, 22, 23, 25, 63, 76, -76, 56},
		{26, -65, 87, 4, -45, 35, 75, 83, 47, 29, 64, 35, -54}}
}

func main() {
	// matr := [][]float64{
	// 	{4, 7, -3, -11, 7, 36},
	// 	{-2, 11, 3, 21, 0, -11},
	// 	{0, 0, 4, -16, -5, 27},
	// 	{26, 0, -7, -1, 18, -5},
	// 	{5, 8, 21, -5, 3, 13}}
	matr := getMediumMatrix()
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

	var bigTest sle.MatrSlice = [][]float64{
		{4, -1, 11, 5, 2},
		{4, -2, 11, -9, 5},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
	}

	bigTest.Print()
}
