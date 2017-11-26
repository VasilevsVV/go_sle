package sle

import (
	"errors"
	"fmt"
)

//MatrSlice a type for making methods for simple slice
type MatrSlice [][]float64

// SwpLines swaps lines in matrix
func swpLines(m *[][]float64, l1, l2 int) {
	buf := (*m)[l1]
	(*m)[l1] = (*m)[l2]
	(*m)[l2] = buf
}

// checkSquarness checks if matrix is square
func (m MatrSlice) checkSquarness() int {
	size := len(m)
	for _, l := range m {
		if len(l) != size {
			return -1
		}
	}
	return size
}

func (m MatrSlice) checkRect() bool {
	size := len(m[0])
	for _, l := range m {
		if len(l) != size {
			return false
		}
	}
	return true
}

// GetMinor gets a n-th minor from matrix
func (m MatrSlice) GetMinor(x, y int) MatrSlice {
	var res [][]float64
	for i, l := range m {
		if i != x {
			size := len(l)
			newl := make([]float64, size-1, size-1)
			copy(newl, l[:y])
			for j := y + 1; j < size; j++ {
				newl[j-1] = l[j]
			}
			res = append(res, newl)
		}
	}
	return res
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

//Determinant gets determinant of matrix
func (m MatrSlice) Determinant() (float64, error) {
	size := m.checkSquarness()
	if size < 0 {
		return 0, errors.New("Matrix is not square")
	}
	if size > 1 {
		return m.determinant(), nil
	}
	return 0, errors.New("Matrix is less then 2x2")
}

// Transponate transponates matrix
func (m MatrSlice) Transponate() (MatrSlice, error) {
	if !m.checkRect() {
		return nil, errors.New("Matrix is not Square")
	}
	res := MakeMatrix(len(m), len(m[0]))
	for i, l := range m {
		for j, el := range l {
			res[j][i] = el
		}
	}
	return res, nil
}

//Mult makes new matrix with every value multipied by value f
func (m MatrSlice) Mult(f float64) MatrSlice {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			m[i][j] *= f
		}
	}
	return m
}

//MakeMatrix make and returns MatrSlice
func MakeMatrix(x, y int) MatrSlice {
	res := make(MatrSlice, x, x)
	for i := 0; i < x; i++ {
		res[i] = make([]float64, y, y)
	}
	return res
}

func testForMult(m1, m2 MatrSlice) error {
	if !m1.checkRect() || !m2.checkRect() {
		return fmt.Errorf("Some of matrixes is not rectangle")
	}
	if len(m1[0]) != len(m2) {
		return fmt.Errorf("Number of columns in 1-st matrix not equal to number of lines in 2-nd:\ncol1 = %d | line2 = %d",
			len(m1[0]), len(m2))
	}
	return nil
}

//Multm multipies two matrixes
func (m MatrSlice) Multm(m1 MatrSlice) (MatrSlice, error) {
	err := testForMult(m, m1)
	if err != nil {
		return nil, err
	}
	size := len(m1)
	res := MakeMatrix(len(m1), len(m1[0]))
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			for k := 0; k < size; k++ {
				res[i][j] += m[i][k] * m1[k][j]
			}
		}
	}
	return res, nil
}

func vectorToMatr(vector []float64) MatrSlice {
	res := MakeMatrix(len(vector), 1)
	for i, el := range vector {
		res[i][0] = el
	}
	return res
}

func (m MatrSlice) getSize() int {
	var res int
	for _, l := range m {
		res += len(l)
	}
	return res
}

func (m MatrSlice) matrToVector() []float64 {
	size := m.getSize()
	res := make([]float64, size, size)
	var i int
	for _, l := range m {
		for _, el := range l {
			res[i] = el
			i++
		}
	}
	return res
}

//Print prints a matrix
func (m MatrSlice) Print() {
	for _, l := range m {
		j := 0
		fmt.Printf("| ")
		for ; j < len(l)-1; j++ {
			fmt.Printf("%f, ", l[j])
		}
		fmt.Printf("%f |\n", l[j])
	}
}
