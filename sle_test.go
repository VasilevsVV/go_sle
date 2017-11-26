package main

import (
	"testing"

	"./sle"
)

func getMatrixes() []sle.MatrSlice {
	return []sle.MatrSlice{[][]float64{{1, 2, -3, 10}, {4, -5, 6, 20}, {-7, 8, 9, 30}},
		[][]float64{
			{4, 7, -3, -11, 7, 36},
			{-2, 11, 3, 21, 0, -11},
			{0, 0, 4, -16, -5, 27},
			{26, 0, -7, -1, 18, -5},
			{5, 8, 21, -5, 3, 13}}}
}

func TestSle(t *testing.T) {
	ins := getMatrixes()
	results := [][]float64{
		[]float64{7.75, 6.5, float64(43.0 / 12.0)},
		[]float64{-40237.0 / 164073.0, 58706.0 / 23439.0, -35953.0 / 54691.0,
			-289622.0 / 164073.0, -45491.0 / 164073},
	}
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
