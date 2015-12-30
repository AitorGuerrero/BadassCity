package BadassCity

type player struct {
	merchant
}

func NewPlayer() player {
	return player{}
}

func (p *player) BuyLocal(l *local) (err error) {
	err = trespassTransaction(l.price(), &(p.merchant), l.o)
	if (err != nil) {
		return
	}
	l.changeOwner(&(p.merchant))

	return
}