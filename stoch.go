package ta

import (
	"math"
)

type stoch struct {
	highs *Buffer
	lows  *Buffer
	k     *sma
	d     *sma
}

func NewStoch(kperiod, kslow, dperiod uint) *stoch {
	return &stoch{
		highs: NewBuffer(kperiod),
		lows:  NewBuffer(kperiod),
		k:     NewSMA(kslow),
		d:     NewSMA(dperiod),
	}
}

func (i *stoch) Write(v OHLCV) (float64, float64) {
	i.highs.Add(v.High)
	i.lows.Add(v.Low)

	if i.highs.pushes < i.highs.size {
		return math.NaN(), math.NaN()
	}

	max, min := i.highs.Max(), i.lows.Min()

	kfast := 100 * (v.Close - min) / (max - min)

	k := i.k.Write(kfast)
	if math.IsNaN(k) {
		return k, k
	}
	d := i.d.Write(k)
	if math.IsNaN(d) {
		return d, d
	}

	return k, d
}
