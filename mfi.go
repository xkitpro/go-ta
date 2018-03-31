package ta

import (
	"math"
)

type mfi struct {
	period uint
	prev   float64
	up     *Buffer
	down   *Buffer
}

func NewMFI(period uint) *mfi {
	return &mfi{
		period: period,
		up:     NewBuffer(period),
		down:   NewBuffer(period),
	}
}

func (i *mfi) Write(v OHLCV) float64 {
	typ := TypicalPrice(v)
	bar := typ * v.Volume

	if typ > i.prev {
		i.up.Add(bar)
		i.down.Add(0)
	} else if typ < i.prev {
		i.up.Add(0)
		i.down.Add(bar)
	} else {
		i.up.Add(0)
		i.down.Add(0)
	}

	i.prev = typ

	if i.up.pushes <= i.period {
		return math.NaN()
	}

	return i.up.Sum() / (i.up.Sum() + i.down.Sum()) * 100.
}
