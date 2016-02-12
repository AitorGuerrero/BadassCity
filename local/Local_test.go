package local

import (
	t "testing"
	"github.com/AitorGuerrero/BadassCity/timedEvents/turnsClock"
	"github.com/AitorGuerrero/goConomy"
	"github.com/AitorGuerrero/BadassCity/timedEvents"
)

type testOwner struct {
	goConomy.Merchant
}

func makeLocal() local {
	clock := turnsClock.Clock{}
	return local{
		clock: &clock,
		owner: &testOwner{},
		room: localRoom(2),
	}
}

func TestGivenALocalWithNotEnoughRoomWhenAskedToStartABussinessShouldReturnError(t *t.T) {
	l := makeLocal()
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

func TestWhenAskedToStartABussinessThenLocalBusinessShouldHaveLevel1(t *t.T) {
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
	if l.businessLevel() != businessLevel(1) {
		t.Error("Should have level 1. Has level", l.businessLevel())
	}
}

func TestGivenALocalWithNoBusinessWhenAskedToImproveThenReturnsError(t *t.T) {
	local := makeLocal()
	err := local.ImproveBusiness()
	if _, ok := err.(localDoesNotHaveABusiness); !ok {
		t.Error("Should return local does not have a business error")
	}
}

func TestGivenTotallyImprovedLocalWhenAskedToImproveThenReturnsError(t *t.T) {
	local := local{
		business: business{
			model: businessModel{
				maxLevel: businessLevel(2),
			},
			level: businessLevel(2),
		},
		hasBusiness: true,
	}
	err := local.ImproveBusiness()
	if _, ok := err.(totallyImprovedBusiness); !ok {
		t.Error("Should return totally improved business error. Returns: ", err)
	}
}

func TestWhenOwnerHasNotEnoughMoneyShouldReturnError(t *t.T) {
	o := testOwner{}
	mg := goConomy.MoneyGenerator{}
	mg.GenerateMoney(&o, goConomy.Money(9))
	local := local{
		owner: &o,
		hasBusiness: true,
		business: business {
			room: localRoom(2),
			model: businessModel{
				maxLevel: businessLevel(2),
				pricesForImprovePerRoom: map[businessLevel]goConomy.Money{
					businessLevel(2): goConomy.Money(5),
				},
			},
			level: businessLevel(1),
		},
	}
	err := local.ImproveBusiness()
	if _, ok := err.(goConomy.NotEnoughMoney); !ok {
		t.Error("Should return not enough money error. Returns: ", err)
	}
}

func TestShouldImproveTheBusiness(t *t.T) {
	o := testOwner{}
	mg := goConomy.MoneyGenerator{}
	mg.GenerateMoney(&o, goConomy.Money(10))
	local := local{
		owner: &o,
		hasBusiness: true,
		business: business {
			room: localRoom(2),
			model: businessModel{
				maxLevel: businessLevel(2),
				pricesForImprovePerRoom: map[businessLevel]goConomy.Money{
					businessLevel(2): goConomy.Money(5),
				},
			},
			level: businessLevel(1),
		},
	}
	err := local.ImproveBusiness()
	if nil != err {
		t.Error(err)
	}
}

func TestStartingABusinessAndPassedAWeekShouldReturnBenefitsPassingTheTime(t *t.T) {
	clock := turnsClock.Clock{}
	o := testOwner{}
	bm := businessModel{
		maxLevel: businessLevel(1),
		pricesForImprovePerRoom: map[businessLevel]goConomy.Money{
			businessLevel(2): goConomy.Money(1),
		},
		revenueByLevel: map[businessLevel]goConomy.Money{
			businessLevel(1): goConomy.Money(2),
		},
	}
	local := local{
		owner: &o,
		clock: &clock,
		room: localRoom(1),
	}
	err := local.StartABusiness(bm)
	if nil != err {
		t.Error(err)
	}
	clock.Advance(timedEvents.Week)
	if o.TotalMoneyAmount() != 2 {
		t.Error("Should have 2 money. Have", o.TotalMoneyAmount())
	}
}