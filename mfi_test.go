package ta

import (
	"math"
	"testing"
)

func TestMFI(t *testing.T) {
	nan := math.NaN()

	type testpair struct {
		ohlcv OHLCV
		idx   float64
	}

	var tests = []testpair{
		{OHLCV{0, 82.15, 81.29, 81.59, 5653100}, nan},
		{OHLCV{0, 81.89, 80.64, 81.06, 6447400}, nan},
		{OHLCV{0, 83.03, 81.31, 82.87, 7690900}, nan},
		{OHLCV{0, 83.3, 82.65, 83, 3831400}, nan},
		{OHLCV{0, 83.85, 83.07, 83.61, 4455100}, nan},
		{OHLCV{0, 83.9, 83.11, 83.15, 3798000}, 61.17},
		{OHLCV{0, 83.33, 82.49, 82.84, 3936200}, 67.31},
		{OHLCV{0, 84.3, 82.3, 83.99, 4732000}, 62.80},
		{OHLCV{0, 84.84, 84.15, 84.55, 4841300}, 64.66},
		{OHLCV{0, 85, 84.11, 84.36, 3915300}, 45.24},
		{OHLCV{0, 85.9, 84.03, 85.53, 6830800}, 67.84},
		{OHLCV{0, 86.58, 85.39, 86.54, 6694100}, 85.58},
		{OHLCV{0, 86.98, 85.76, 86.89, 5293600}, 85.96},
		{OHLCV{0, 88, 87.17, 87.77, 7985800}, 87.50},
		{OHLCV{0, 87.87, 87.01, 87.29, 4807900}, 84.65},
	}

	i := NewMFI(5)

	var idx float64
	for _, pair := range tests {
		idx = i.Write(pair.ohlcv)
		if !isEqualFloat(pair.idx, idx, 0.01) {
			t.Error(
				"For", pair.ohlcv,
				"expected", pair.idx,
				"got", idx,
			)
		}
	}
}
