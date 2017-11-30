package sle

import (
	"fmt"
	"runtime"
	"sync"
)

// Sle is a structure to represent System of Linear Equations
type Sle struct {
	matrix          MatrSlice
	solutions       []float64
	enableLog       bool
	lineIds, colIds []uint64
	depth           int
	cache           *map[uint64]float64
	mutex           *sync.Mutex
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
		return Sle{matrix, solutions, false, lines, cols, len(matrix),
			&map[uint64]float64{}, &sync.Mutex{}}, nil
	}
	return Sle{nil, nil, false, nil, nil, 0, nil, nil}, fmt.Errorf("Not valid matrix\n %f\n passed to CreateSle", m)
}

func (sle Sle) cloneSle(matrix MatrSlice, solutions []float64) Sle {
	return Sle{matrix, solutions, sle.enableLog, sle.lineIds, sle.colIds, sle.depth, sle.cache, sle.mutex}
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

func (sle Sle) cloneMinor(matrix MatrSlice, lineIds, colIds []uint64) Sle {
	return Sle{matrix, nil, sle.enableLog, lineIds, colIds, sle.depth - 1, sle.cache, sle.mutex}
}

func (sle Sle) getMinor(x, y int) Sle {
	sle.matrix = sle.matrix.GetMinor(x, y)
	sle.lineIds = removeUint(sle.lineIds, x)
	sle.colIds = removeUint(sle.colIds, y)
	sle.depth--
	return sle
}

func (sle Sle) getIndex() uint64 {
	var res uint64
	for i := range sle.lineIds {
		res += sle.lineIds[i] + sle.colIds[i]
	}
	return res
}

func (sle Sle) determinantAux() float64 {
	switch {
	case sle.depth == 2:
		return sle.matrix[0][0]*sle.matrix[1][1] -
			sle.matrix[0][1]*sle.matrix[1][0]
	default:
		var res float64
		for i, f := 0, 1.0; i < sle.depth; i, f = i+1, -f {
			res += f * sle.matrix[i][0] * sle.getMinor(i, 0).determinant()
		}
		return res
	}
}

func (sle Sle) determinant() float64 {
	index := sle.getIndex()
	sle.mutex.Lock()
	res, ok := (*sle.cache)[index]
	sle.mutex.Unlock()
	if ok {
		if sle.depth >= len(sle.solutions)-1 {
			sle.log(fmt.Sprintf("Determinant = %f\n", res))
		}
		return res
	}
	res = sle.determinantAux()
	sle.mutex.Lock()
	(*sle.cache)[index] = res
	sle.mutex.Unlock()
	return res
}

func (sle Sle) getMinorsMatrix() Sle {
	size := len(sle.matrix)
	res := MakeMatrix(size, size)
	var channels []*chan []float64
	for i, l := range sle.matrix {
		ch := make(chan []float64)
		channels = append(channels, &ch)
		go func(i int, l []float64, ch chan []float64) {
			vector := make([]float64, len(l), len(l))
			for j := range l {
				det := sle.getMinor(i, j).determinant()
				sle.log(fmt.Sprintf("[DETERMINANT [%d, %d] = %f]\n", i, j, det))
				vector[j] = det
			}
			ch <- vector
			close(ch)
		}(i, l, ch)
	}
	for i := range res {
		res[i] = <-*channels[i]
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
	det := sle.determinant()
	if det == 0 {
		return sle.cloneSle(nil, nil), fmt.Errorf("Determinant of matrix is equal to 0")
	}
	res := sle.getMinorsMatrix().getAlgComplemetsMatr().transponate()
	res.matrix = res.matrix.Mult(1.0 / det)
	return res, nil
}

//Solve returns a slice of solutions for SLE.
func (sle Sle) Solve() ([]float64, error) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	sle.log(fmt.Sprintf("Num of threads set to : %d\n", runtime.NumCPU()))
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
	runtime.GOMAXPROCS(1)
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
