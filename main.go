package main

import "fmt"

func main() {
	sle, err := CreateSle([][]float64{{1, 2, 3, 10}, {4, 5, 6, 20}, {7, 8, 9, 30}})
	if err == nil {
		sle.Print()
	} else {
		fmt.Println(err)
	}

	var test matrSlice = [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	res, _ := test.GetMinor(1)
	res.Print()
	res1, _ := test.Transponate()
	res1.Print()
}
