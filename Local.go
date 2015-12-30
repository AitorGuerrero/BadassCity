package BadassCity

type localRoom int

type localOwner struct {
	merchant
}

type localDoesNotHaveABusiness struct{}
func (localDoesNotHaveABusiness) Error() string {
	return "Local does not have a business"
}

type local struct {
	p money
	o *merchant
	business business
	room localRoom
	hasBusiness bool
}

func (l local) price() money {
	return l.p
}

func (l local) owner() *merchant {
	return l.o
}

func (l local) hasEnoughRoom(r localRoom) bool {
	return l.room < r
}

func (l local) priceForStartABusiness(b business) money {
	return b.model.priceForStartPerRoom.multiply(float32(l.room))
}

func (l local) priceForImprove() money {
	return l.business.priceForImprove()
}

func (l *local) changeOwner(o *merchant) {
	l.o = o
}

func (l *local) StartABusiness(b business) error {
	if l.hasEnoughRoom(b.model.neededRoom) {
		return notEnoughRoom{}
	}
	if err, _ := l.o.giveTransaction(l.priceForStartABusiness(b)); err != nil {
		return notEnoughMoney{}
	}
	l.business = b
	l.initPaymentTicker()

	return nil
}

func (l *local) collectBenefits() {
	l.o.getTransaction(transaction{l.business.benefits()})
}

func (l *local) initPaymentTicker() {
	initTicker(localPaymentDuration, func() {
		l.collectBenefits()
	})
}

func (l *local) ImproveBusiness() error {
	if (!l.hasBusiness) {
		return localDoesNotHaveABusiness{}
	}
	if l.business.isTotallyImproved() {
		return totallyImprovedBusiness{}
	}
	err, _ := l.o.giveTransaction(l.priceForImprove())
	if (err != nil) {
		return notEnoughMoney{}
	}
	l.business.improve()

	return nil
}