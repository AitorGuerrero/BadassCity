package economy

type Money uint

const (
	Unit = Money(1)
)

func (m Money) Multiply(a float32) Money {
	return Money(float32(m) * a)
}
