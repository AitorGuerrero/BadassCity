package BadassCity

import t "testing"

func TestGivenAPlayerWithNotEnoughMoneyShouldThrow(t *t.T) {
	p := givenAPlayerWithNotEnoughMoney()
	l := givenALocalOwnedByNpcOwner()
	err := p.BuyLocal(&l)
	if (err == nil) {
		t.Error("Should return error")
	}
	if l.owner() == &p {
		t.Error("Should not change the owner")
	}
}

func TestShouldChangeLocalOwner(t *t.T) {
	l := givenALocalOwnedByNpcOwner()
	p := givenAPlayerWithEnoughMoney()

	p.BuyLocal(&l)
	if l.owner() != &p {
		t.Error("Should change local owner")
	}
}

func TestPlayerShouldHaveANegativeTransaction(t *t.T) {
	l := givenALocalOwnedByNpcOwner()
	p := givenAPlayerWithEnoughMoney()

	p.BuyLocal(&l)
	if (p.w.transactions[1].a != -30) {
		t.Error("Should have a negative transaction")
	}
}

func TestOwnerShouldHaveATransaction(t *t.T) {
	l := givenALocalOwnedByNpcOwner()
	p := givenAPlayerWithEnoughMoney()
	o := l.o

	p.BuyLocal(&l)
	a := o.(*npcLocalOwner)
	if (a.w.transactions[0].a != 30) {
		t.Error("Should have a transaction")
	}
}

func givenALocalOwnedByNpcOwner() (l local) {
	var localPrice float32 = 30
	l = local{p: localPrice}
	l.o = &npcLocalOwner{}

	return
}

func givenAPlayerWithEnoughMoney() player {
	var enoughMoney float32 = 30
	return givenAPlayerWithMoney(enoughMoney)
}

func givenAPlayerWithNotEnoughMoney() player {
	var notEnoughMoney float32 = 10
	return givenAPlayerWithMoney(notEnoughMoney)
}

func givenAPlayerWithMoney(m float32) player {
	p := player{}
	p.w.transactions = append(p.w.transactions, transaction{m})
	return p
}