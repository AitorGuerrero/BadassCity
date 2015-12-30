package BadassCity

type transaction struct {
	a float32
}

func (t transaction) amount() float32 {
	return t.a
}