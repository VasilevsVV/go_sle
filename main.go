package main

import (
	"fmt"

	"./sle"
)

func main() {
	matr := [][]float64{{1, 2, -3, 10}, {4, -5, 6, 20}, {-7, 8, 9, 30}}
	Sle, err := sle.CreateSle(matr)
	// if err == nil {
	// 	Sle.Print()
	// } else {
	// 	fmt.Println(err)
	// }

	// var test sle.MatrSlice = [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// res, err := test.Multm([][]float64{{10}, {20}, {30}})
	// if err == nil {
	// 	res.Print()
	// } else {
	// 	fmt.Println(err)
	// }

	// var bigTest sle.MatrSlice = [][]float64{{4, -1, 11, 5, 2}, {4, -2, 11, -9, 5}, {1, 2, 3, 4, 5}, {5, 4, 3, 2, 1}, {21, 3, -7, -1, 6}}

	// transp, _ := bigTest.Transponate()
	// transp.Print()
	// det, err := bigTest.Determinant()
	// if err == nil {
	// 	fmt.Printf("Det = %f\n", det)
	// } else {
	// 	fmt.Println(err)
	// }

	// bigTest.GetMinor(2, 2).Print()

	sleRes, err := Sle.Solve()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%f\n", sleRes)
	}
	fmt.Println(sle.CompareSlices(sleRes, []float64{7.75, 6.5, 43.0 / 12.0}))
}
