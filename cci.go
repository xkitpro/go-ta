package ta

import "math"

type cci struct {
	period uint
	div    float64
	sum    *Buffer
}

func NewCCI(period uint) *cci {
	return &cci{
		period: period,
		div:    1. / float64(period),
		sum:    NewBuffer(period),
	}
}

func (i *cci) Write(v float64) float64 {
	i.sum.Add(v)

	if i.sum.pushes < i.period {
		return math.NaN()
	}

	avg := i.sum.Sum() * i.div

	acc := 0.
	for _, v := range i.sum.vals {
		acc += math.Abs(avg - v)
	}
	acc *= i.div

	idx := (v - avg) / acc / 0.015

	return idx
}
