package ta

import (
	"fmt"
	"math"
	"testing"
)

func TestBBands(t *testing.T) {
	nan := math.NaN()

	series := []float64{81.59, 81.06, 82.87, 83.00, 83.61, 83.15, 82.84, 83.99, 84.55, 84.36, 85.53, 86.54, 86.89, 87.77, 87.29}
	scale := 2.0
	period := 5
	results := [][3]float64{
		{nan, nan, nan},
		{nan, nan, nan},
		{nan, nan, nan},
		{nan, nan, nan},
		{80.53, 82.43, 84.32},
		{80.99, 82.74, 84.49},
		{82.53, 83.09, 83.65},
		{82.47, 83.32, 84.16},
		{82.42, 83.63, 84.84},
		{82.44, 83.78, 85.12},
		{82.51, 84.25, 86.00},
		{83.14, 84.99, 86.85},
		{83.54, 85.57, 87.61},
		{83.87, 86.22, 88.57},
		{85.29, 86.80, 88.32},
	}

	bbands := NewBBands(uint(period), scale)
	for i, v := range series {
		u, m, l := bbands.Write(v)

		e := results[i]

		if fmt.Sprintf("%.2f", e[0]) != fmt.Sprintf("%.2f", u) ||
			fmt.Sprintf("%.2f", e[1]) != fmt.Sprintf("%.2f", m) ||
			fmt.Sprintf("%.2f", e[2]) != fmt.Sprintf("%.2f", l) {
			t.Error(
				"For", v,
				"expected", e,
				"got", u, m, l,
			)
		}
	}
}
