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

func (m MatrSlice) getMinorsMatrix() MatrSlice {
	size := len(m)
	res := MakeMatrix(size, size)
	for i, l := range m {
		for j := range l {
			det := m.GetMinor(i, j).determinant()
			//fmt.Printf("[DETERMINANT = %f]\n", det)
			res[i][j] = det
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
	return res
}

func (m MatrSlice) getInverseMatrix() (MatrSlice, error) {
	det := m.determinant()
	if det == 0 {
		return nil, fmt.Errorf("Determinant of matrix is equal to 0")
	}
	minors := m.getMinorsMatrix()
	compl := minors.getAlgComplemetsMatr()
	transp, _ := compl.Transponate()
	//res, _ := m.getMinorsMatrix().getAlgComplemetsMatr().Transponate()
	return transp.Mult(1.0 / det), nil
}

//Solve returns a slice of solutions for SLE.
func (sle Sle) Solve() ([]float64, error) {
	inverseMatr, err := sle.matrix.getInverseMatrix()
	if err != nil {
		return nil, err
	}
	solutions := MakeMatrix(len(sle.matrix), 1)
	for i, el := range sle.solutions {
		solutions[i][0] = el
	}
	res, err := inverseMatr.Multm(solutions)
	if err != nil {
		return nil, err
	}
	return res.matrToVector(), nil
}

//Print prints out Sle to console
func (sle Sle) Print() {
	fmt.Print(sle.Prints())
}

//Prints returns string with printed sle
func (sle Sle) Prints() string {
	var res string
	for i, l := range sle.matrix {
		j := 0
		res += "|| "
		for ; j < len(l)-1; j++ {
			res += fmt.Sprintf("%f,\t", l[j])
		}
		res += fmt.Sprintf("%f\t| %f ||\n", l[j], sle.solutions[i])
	}
	return res
}
