package BadassCity

import (
	t "testing"
	"github.com/AitorGuerrero/BadassCity/economy"
)

func TestGivenAPlayerWithNotEnoughMoneyShouldThrow(t *t.T) {
	p := givenAPlayerWithNotEnoughMoney()
	l := givenALocalOwnedByNpcOwner()
	err := p.BuyLocal(&l)
	if (err == nil) {
		t.Error("Should return error")
	}
	if l.owner == &(p.merchant) {
		t.Error("Should not change the owner")
	}
}

func TestShouldChangeLocalOwner(t *t.T) {
	l := givenALocalOwnedByNpcOwner()
	p := givenAPlayerWithEnoughMoney()

	p.BuyLocal(&l)
	if l.owner != &(p.merchant) {
		t.Error("Should change local owner")
	}
}

func TestPlayerShouldHaveANegativeTransaction(t *t.T) {
	l := givenALocalOwnedByNpcOwner()
	p := givenAPlayerWithEnoughMoney()

	p.BuyLocal(&l)
	if (p.merchant.wallet.transactions[1].a != -30) {
		t.Error("Should have a negative transaction")
	}
}

func TestOwnerShouldHaveATransaction(t *t.T) {
	l := givenALocalOwnedByNpcOwner()
	p := givenAPlayerWithEnoughMoney()
	o := l.owner

	p.BuyLocal(&l)
	if (o.wallet.transactions[0].a != 30) {
		t.Error("Should have a transaction")
	}
}

func givenALocalOwnedByNpcOwner() (l local) {
	localPrice := economy.Money(30)
	l = local{price: localPrice}
	l.owner = &merchant{}

	return
}

func givenAPlayerWithEnoughMoney() player {
	return givenAPlayerWithMoney(economy.Money(30))
}

func givenAPlayerWithNotEnoughMoney() player {
	return givenAPlayerWithMoney(economy.Money(10))
}

func givenAPlayerWithMoney(m economy.Money) player {
	p := player{}
	p.wallet.transactions = append(p.wallet.transactions, transaction{m})
	return p
}