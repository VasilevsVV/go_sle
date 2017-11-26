package main

import (
	"testing"

	"./sle"
)

func TestSle1(t *testing.T) {
	cases := []struct {
		in   sle.MatrSlice
		want []float64
	}{
		{[][]float64{{1, 2, -3, 10}, {4, -5, 6, 20}, {-7, 8, 9, 30}},
			[]float64{7.75, 6.5, float64(43.0 / 12.0)}},
		{[][]float64{{4, 7, -3, -11, 7, 36},
			{-2, 11, 3, 21, 0, -11},
			{0, 0, 4, -16, -5, 27},
			{26, 0, -7, -1, 18, -5},
			{5, 8, 21, -5, 3, 13}},
			[]float64{-40237.0 / 164073.0, 58706.0 / 23439.0, -35953.0 / 54691.0,
				-289622.0 / 164073.0, -45491.0 / 164073}},
	}
	for _, c := range cases {
		Sle, err := sle.CreateSle(c.in)
		if err != nil {
			t.Error(err)
		}
		res, err := Sle.Solve()
		if err != nil {
			t.Error(err)
		}
		if !sle.CompareSlices(res, c.want) {
			t.Errorf("Want : %f\nAnd res : %f\n Are not equal",
				c.want, res)
		}
	}
}
