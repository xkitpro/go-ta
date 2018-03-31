package ta

import (
	"math"
)

type cmo struct {
	period uint
	sup    float64
	sdown  float64
	prev   float64
	n      uint
	up     *Buffer
	down   *Buffer
}

func NewCMO(period uint) *cmo {
	return &cmo{
		up:   NewBuffer(period),
		down: NewBuffer(period),
	}
}

func (i *cmo) Write(v float64) float64 {
	if v > i.prev {
		i.up.Add(v - i.prev)
		i.down.Add(0.)
	} else if v < i.prev {
		i.up.Add(0.)
		i.down.Add(i.prev - v)
	} else {
		i.up.Add(0.)
		i.down.Add(0.)
	}

	i.prev = v

	if i.up.pushes <= i.up.size {
		return math.NaN()
	}

	return (i.up.sum - i.down.sum) / (i.up.sum + i.down.sum) * 100
}
