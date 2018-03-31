package ta

import (
	"fmt"
	"math"
	"testing"
)

func TestCCI(t *testing.T) {
	type row struct {
		v OHLCV
		i float64
	}

	nan := math.NaN()

	data := []row{
		{OHLCV{0, 82.15, 81.29, 81.59, 0}, nan},
		{OHLCV{0, 81.89, 80.64, 81.06, 0}, nan},
		{OHLCV{0, 83.03, 81.31, 82.87, 0}, nan},
		{OHLCV{0, 83.30, 82.65, 83.00, 0}, nan},
		{OHLCV{0, 83.85, 83.07, 83.61, 0}, 105.01},
		{OHLCV{0, 83.90, 83.11, 83.15, 0}, 64.24},
		{OHLCV{0, 83.33, 82.49, 82.84, 0}, -29.63},
		{OHLCV{0, 84.30, 82.30, 83.99, 0}, 69.54},
		{OHLCV{0, 84.84, 84.15, 84.55, 0}, 166.67},
		{OHLCV{0, 85.00, 84.11, 84.36, 0}, 82.02},
		{OHLCV{0, 85.90, 84.03, 85.53, 0}, 95.50},
		{OHLCV{0, 86.58, 85.39, 86.54, 0}, 130.91},
		{OHLCV{0, 86.98, 85.76, 86.89, 0}, 99.16},
		{OHLCV{0, 88.00, 87.17, 87.77, 0}, 116.34},
		{OHLCV{0, 87.87, 87.01, 87.29, 0}, 71.93},
	}

	i := NewCCI(5)
	for _, r := range data {
		p := TypicalPrice(r.v)
		cci := i.Write(p)
		if fmt.Sprintf("%.2f", cci) != fmt.Sprintf("%.2f", r.i) {
			t.Error(
				"For", r.v,
				"expected", r.i,
				"got", cci,
			)
		}
	}
}
