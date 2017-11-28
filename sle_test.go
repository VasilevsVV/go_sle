package main

import (
	"testing"

	"./sle"
)

func getMediumMatrix() sle.MatrSlice {
	return [][]float64{
		{12, 15, -3, 56, 46, 3, 56, -97, 6},
		{34, -8, 87, 5, 45, 6, 34, 6, -87},
		{43, 24, -9, 8, 36, 21, 34, 67, -9},
		{45, 6, 56, -76, 65, 42, 54, -8, -76},
		{56, 43, 76, 5, -8, 87, 53, 24, -87},
		{-56, 20, 43, 43, 34, -17, 72, 43, 23},
		{76, 46, -73, 35, 85, 84, 35, 89, -65},
		{45, 76, -73, 23, 16, 8, -43, 24, 75},
	}
}

func getMediumResult() []float64 {
	return []float64{-0.0854486486, 1.0893912016, -0.6491091389, -0.2816926027,
		-0.6185427379, -1.2671235885, 0.5395247648, -0.0675559173}
}

func getBigMatrix() sle.MatrSlice {
	return [][]float64{
		{12, 15, -3, 56, 46, 3, 56, -97, 6, 23, -9},
		{34, -8, 87, 5, 45, 6, 34, 6, -87, 87, 67},
		{43, 24, -9, 8, 36, 21, 34, 67, -9, -76, 8},
		{45, 6, 56, -76, 65, 42, 54, -8, -76, 53, 35},
		{56, 43, 76, 5, -8, 87, 53, 24, -87, 27, 28},
		{-56, 20, 43, 43, 34, -17, 72, 43, 23, -54, 28},
		{76, 46, -73, 35, 85, 84, 35, 89, -65, 18, -63},
		{45, 76, -73, 23, 16, 8, -43, 24, 75, 38, 29},
		{53, -75, 75, 34, 15, 7, -65, 7, 53, -76, 46},
		{-76, 5, 76, 55, 43, 65, 67, 87, 98, -86, 45},
	}
}

func getBigResult() []float64 {
	return []float64{1.2106107148, -0.5511810357, 0.5321367500, -0.2105863700, -0.7306853426,
		-0.7378783187, 1.2849789619, 0.6488343985, 1.1705003456, 0.7786916242}
}

func getBigTestMatrix() sle.MatrSlice {
	return [][]float64{
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 20},
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 30},
		{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 40},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 50},
		{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 60},
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 70},
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 80},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 90},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 100},
	}
}

func getBigTestResult() []float64 {
	return []float64{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
}

func getLargeMatrix() sle.MatrSlice {
	return [][]float64{
		{12, 15, -3, 56, 46, 3, 56, -97, 6, 87, 13, -9, 98},
		{34, -8, 87, 5, 45, 6, 34, 6, -87, 6, 54, 35, -7},
		{43, 24, -9, 8, 36, 21, 34, 67, -9, 8, 6, 4, 8},
		{45, 6, 56, -76, 65, 42, 54, -8, -76, 53, 5, 6, -43},
		{56, 43, 76, 5, -8, 87, 53, 24, -87, 54, 34, 3, 16},
		{-56, 20, 43, 43, 34, -17, 72, 43, 23, -6, 13, -32, 2},
		{76, 46, -73, 35, 85, 84, 35, 89, -65, 42, 56, 73, 91},
		{45, 76, -73, 23, 16, 8, -43, 24, 75, 27, -81, 30, 73},
		{4, 26, 34, 95, 27, 96, -7, -36, 75, 82, 93, 26, 75},
		{36, 8, 35, 75, -6, 49, 76, 56, 45, 75, -65, 65, 86},
		{56, 86, 45, 14, -63, 16, 22, 23, 25, 63, 76, -76, 56},
		{26, -65, 87, 4, -45, 35, 75, 83, 47, 29, 64, 35, -54}}
}

func getLargeResult() []float64 {
	return []float64{-0.3115411222, 0.0683765696, -0.3478120152,
		0.8801386556, -0.0703200598, -0.6768418131,
		-0.6187889943, 0.4469540338, -0.5692153340,
		1.5853090166, -0.0647483120, 0.0012496748}
}

func getMatrixes() []sle.MatrSlice {
	return []sle.MatrSlice{
		[][]float64{
			{1, 2, -3, 10},
			{4, -5, 6, 20},
			{-7, 8, 9, 30}},
		[][]float64{
			{4, 7, -3, -11, 7, 36},
			{-2, 11, 3, 21, 0, -11},
			{0, 0, 4, -16, -5, 27},
			{26, 0, -7, -1, 18, -5},
			{5, 8, 21, -5, 3, 13}}}
}

func getSqrMatrixes() []sle.MatrSlice {
	return []sle.MatrSlice{
		[][]float64{
			{1, 2, -3},
			{4, -5, 6},
			{-7, 8, 9}},
		[][]float64{
			{4, 7, -3, -11, 7},
			{-2, 11, 3, 21, 0},
			{0, 0, 4, -16, -5},
			{26, 0, -7, -1, 18},
			{5, 8, 21, -5, 3}}}
}

func TestSle(t *testing.T) {
	ins := getMatrixes()
	results := [][]float64{
		[]float64{7.75, 6.5, float64(43.0 / 12.0)},
		[]float64{-40237.0 / 164073.0, 58706.0 / 23439.0, -35953.0 / 54691.0,
			-289622.0 / 164073.0, -45491.0 / 164073},
	}
	//ins = append(ins, getBigTestMatrix())
	// results = append(results, getBigTestResult())

	ins = append(ins, getMediumMatrix())
	results = append(results, getMediumResult())

	ins = append(ins, getBigMatrix())
	results = append(results, getBigResult())

	ins = append(ins, getLargeMatrix())
	results = append(results, getLargeResult())

	for i, in := range ins {
		Sle, err := sle.CreateSle(in)
		if err != nil {
			t.Error(err)
		}
		res, err := Sle.Solve()
		if err != nil {
			t.Error(err)
		}
		if !sle.CompareSlices(res, results[i]) {
			t.Errorf("Want : %f\nAnd res : %f\n Are not equal",
				results[i], res)
		}
	}
}

func TestTransponate(t *testing.T) {
	ins := getMatrixes()
	results := []sle.MatrSlice{
		[][]float64{{1, 4, -7}, {2, -5, 8}, {-3, 6, 9}, {10, 20, 30}},
		[][]float64{
			{4, -2, 0, 26, 5},
			{7, 11, 0, 0, 8},
			{-3, 3, 4, -7, 21},
			{-11, 21, -16, -1, -5},
			{7, 0, -5, 18, 3},
			{36, -11, 27, -5, 13}},
	}
	for i, in := range ins {
		res, err := in.Transponate()
		if err != nil {
			t.Error(err)
		} else {
			if !res.EqualTo(results[i]) {
				t.Error("Results are not equal")
			}
		}
	}
}

func TestMinor(t *testing.T) {
	ins := getMatrixes()
	results := []sle.MatrSlice{
		[][]float64{{1, 2, 10}, {-7, 8, 30}},
		[][]float64{
			{4, 7, -11, 7, 36},
			{0, 0, -16, -5, 27},
			{26, 0, -1, 18, -5},
			{5, 8, -5, 3, 13}},
	}
	for i, in := range ins {
		if !in.GetMinor(1, 2).EqualTo(results[i]) {
			t.Errorf("Result response in not equal to expected matrix")
		}
	}
}

func TestDeterminant(t *testing.T) {
	ins := getSqrMatrixes()
	results := []float64{-240, 984438}
	for i, in := range ins {
		res, err := in.Determinant()
		if err != nil {
			t.Error(err)
		} else {
			if res != results[i] {
				t.Errorf("Result is not equal to expected")
			}
		}
	}
}
