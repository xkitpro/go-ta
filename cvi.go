package ta

import (
	"math"
)

type cvi struct {
	ema *ema
	buf *Buffer
}

func NewCVI(period uint) *cvi {
	return &cvi{
		ema: NewEMA(period),
		buf: NewBuffer(period),
	}
}

func (i *cvi) Write(h, l float64) float64 {
	curr := i.ema.Write(h - l)
	prev := i.buf.Get(i.buf.size - 1)
	i.buf.Add(curr)

	if i.buf.pushes <= i.buf.size {
		return math.NaN()
	}

	return 100. * (curr/prev - 1.)
}
