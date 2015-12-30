package BadassCity

type transaction struct {
	a money
}

func (t transaction) amount() money {
	return t.a
}

func trespassTransaction (amount money, giver, receiver *merchant) (err error) {
	err, t := giver.giveTransaction(amount)
	if err != nil {
		return
	}
	receiver.getTransaction(t)
	return
}