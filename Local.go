package BadassCity
import "github.com/AitorGuerrero/BadassCity/economy"

type localRoom int

type localOwner struct {
	merchant
}

type localDoesNotHaveABusiness struct{}
func (localDoesNotHaveABusiness) Error() string {
	return "Local does not have a business"
}

type local struct {
	price       economy.Money
	owner       *merchant
	business    business
	room        localRoom
	hasBusiness bool
}

func (l local) hasEnoughRoom(r localRoom) bool {
	return l.room < r
}

func (l local) priceForStartABusiness(b business) economy.Money {
	return b.model.priceForStartPerRoom.Multiply(float32(l.room))
}

func (l local) priceForImprove() economy.Money {
	return l.business.priceForImprove()
}

func (l *local) changeOwner(o *merchant) {
	l.owner = o
}

func (l *local) StartABusiness(b business) error {
	if l.hasEnoughRoom(b.model.neededRoom) {
		return notEnoughRoom{}
	}
	if err, _ := l.owner.giveTransaction(l.priceForStartABusiness(b)); err != nil {
		return notEnoughMoney{}
	}
	l.business = b
	l.initPaymentTicker()

	return nil
}

func (l *local) collectBenefits() {
	l.owner.getTransaction(transaction{l.business.benefits()})
}

func (l *local) initPaymentTicker() {
	initTicker(localPaymentDuration, func() {
		l.collectBenefits()
	})
}

func (l *local) ImproveBusiness() (err error) {
	if err = l.canImproveBusiness(); err != nil {
		return
	}
	if err, _ = l.owner.giveTransaction(l.priceForImprove()); err != nil {
		return
	}
	if err = l.business.improve(); err != nil {
		return
	}

	return
}

func (l local) canImproveBusiness() (error) {
	if !l.hasBusiness {
		return localDoesNotHaveABusiness{}
	}
	if l.business.isTotallyImproved() {
		return totallyImprovedBusiness{}
	}
	if !l.owner.hasEnoughMoney(l.priceForImprove()) {
		return notEnoughMoney{}
	}

	return nil
}