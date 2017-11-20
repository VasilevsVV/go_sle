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
func (m matrSlice) checkSquarness() int {
	size := len(m)
	for _, l := range m {
		if len(l) != size {
			return -1
		}
	}
	return size
}

func (m matrSlice) checkRect() bool {
	size := len(m[0])
	for _, l := range m {
		if len(l) != size {
			return false
		}
	}
	return true
}

// GetMinor gets a n-th minor from matrix
func (m matrSlice) getMinor(n int) matrSlice {
	var res [][]float64
	for i, l := range m {
		if i == n {
			continue
		}
		res = append(res, l[1:])
	}
	return res
}

func (m matrSlice) determinant() float64 {
	length := len(m)
	if length == 2 {
		return m[0][0]*m[1][1] - m[0][1]*m[1][0]
	}
	var res float64
	for i, f := 0, 1.0; i < length; i, f = i+1, -f {
		res += f * m[i][0] * m.getMinor(i).determinant()
	}
	return res
}

//Determinant gets determinant of matrix
func (m matrSlice) Determinant() (float64, error) {
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
func (m matrSlice) Transponate() (matrSlice, error) {
	if m.checkSquarness() >= 0 {
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

func (m matrSlice) Mult(f float64) matrSlice {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			m[i][j] *= f
		}
	}
	return m
}

func makeMatrix(x, y int) matrSlice {
	res := make(matrSlice, x, x)
	for i := 0; i < x; i++ {
		res[i] = make([]float64, y, y)
	}
	return res
}

func testForMult(m1, m2 matrSlice) error {
	if !m1.checkRect() || !m2.checkRect() {
		return fmt.Errorf("Some of matrixes is not rectangle")
	}
	if len(m1[0]) != len(m2) {
		return fmt.Errorf("Number of columns in 1-st matrix not equal to number of lines in 2-nd:\ncol1 = %d | line2 = %d",
			len(m1[0]), len(m2))
	}
	return nil
}

func (m matrSlice) Multm(m1 matrSlice) (matrSlice, error) {
	err := testForMult(m, m1)
	if err != nil {
		return nil, err
	}
	size := len(m1)
	res := makeMatrix(len(m1), len(m1[0]))
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			for k := 0; k < size; k++ {
				res[i][j] += m[i][k] * m1[k][j]
			}
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
