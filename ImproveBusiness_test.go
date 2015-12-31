package BadassCity

import (
	t "testing"
	"github.com/AitorGuerrero/BadassCity/economy"
	"github.com/AitorGuerrero/BadassCity/timedEvents"
	"github.com/AitorGuerrero/BadassCity/timedEvents/turnsClock"
)

func TestGivenALocalWithNotBusinessShouldThrow(t *t.T) {
	c := &turnsClock.Clock{}
	l := makeFakeLocal(c)
	l.hasBusiness = false
	err := l.ImproveBusiness()

	if _, ok := err.(localDoesNotHaveABusiness); !ok {
		t.Error("should throw a localDoesNotHaveABusiness. Thrown: ", err)
	}
}

func TestGivenATotallyImprovedBusinessShouldThrow(t *t.T) {
	c := &turnsClock.Clock{}
	l := makeFakeLocal(c)
	l.business.level = businessLevel(2)
	err := l.ImproveBusiness()
	if _, ok := err.(totallyImprovedBusiness); !ok {
		t.Error("should throw a localDoesNotHaveABusiness. Thrown: ", err)
	}
}

func TestGivenAOwnerWithNotEnoughMoneyForImproveShouldThrow(t *t.T) {
	mg := economy.MoneyGenerator{}
	c := &turnsClock.Clock{}
	l := makeFakeLocal(c)
	mg.GenerateMoney(&l.owner.Wallet, economy.Money(-5))
	err := l.ImproveBusiness()
	if _, ok := err.(economy.NotEnoughMoney); !ok {
		t.Error("should throw a notEnoughMoney. Thrown: ", err)
	}
}

func TestOwnerShouldSpentMoney(t *t.T) {
	c := &turnsClock.Clock{}
	l := makeFakeLocal(c)
	l.ImproveBusiness()
	if l.owner.Wallet.TotalAmount() != 0 {
		t.Error("Owner should not have money, has:", l.owner.Wallet.TotalAmount())
	}
}

func TestShouldImproveTheBusiness(t *t.T) {
	c := &turnsClock.Clock{}
	l := makeFakeLocal(c)
	err := l.ImproveBusiness()
	if err != nil {
		t.Error(err)
	}
	if l.business.level != 2 {
		t.Error("Should have level 2, has: ", l.business.level)
	}
}

func makeFakeLocal(c timedEvents.Clock) local {
	w := economy.Wallet{}
	mg := economy.MoneyGenerator{}
	mg.GenerateMoney(&w, economy.Money(20))
	return local{
		room: localRoom(2),
		business: makeFakeBusiness(),
		hasBusiness: true,
		owner: &economy.Merchant{
			Wallet: w,
		},
		clock: c,
	}
}

func makeFakeBusiness() business {
	pricesForImprovePerRoom := make(map[businessLevel]economy.Money)
	pricesForImprovePerRoom[businessLevel(1)] = 10
	return business{
		room: localRoom(2),
		level: businessLevel(1),
		model: businessModel{
			neededRoom: localRoom(2),
			pricesForImprovePerRoom: pricesForImprovePerRoom,
			maxLevel: businessLevel(2),
		},
	}
}