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

func NewPlayer() player {
	return player{}
}

func (p *player) BuyLocal(l *local) (err error) {
	err = trespassTransaction(l.price(), p, l.o)
	if (err != nil) {
		return
	}
 	l.changeOwner(p)

	return
}

func (pl player) hasEnoughMoney(a float32) bool {
	return pl.w.totalAmount() >= a
}

func (pl *player) getTransaction(t transaction) {
	pl.w.transactions = append(pl.w.transactions, t)
}

func (pl *player) giveTransaction(a float32) (err error, t transaction) {
	if !pl.hasEnoughMoney(a) {
		err = errors.New("Player does not have enough money")
		return
	}
	pl.getTransaction(transaction{-a})
	t = transaction{a}
	return
}