package ta

import (
	"math"
)

type StdDev struct {
	period uint
	div    float64
	n      uint
	sum    float64
	sum2   float64
	prev   []float64
}

func NewStdDev(period uint) *StdDev {
	return &StdDev{
		period: period,
		div:    1. / float64(period),
		prev:   make([]float64, 1, period),
	}
}

func (i *StdDev) Write(v float64) float64 {
	i.n++

	i.sum += v
	i.sum2 += v * v

	if i.n < i.period {
		i.prev = append(i.prev, v)
		return math.NaN()
	}

	prev := i.prev[0]
	i.prev = i.prev[1:]
	i.prev = append(i.prev, v)

	i.sum -= prev
	i.sum2 -= prev * prev

	s2s2 := (i.sum2*i.div - (i.sum*i.div)*(i.sum*i.div))
	return math.Sqrt(s2s2)

}
