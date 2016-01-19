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
	getPayment(economy.MoneyReceiver, economy.Money) (error)
}

type loyalty int

type gangster struct {
	payer payer
	clock timedEvents.Clock
	loyalty loyalty
}

func (g *gangster) 	changePayer(p payer) {
	g.payer = p
	g.initPaymentTicker()
}

func (g *gangster) payment() economy.Money {
	return payment
}

func (g *gangster) decreaseLoyalty(l loyalty) {
	g.loyalty = g.loyalty - loyalty(math.Abs(float64(l)))
}

func (g *gangster) initPaymentTicker() {
	g.clock.AddTicker(daysForPayment, func() {
		err := g.payer.getPayment(g, g.payment())
		if err != nil {
			g.decreaseLoyalty(loyaltyLossForNotPayed)
		}
	})
}

func (g *gangster) ReceiveMoney(m economy.Money, w *economy.Wallet) {
	economy.Consume(w, m)
}