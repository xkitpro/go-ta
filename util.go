package ta

type Buffer struct {
	size   uint
	pushes uint
	index  uint
	sum    float64
	vals   []float64
}

func NewBuffer(size uint) *Buffer {
	return &Buffer{
		size: size,
		vals: make([]float64, size),
	}
}

func (b *Buffer) Add(v float64) {
	if b.pushes >= b.size {
		b.sum -= b.vals[b.index]
	}

	b.sum += v
	b.vals[b.index] = v
	b.pushes++
	b.index++
	if b.index >= b.size {
		b.index = 0
	}
}

func (b *Buffer) Sum() float64 {
	return b.sum
}

func (b *Buffer) Max() float64 {
	m := b.vals[0]

	for _, v := range b.vals[1:] {
		if v > m {
			m = v
		}
	}

	return m
}

func (b *Buffer) Min() float64 {
	m := b.vals[0]

	for _, v := range b.vals[1:] {
		if v < m {
			m = v
		}
	}

	return m
}

func (b *Buffer) Get(i uint) float64 {
	return b.vals[(b.index+b.size-1-i)%b.size]
}

func TypicalPrice(v OHLCV) float64 {
	return (v.High + v.Low + v.Close) / 3.0
}
