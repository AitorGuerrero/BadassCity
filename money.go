package BadassCity

type money float32

func (m money) multiply(a float32) money {
	return money(float32(m) * a)
}
