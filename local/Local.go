package local

import "github.com/AitorGuerrero/goConomy"
import "github.com/AitorGuerrero/BadassCity/timedEvents"

const localPaymentDuration = timedEvents.Week

type localRoom int

type local struct {
	goConomy.MoneyGenerator
	goConomy.MoneyDestroyer
	owner       Owner
	business    business
	hasBusiness bool
	room        localRoom
	clock       timedEvents.Clock
}

func (l *local) StartABusiness(bm businessModel) error {
	if !l.hasEnoughRoomForBusiness(bm) {
		return notEnoughRoom{}
	}
	if err := l.ConsumeMoney(l.owner, l.priceForStartABusiness(bm)); err != nil {
		return goConomy.NotEnoughMoney{}
	}
	l.business = business{
		model: bm,
		room: l.room,
		level: businessLevel(1),
	}
	l.hasBusiness = true
	l.initPaymentTicker()

	return nil
}

func (l *local) ImproveBusiness() (err error) {
	if err = l.canImproveBusiness(); err != nil {
		return
	}
	if err = l.ConsumeMoney(l.owner, l.priceForImprove()); err != nil {
		return
	}
	l.business.improve()

	return
}

func (l local) hasEnoughRoomForBusiness(bm businessModel) bool {
	return l.hasEnoughRoom(bm.neededRoom)
}

func (l local) hasEnoughRoom(r localRoom) bool {
	return l.room >= r
}

func (l local) priceForStartABusiness(bm businessModel) goConomy.Money {
	return bm.priceForStartPerRoom.Multiply(float32(l.room))
}

func (l local) priceForImprove() goConomy.Money {
	return l.business.priceForImprove()
}

func (l *local) collectBenefits() {
	l.GenerateMoney(l.owner, l.business.benefits())
}

func (l *local) initPaymentTicker() {
	l.clock.AddTicker(localPaymentDuration, func() {
		l.collectBenefits()
	})
}

func (l local) canImproveBusiness() error {
	if !l.hasBusiness {
		return localDoesNotHaveABusiness{}
	}
	if l.business.isTotallyImproved() {
		return totallyImprovedBusiness{}
	}
	return nil
}

func (l local) businessLevel() businessLevel {
	return l.business.level
}