package ta

import (
	"math"
	"testing"
)

func isEqualFloat(a, b, e float64) bool {
	if a == b {
		return true
	}

	if math.IsNaN(a) && math.IsNaN(b) {
		return true
	}

	return math.Abs(a-b) < e
}
func TestRSI(t *testing.T) {
	nan := math.NaN()
	series := []float64{81.59, 81.06, 82.87, 83.00, 83.61, 83.15, 82.84, 83.99}
	expected := []float64{nan, nan, nan, nan, nan, 72.03, 64.93, 75.94}
	rsi := NewRSI(5)
	for i, v := range series {
		g := rsi.Write(v)
		if !isEqualFloat(g, expected[i], 0.01) {
			t.Error(
				"For", v,
				"expected", expected[i],
				"got", g,
			)
		}
	}
}
