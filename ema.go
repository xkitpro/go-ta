package ta

import (
	"math"
)

type ema struct {
	div  float64
	prev float64
}

func NewEMA(period uint) *ema {
	return &ema{
		div:  2. / (float64(period) + 1.),
		prev: math.NaN(),
	}
}

func (i *ema) Write(v float64) float64 {
	if math.IsNaN(i.prev) {
		i.prev = v
	} else {
		i.prev = i.prev + i.div*(v-i.prev)
	}

	return i.prev
}
