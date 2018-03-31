package ta

import (
	"testing"
)

func TestBuffer(t *testing.T) {
	type testpair struct {
		v float64
		s float64
	}

	var tests = []testpair{
		{1.0, 1.0},
		{2.0, 3.0},
		{3.0, 6.0},
		{4.0, 10.0},
		{5.0, 15.0},
		{6.0, 20.0},
		{7.0, 25.0},
	}

	buf := NewBuffer(5)

	for _, pair := range tests {
		buf.Add(pair.v)
		if buf.Sum() != pair.s {
			t.Error(
				"For", pair.v,
				"expected", pair.s,
				"got", buf.Sum(),
			)
		}
	}
}

func TestGetBuffer(t *testing.T) {
	in := []float64{1., 2., 3., 4., 5.}

	buf := NewBuffer(5)

	for _, v := range in {
		buf.Add(v)
	}
}
