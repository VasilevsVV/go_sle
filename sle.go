package main

import (
	"errors"
	"fmt"
)

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
		resMatrix = append(resMatrix, m[i][:length-1])
		solutions = append(solutions, m[i][length-1])
	}
	return true, resMatrix, solutions
}

//CreateSle creates System of Linear Equations from simple matrix
func CreateSle(m [][]float64) (Sle, error) {
	test, matrix, solutions := validateMatrix(m)
	if test {
		return Sle{matrix, solutions}, nil
	}
	return Sle{nil, nil}, errors.New("Not valid matrix passed to CreateSle")
}

//Print prints out Sle to console
func (sle Sle) Print() {
	for i, l := range sle.matrix {
		j := 0
		for ; j < len(l)-1; j++ {
			fmt.Printf("%f, ", l[j])
		}
		fmt.Printf("%f | %f\n", l[j], sle.solutions[i])
	}
}
