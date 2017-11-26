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
			[]float64{7.75, 6.5, float64(43 / 12)}},
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
		if res != c.want {
			t.Errorf("Want : [elem0 , elem1 ...]\nAnd res : [elem0, elem1 ...]\n Are not equal",
				c.want, res)
		}
	}
}
