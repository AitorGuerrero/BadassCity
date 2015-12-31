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
	getPayment(economy.Money) (error, economy.Transaction)
}

type loyalty int

type Gangster struct {
	payer payer
	clock timedEvents.Clock
	loyalty loyalty
}

func (g *Gangster) changePayer(p payer) {
	g.payer = p
	g.clock.AddTicker(daysForPayment, func() {
		err, _ := p.getPayment(g.payment())
		if err != nil {
			g.decreaseLoyalty(loyaltyLossForNotPayed)
		}
	})
}

func (g *Gangster) payment() economy.Money {
	return payment
}

func (g *Gangster) decreaseLoyalty(l loyalty) {
	g.loyalty = g.loyalty - loyalty(math.Abs(float64(l)))
}