package ta

import (
	"math"
	"testing"
)

func TestStoch(t *testing.T) {
	series := []OHLCV{
		{0, 82.15, 81.29, 81.59, 0},
		{0, 81.89, 80.64, 81.06, 0},
		{0, 83.03, 81.31, 82.87, 0},
		{0, 83.30, 82.65, 83.00, 0},
		{0, 83.85, 83.07, 83.61, 0},
		{0, 83.90, 83.11, 83.15, 0},
		{0, 83.33, 82.49, 82.84, 0},
		{0, 84.30, 82.30, 83.99, 0},
		{0, 84.84, 84.15, 84.55, 0},
		{0, 85.00, 84.11, 84.36, 0},
		{0, 85.90, 84.03, 85.53, 0},
		{0, 86.58, 85.39, 86.54, 0},
		{0, 86.98, 85.76, 86.89, 0},
		{0, 88.00, 87.17, 87.77, 0},
		{0, 87.87, 87.01, 87.29, 0},
	}

	nan := math.NaN()
	expected := [][2]float64{
		{nan, nan},
		{nan, nan},
		{nan, nan},
		{nan, nan},
		{nan, nan},
		{nan, nan},
		{nan, nan},
		{nan, nan},
		{77.39, 75.70},
		{83.13, 78.01},
		{84.87, 81.79},
		{88.36, 85.45},
		{95.25, 89.49},
		{96.74, 93.45},
		{91.09, 94.36},
	}

	stoch := NewStoch(5, 3, 3)
	for i, v := range series {
		k, d := stoch.Write(v)
		e := expected[i]
		if !isEqualFloat(k, e[0], 0.01) {
			t.Error(
				"For", v,
				"expected", e[0],
				"got", k,
			)
		}
		if !isEqualFloat(d, e[1], 0.01) {
			t.Error(
				"For", v,
				"expected", e[1],
				"got", d,
			)
		}
	}
}