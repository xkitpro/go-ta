package ta

import (
	"fmt"
	"math"
	"testing"
)

func TestCMO(t *testing.T) {
	type row struct {
		v float64
		i float64
	}

	nan := math.NaN()

	data := []row{
		{81.59, nan},
		{81.06, nan},
		{82.87, nan},
		{83.00, nan},
		{83.61, nan},
		{83.15, 44.07},
		{82.84, 53.61},
		{83.99, 42.11},
		{84.55, 50.16},
		{84.36, 28.09},
		{85.53, 70.41},
		{86.54, 90.69},
		{86.89, 88.41},
		{87.77, 89.44},
		{87.29, 75.32},
	}

	i := NewCMO(5)
	for _, r := range data {
		cmo := i.Write(r.v)
		if fmt.Sprintf("%.2f", cmo) != fmt.Sprintf("%.2f", r.i) {
			t.Error(
				"For", r.v,
				"expected", r.i,
				"got", cmo,
			)
		}
	}
}
