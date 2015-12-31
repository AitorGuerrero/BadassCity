package gangsters
import (
	"github.com/AitorGuerrero/BadassCity/economy"
	"github.com/AitorGuerrero/BadassCity/timedEvents"
)

const (
	payment = economy.Unit * 10
	daysForPayment = timedEvents.Week
)

type payer interface {
	getPayment(economy.Money) (error, economy.Transaction)
}

type Gangster struct {
	payer payer
}

func (g *Gangster) changePayer(p payer) {
	g.payer = p
	timedEvents.InitTicker(daysForPayment, func() {
		p.getPayment(g.payment())
	})
}

func (g *Gangster) payment() economy.Money {
	return payment
}