package macd

import (
	"gitlab.com/xarbit/indicator/ema"
)

type MACD struct {
	fast   *ema.Ema
	slow   *ema.Ema
	signal *ema.Ema
}

func New(fast, slow, signal int) *MACD {
	a := &MACD{
		fast:   ema.New(fast),
		slow:   ema.New(slow),
		signal: ema.New(signal),
	}

	return a
}

func (a *MACD) Write(v float64) {
	a.fast.Write(v)
	a.slow.Write(v)

	if a.slow.Valid() {
		a.signal.Write(a.fast.Value() - a.slow.Value())
	}
}

func (a MACD) Value() (float64, float64, float64, float64) {
	return a.fast.Value(), a.slow.Value(), a.fast.Value() - a.slow.Value(), a.signal.Value()
}

func (a MACD) Valid() bool {
	return a.fast.Valid() && a.slow.Valid() && a.signal.Valid()
}
