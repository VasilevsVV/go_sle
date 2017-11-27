package sle

import (
	"fmt"
)

// Sle is a structure to represent System of Linear Equations
type Sle struct {
	matrix          MatrSlice
	solutions       []float64
	enableLog       bool
	lineIds, colIds []uint64
	depth           int
}

func validateMatrix(m [][]float64) (bool, MatrSlice, []float64) {
	size := len(m)
	var resMatrix MatrSlice
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
		lines, cols := matrix.genMatrixIds()
		return Sle{matrix, solutions, false, lines, cols, len(matrix)}, nil
	}
	return Sle{nil, nil, false, nil, nil, 0}, fmt.Errorf("Not valid matrix\n %f\n passed to CreateSle", m)
}

func (sle Sle) cloneSle(matrix MatrSlice, solutions []float64) Sle {
	return Sle{matrix, solutions, sle.enableLog, sle.lineIds, sle.colIds, sle.depth}
}

func (sle Sle) log(str string) {
	if sle.enableLog {
		fmt.Print(str)
	}
}

// EnableLog enables log messages in sle
func (sle *Sle) EnableLog() {
	(*sle).enableLog = true
}

// DisableLog disables log messages in sle
func (sle *Sle) DisableLog() {
	(*sle).enableLog = false
}

func (sle Sle) getMinorsMatrix() Sle {
	size := len(sle.matrix)
	res := MakeMatrix(size, size)
	for i, l := range sle.matrix {
		for j := range l {
			det := sle.matrix.GetMinor(i, j).determinant()
			sle.log(fmt.Sprintf("[DETERMINANT = %f]\n", det))
			res[i][j] = det
		}
	}
	return sle.cloneSle(res, sle.solutions)
}

func (sle Sle) getAlgComplemetsMatr() Sle {
	flag, lineFlag := 1, 1
	res := MakeMatrix(len(sle.matrix), len(sle.matrix))
	for i, l := range sle.matrix {
		flag = lineFlag
		for j, el := range l {
			res[i][j] = el * float64(flag)
			flag = -flag
		}
		lineFlag = -lineFlag
	}
	return sle.cloneSle(res, sle.solutions)
}

func (sle Sle) transponate() Sle {
	transp, _ := sle.matrix.Transponate()
	return sle.cloneSle(transp, sle.solutions)
}

func (sle Sle) getInverseMatrix() (Sle, error) {
	det := sle.matrix.determinant()
	if det == 0 {
		return sle.cloneSle(nil, nil), fmt.Errorf("Determinant of matrix is equal to 0")
	}
	res := sle.getMinorsMatrix().getAlgComplemetsMatr().transponate()
	res.matrix = res.matrix.Mult(1.0 / det)
	return res, nil
}

//Solve returns a slice of solutions for SLE.
func (sle Sle) Solve() ([]float64, error) {
	sle.log(fmt.Sprintf("Lines : %d\nColumns : %d\n", sle.lineIds, sle.colIds))
	inverseMatr, err := sle.getInverseMatrix()
	if err != nil {
		return nil, err
	}
	solutions := MakeMatrix(len(sle.matrix), 1)
	for i, el := range sle.solutions {
		solutions[i][0] = el
	}
	res, err := inverseMatr.matrix.Multm(solutions)
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
