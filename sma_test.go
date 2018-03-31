package ta

import (
	"math"
	"testing"
)

func TestSMA(t *testing.T) {
	series := []float64{5, 6, 7, 8, 9}
	period := uint(3)
	nan := math.NaN()
	returns := []float64{nan, nan, 6, 7, 8}

	sma := NewSMA(period)
	for i, v := range series {
		a := sma.Write(v)
		if !isEqualFloat(a, returns[i], 0.01) {
			t.Error(
				"For", series,
				"expected", returns[i],
				"got", a,
			)
		}
	}
}
