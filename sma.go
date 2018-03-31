package ta

import (
	"math"
)

type sma struct {
	div float64
	buf *Buffer
}

func NewSMA(period uint) *sma {
	return &sma{
		div: 1. / float64(period),
		buf: NewBuffer(period),
	}
}

func (i *sma) Write(v float64) float64 {
	i.buf.Add(v)

	if i.buf.pushes < i.buf.size {
		return math.NaN()
	}

	return i.buf.sum * i.div
}
