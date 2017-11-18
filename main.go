package main

import "fmt"

func main() {
	fmt.Println("Test")
	m := [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	SwpLines(&m, 0, 2)
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(len(m[1]))
}
