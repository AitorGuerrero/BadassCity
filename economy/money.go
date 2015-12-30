package economy

type Money float32

func (m Money) Multiply(a float32) Money {
	return Money(float32(m) * a)
}
