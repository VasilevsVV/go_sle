package main

import "fmt"

func main() {
	matr := [][]float64{{1, 2, 3, 10}, {4, 5, 6, 20}, {7, 8, 9, 30}}
	sle, err := CreateSle(matr)
	if err == nil {
		sle.Print()
	} else {
		fmt.Println(err)
	}

	var test matrSlice = [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	res, err := test.Multm([][]float64{{10}, {20}, {30}})
	if err == nil {
		res.Print()
	} else {
		fmt.Println(err)
	}

	var bigTest matrSlice = [][]float64{{4, -1, 11, 5, 2}, {4, -2, 11, -9, 5}, {1, 2, 3, 4, 5}, {5, 4, 3, 2, 1}, {21, 3, -7, -1, 6}}

	det, err := bigTest.Determinant()
	if err == nil {
		fmt.Printf("Det = %f\n", det)
	} else {
		fmt.Println(err)
	}

	bigTest.getMinor(2, 2).Print()
}
