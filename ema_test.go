package ta

import "testing"

func TestEMA(t *testing.T) {
	series := []float64{81.59, 81.06, 82.87, 83.00, 83.61, 83.15, 82.84, 83.99, 84.55}
	expected := []float64{81.59, 81.41, 81.90, 82.27, 82.71, 82.86, 82.85, 83.23, 83.67}
	ema := NewEMA(5)

	var g float64
	for i, v := range series {
		g = ema.Write(v)
		if !isEqualFloat(g, expected[i], 0.01) {
			t.Error(
				"For", v,
				"expected", expected[i],
				"got", g,
			)
		}
	}
}
