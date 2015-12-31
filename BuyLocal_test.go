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
	if l.owner == &(p.Merchant) {
		t.Error("Should not change the owner")
	}
}

func TestShouldChangeLocalOwner(t *t.T) {
	l := givenALocalOwnedByNpcOwner()
	p := givenAPlayerWithEnoughMoney()

	p.BuyLocal(&l)
	if l.owner != &(p.Merchant) {
		t.Error("Should change local owner")
	}
}

func TestPlayerShouldHaveANegativeTransaction(t *t.T) {
	l := givenALocalOwnedByNpcOwner()
	p := givenAPlayerWithEnoughMoney()

	p.BuyLocal(&l)
	if (p.Wallet.TotalAmount() != 0) {
		t.Error("Should have a negative transaction")
	}
}

func TestOwnerShouldHaveATransaction(t *t.T) {
	l := givenALocalOwnedByNpcOwner()
	p := givenAPlayerWithEnoughMoney()
	o := l.owner

	p.BuyLocal(&l)
	if (o.Wallet.TotalAmount() != 30) {
		t.Error("Should have a transaction")
	}
}

func givenALocalOwnedByNpcOwner() (l local) {
	l = local{
		price: economy.Money(30),
		owner: &economy.Merchant{},
	}
	return
}

func givenAPlayerWithEnoughMoney() player {
	return givenAPlayerWithMoney(economy.Money(30))
}

func givenAPlayerWithNotEnoughMoney() player {
	return givenAPlayerWithMoney(economy.Money(10))
}

func givenAPlayerWithMoney(m economy.Money) player {
	mg := economy.MoneyGenerator{}
	p := player{}
	mg.GenerateMoney(&p.Wallet, m)
	return p
}