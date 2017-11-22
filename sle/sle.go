package sle

import (
	"fmt"
)

// Sle is a structure to represent System of Linear Equations
type Sle struct {
	matrix    MatrSlice
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
	return Sle{nil, nil}, fmt.Errorf("Not valid matrix\n %f\n passed to CreateSle", m)
}

func (m MatrSlice) determinant() float64 {
	length := len(m)
	if length == 2 {
		return m[0][0]*m[1][1] - m[0][1]*m[1][0]
	}
	var res float64
	for i, f := 0, 1.0; i < length; i, f = i+1, -f {
		res += f * m[i][0] * m.GetMinor(i, 0).determinant()
	}
	return res
}

func (m MatrSlice) getMinorsMatrix() MatrSlice {
	size := len(m)
	res := MakeMatrix(size, size)
	for i, l := range m {
		for j := range l {
			res[i][j] = m.GetMinor(i, j).determinant()
		}
	}
	return res
}

func (m MatrSlice) getAlgComplemetsMatr() MatrSlice {
	flag := 1
	res := MakeMatrix(len(m), len(m))
	for i, l := range m {
		for j, el := range l {
			res[i][j] = el * float64(flag)
			flag = -flag
		}
	}
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
