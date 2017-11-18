package main

import (
	"errors"
	"fmt"
)

type matrSlice [][]float64

// SwpLines swaps lines in matrix
func swpLines(m *[][]float64, l1, l2 int) {
	buf := (*m)[l1]
	(*m)[l1] = (*m)[l2]
	(*m)[l2] = buf
}

// checkSquarness checks if matrix is square
func (m matrSlice) checkSquarness() bool {
	size := len(m)
	for _, l := range m {
		if len(l) != size {
			return false
		}
	}
	return true
}

// GetMinor gets a n-th minor from matrix
func (m matrSlice) GetMinor(n int) (matrSlice, error) {
	if !m.checkSquarness() {
		return nil, errors.New("Matrix is not square")
	}
	var res [][]float64
	for i, l := range m {
		if i == n {
			continue
		}
		res = append(res, l[1:])
	}
	return res, nil
}

// Transponate transponates matrix
func (m matrSlice) Transponate() (matrSlice, error) {
	if !m.checkSquarness() {
		return nil, errors.New("Matrix is not Square")
	}
	res := make(matrSlice, len(m), len(m))
	for _, l := range m {
		for j, el := range l {
			res[j] = append(res[j], el)
		}
	}
	return res, nil
}

func (m matrSlice) Print() {
	for _, l := range m {
		j := 0
		fmt.Printf("| ")
		for ; j < len(l)-1; j++ {
			fmt.Printf("%f, ", l[j])
		}
		fmt.Printf("%f |\n", l[j])
	}
}
