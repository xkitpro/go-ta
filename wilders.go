package ta

import (
	"math"
)

type wilders struct {
	period uint
	real   float64
	count  uint
	factor float64
}

func NewWilders(period uint) *wilders {
	return &wilders{
		period: period,
		factor: 1. / float64(period),
	}
}

func (i *wilders) Write(v float64) float64 {
	i.count++

	if i.count < i.period {
		i.real += v
		return math.NaN()
	} else if i.count == i.period {
		i.real = (i.real + v) * i.factor
		return i.real
	}

	i.real = (v-i.real)*i.factor + i.real
	return i.real
}
