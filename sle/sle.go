package sle

import (
	"fmt"
)

// Sle is a structure to represent System of Linear Equations
type Sle struct {
	matrix    MatrSlice
	solutions []float64
	size      int
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
		return Sle{matrix, solutions, len(matrix)}, nil
	}
	return Sle{nil, nil, 0}, fmt.Errorf("Not valid matrix\n %f\n passed to CreateSle", m)
}

func (sle Sle) getMinorsMatrix() MatrSlice {
	size := len(sle.matrix)
	res := MakeMatrix(size, size)
	for i, l := range sle.matrix {
		for j := range l {
			el, _ := sle.matrix.GetMinor(i, j).Determinant()
			res[i][j] = el
		}
	}
	return res
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
