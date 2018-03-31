package ta

import (
	"math"
	"testing"
)

func TestCVI(t *testing.T) {
	nan := math.NaN()

	in := [][2]float64{
		{82.15, 81.29},
		{81.89, 80.64},
		{83.03, 81.31},
		{83.30, 82.65},
		{83.85, 83.07},
		{83.90, 83.11},
		{83.33, 82.49},
		{84.30, 82.30},
		{84.84, 84.15},
		{85.00, 84.11},
		{85.90, 84.03},
		{86.58, 85.39},
		{86.98, 85.76},
	}
	out := []float64{nan, nan, nan, nan, nan, 4.46, -11.22, 1.56, 2.52, 5.68, 44.09, 43.32, -0.49}

	cvi := NewCVI(5)
	for i, v := range in {
		g := cvi.Write(v[0], v[1])
		if !isEqualFloat(out[i], g, 0.01) {
			t.Error(
				"For", v,
				"expected", out[i],
				"got", g,
			)
		}
	}

}
