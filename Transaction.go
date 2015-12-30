package BadassCity

type transaction struct {
	a float32
}

func (t transaction) amount() float32 {
	return t.a
}

func trespassTransaction (amount float32, giver, receiver merchant) (err error) {
	err, t := giver.giveTransaction(amount)
	if err != nil {
		return
	}
	receiver.getTransaction(t)
	return
}