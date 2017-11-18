package main

import "errors"

// Sle is a structure to represent System of Linear Equations
type Sle struct {
	matrix    [][]float64
	solutions []float64
}

func validateMatrix(m [][]float64) (bool, [][]float64, []float64) {
	size := len(m)
	var resMatrix [][]float64
	var solutions []float64
	for i := 0; i < size; i++ {
		length := len(m[i])
		if length != size+1 {
			return false, nil, nil
		}
		for j := 0; j < size-1; j++ {
			resMatrix[i][j] = m[i][j]
		}
		solutions[i] = m[i][length-1]
	}
	return false, resMatrix, solutions
}

//CreateSle create System of Linear Equations from simple matrix
func CreateSle(m [][]float64) (Sle, error) {
	test, matrix, solutions := validateMatrix(m)
	if test {
		return Sle{matrix, solutions}, nil
	}
	return Sle{nil, nil}, errors.New("Not valid matrix passed to CreateSle")

}

// SwpLines swaps lines in matrix
func SwpLines(m *[][]float64, l1, l2 int) {
	buf := (*m)[l1]
	(*m)[l1] = (*m)[l2]
	(*m)[l2] = buf
}
