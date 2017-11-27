package sle

import (
	"errors"
	"fmt"
	"math"
)

//MatrSlice a type for making methods for simple slice
type MatrSlice [][]float64

// SwpLines swaps lines in matrix
func swpLines(m *[][]float64, l1, l2 int) {
	buf := (*m)[l1]
	(*m)[l1] = (*m)[l2]
	(*m)[l2] = buf
}

func (m MatrSlice) genMatrixIds() ([]uint64, []uint64) {
	length := len(m)
	lines := make([]uint64, length, length)
	cols := make([]uint64, len(m[0]), len(m[0]))
	for i := range lines {
		lines[i] = uint64(math.Pow(2.0, float64(i)))
	}
	for i := range cols {
		cols[i] = uint64(math.Pow(2.0, float64(i+length)))
	}
	return lines, cols
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
			res = append(res, remove(l, y))
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
	res := MakeMatrix(len(m[0]), len(m))
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

func remove(slice []float64, i int) []float64 {
	size := len(slice)
	res := make([]float64, size-1, size-1)
	copy(res, slice[:i])
	for j := i + 1; j < size; j++ {
		res[j-1] = slice[j]
	}
	return res
}

func removeUint(slice []uint64, i int) []uint64 {
	size := len(slice)
	res := make([]uint64, size-1, size-1)
	copy(res, slice[:i])
	for j := i + 1; j < size; j++ {
		res[j-1] = slice[j]
	}
	return res
}

//CompareSlices is comparing two slices
func CompareSlices(sl1, sl2 []float64) bool {
	if len(sl1) != len(sl2) {
		return false
	}
	for i, el := range sl1 {
		if el != sl2[i] {
			if math.Abs(el-sl2[i]) > math.Pow10(-9) {
				return false
			}
		}
	}
	return true
}

//EqualTo compares two MtrSlices
func (m MatrSlice) EqualTo(m2 MatrSlice) bool {
	if len(m) != len(m2) {
		return false
	}
	for i, l := range m {
		if !CompareSlices(l, m2[i]) {
			return false
		}
	}
	return true
}

//Print prints a matrix
func (m MatrSlice) Print() {
	fmt.Print(m.Prints())
}

//Prints return string with printed matrix
func (m MatrSlice) Prints() string {
	var res string
	for _, l := range m {
		j := 0
		res += "| "
		for ; j < len(l)-1; j++ {
			res += fmt.Sprintf("%f,\t", l[j])
		}
		res += fmt.Sprintf("%f |\n", l[j])
	}
	return res
}
