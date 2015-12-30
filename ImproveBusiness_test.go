package BadassCity

import (
	t "testing"
	"github.com/AitorGuerrero/BadassCity/economy"
)

func TestGivenALocalWithNotBusinessShouldThrow(t *t.T) {
	l := makeFakeLocal()
	l.hasBusiness = false
	err := l.ImproveBusiness()

	if _, ok := err.(localDoesNotHaveABusiness); !ok {
		t.Error("should throw a localDoesNotHaveABusiness. Thrown: ", err)
	}
}

func TestGivenATotallyImprovedBusinessShouldThrow(t *t.T) {
	l := makeFakeLocal()
	l.business.level = businessLevel(2)
	err := l.ImproveBusiness()
	if _, ok := err.(totallyImprovedBusiness); !ok {
		t.Error("should throw a localDoesNotHaveABusiness. Thrown: ", err)
	}
}

func TestGivenAOwnerWithNotEnoughMoneyForImproveShouldThrow(t *t.T) {
	l := makeFakeLocal()
	l.owner.wallet.transactions = []transaction{transaction{economy.Money(5)}}
	err := l.ImproveBusiness()
	if _, ok := err.(notEnoughMoney); !ok {
		t.Error("should throw a notEnoughMoney. Thrown: ", err)
	}
}

func TestOwnerShouldSpentMoney(t *t.T) {
	l := makeFakeLocal()
	l.ImproveBusiness()
	if l.owner.wallet.totalAmount() != 0 {
		t.Error("Owner should not have money, has:", l.owner.wallet.totalAmount())
	}
}

func TestShouldImproveTheBusiness(t *t.T) {
	l := makeFakeLocal()
	err := l.ImproveBusiness()
	if err != nil {
		t.Error(err)
	}
	if l.business.level != 2 {
		t.Error("Should have level 2, has: ", l.business.level)
	}
}

func makeFakeLocal() local {
	return local{
		room: localRoom(2),
		business: makeFakeBusiness(),
		hasBusiness: true,
		owner: &merchant{
			wallet: wallet{
				transactions: []transaction{transaction{economy.Money(20)}},
			},
		},
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