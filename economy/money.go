package economy

type Money int

const (
	Unit = Money(1)
)

func (m Money) Multiply(a float32) Money {
	return Money(float32(m) * a)
}
