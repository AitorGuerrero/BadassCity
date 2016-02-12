package local

import (
	t "testing"
	"github.com/AitorGuerrero/BadassCity/timedEvents/turnsClock"
	"github.com/AitorGuerrero/goConomy"
)

type testOwner struct {
	goConomy.Merchant
}

func TestGivenALocalWithNotEnoughRoomWhenAskedToStartABussinessShouldReturnError(t *t.T) {
	clock := turnsClock.Clock{}
	l := local{
		clock: &clock,
		owner: &testOwner{},
		room: localRoom(2),
	}
	bm := businessModel{
		neededRoom: localRoom(3),
	}
	err := l.StartABusiness(bm)
	if _, ok := err.(notEnoughRoom); !ok {
		t.Error("Should return not enoguh room error")
	}
}

func TestGivenAOwnerWithNotEnoughMoneyWhenAskedToStartABussinessShouldReturnError(t *t.T) {
	clock := turnsClock.Clock{}
	o := testOwner{}
	mg := goConomy.MoneyGenerator{}
	mg.GenerateMoney(&o, goConomy.Money(24))
	l := local{
		clock: &clock,
		owner: &o,
		room: localRoom(5),
	}
	bm := businessModel{
		priceForStartPerRoom: goConomy.Money(5),
		neededRoom: localRoom(3),
	}
	err := l.StartABusiness(bm)
	if _, ok := err.(goConomy.NotEnoughMoney); !ok {
		t.Error("Should return not enoguh money error")
	}
}

func TestWhenAskedToStartABussinessThenStartsABussinessInTheLocal(t *t.T) {
	clock := turnsClock.Clock{}
	o := testOwner{}
	mg := goConomy.MoneyGenerator{}
	mg.GenerateMoney(&o, goConomy.Money(25))
	l := local{
		clock: &clock,
		owner: &o,
		room: localRoom(5),
	}
	bm := businessModel{
		priceForStartPerRoom: goConomy.Money(5),
		neededRoom: localRoom(3),
	}
	err := l.StartABusiness(bm)
	if (nil != err) {
		t.Error(err)
	}
}