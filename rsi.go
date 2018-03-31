package ta

import (
	"math"
)

type rsi struct {
	prev float64
	up   *wilders
	down *wilders
}

func NewRSI(period uint) *rsi {
	return &rsi{
		prev: math.NaN(),
		up:   NewWilders(period),
		down: NewWilders(period),
	}
}

func (i *rsi) Write(v float64) float64 {
	if math.IsNaN(i.prev) {
		i.prev = v
		return math.NaN()
	}

	var (
		up, down   float64
		sup, sdown float64
	)

	if v > i.prev {
		up = v - i.prev
	} else if v < i.prev {
		down = i.prev - v
	}

	i.prev = v

	sup, sdown = i.up.Write(up), i.down.Write(down)

	return 100. * (sup / (sup + sdown))
}
