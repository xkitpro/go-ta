package ta

import (
	"math"
)

type bbands struct {
	scale     float64
	middle    *sma
	deviation *StdDev
}

func NewBBands(period uint, scale float64) *bbands {
	return &bbands{
		scale:     scale,
		middle:    NewSMA(period),
		deviation: NewStdDev(period),
	}
}

func (i *bbands) Write(v float64) (float64, float64, float64) {
	avg := i.middle.Write(v)
	dev := i.deviation.Write(v)
	if math.IsNaN(dev) {
		nan := math.NaN()
		return nan, nan, nan
	}

	return avg - dev*i.scale,
		avg,
		avg + dev*i.scale
}
