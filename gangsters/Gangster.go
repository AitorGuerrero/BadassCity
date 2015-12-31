package gangsters
import (
	"github.com/AitorGuerrero/BadassCity/economy"
	"github.com/AitorGuerrero/BadassCity/timedEvents"
	"math"
)

const (
	payment = economy.Unit * 10
	daysForPayment = timedEvents.Week
	loyaltyLossForNotPayed = 1
)

type payer interface {
	getPayment(economy.Wallet, economy.Money) (error)
}

type loyalty int

type gangster struct {
	economy.MoneyReceiver
	payer payer
	clock timedEvents.Clock
	loyalty loyalty
}

func (g *gangster) changePayer(p payer) {
	g.payer = p
	g.clock.AddTicker(daysForPayment, func() {
		err, _ := p.getPayment(g.payment())
		if err != nil {
			g.decreaseLoyalty(loyaltyLossForNotPayed)
		}
	})
}

func (g *gangster) payment() economy.Money {
	return payment
}

func (g *gangster) decreaseLoyalty(l loyalty) {
	g.loyalty = g.loyalty - loyalty(math.Abs(float64(l)))
}