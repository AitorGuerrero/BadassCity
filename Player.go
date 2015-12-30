package BadassCity

import (
	"errors"
)

type playerId string

type player struct {
	i playerId
	w wallet
	basePayment float32
	ownedLocals []local
}

func (p *player) BuyLocal(l *local) (err error) {
	price := l.price()
	if err = p.spend(price); err != nil {
		return
	}
	l.o.giveTransaction(transaction{price})
 	l.changeOwner(p)

	return
}

func (pl player) hasEnoughMoney(a float32) bool {
	return pl.w.totalAmount() >= a
}

func (pl *player) spend (a float32) (err error) {
	if !pl.hasEnoughMoney(a) {
		err = errors.New("Player does not have enough money")
		return
	}
	pl.giveTransaction(transaction{-a})
	return
}