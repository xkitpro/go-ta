package ta

import (
	"math"
	"testing"
)

func TestWilders(t *testing.T) {
	nan := math.NaN()
	series := []float64{81.59, 81.06, 82.87, 83., 83.61, 83.15, 82.84}
	expected := []float64{nan, nan, nan, nan, 82.43, 82.57, 82.63, 82.90}

	wilders := NewWilders(5)
	for i, v := range series {
		g := wilders.Write(v)
		if !isEqualFloat(g, expected[i], 0.01) {
			t.Error(
				"For", v,
				"expected", expected[i],
				"got", g,
			)
		}
	}

}
