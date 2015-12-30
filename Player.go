package BadassCity

type player struct {
	merchant
}

func NewPlayer() player {
	return player{}
}

func (p *player) BuyLocal(l *local) (err error) {
	err = p.merchant.wallet.TransferTo(&l.owner.wallet, l.price)
	if (err != nil) {
		return
	}
	l.changeOwner(&(p.merchant))

	return
}