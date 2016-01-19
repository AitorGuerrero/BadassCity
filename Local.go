package BadassCity
import "github.com/AitorGuerrero/BadassCity/economy"
import "github.com/AitorGuerrero/BadassCity/timedEvents"

const localPaymentDuration = timedEvents.Week

type localRoom int

type localOwner struct {
	economy.Merchant
}

type localDoesNotHaveABusiness struct{}
func (localDoesNotHaveABusiness) Error() string {
	return "Local does not have a business"
}

type local struct {
	price       economy.Money
	owner       *economy.Merchant
	business    business
	room        localRoom
	hasBusiness bool
	clock 		timedEvents.Clock
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

func (l *local) changeOwner(o *economy.Merchant) {
	l.owner = o
}

func (l *local) StartABusiness(b business) error {
	if l.hasEnoughRoom(b.model.neededRoom) {
		return notEnoughRoom{}
	}
	if err := economy.Consume(&l.owner.Wallet, l.priceForStartABusiness(b)); err != nil {
		return economy.NotEnoughMoney{}
	}
	l.business = b
	l.initPaymentTicker()

	return nil
}

func (l *local) collectBenefits() {
	economy.GenerateMoney(&l.owner.Wallet, l.business.benefits())
}

func (l *local) initPaymentTicker() {
	l.clock.AddTicker(localPaymentDuration, func() {
		l.collectBenefits()
	})
}

func (l *local) ImproveBusiness() (err error) {
	if err = l.canImproveBusiness(); err != nil {
		return
	}
	if err = economy.Consume(&l.owner.Wallet, l.priceForImprove()); err != nil {
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
	if !l.owner.Wallet.HasEnoughMoney(l.priceForImprove()) {
		return economy.NotEnoughMoney{}
	}

	return nil
}