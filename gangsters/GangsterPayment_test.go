package gangsters

import (
	t "testing"
	"github.com/AitorGuerrero/BadassCity/economy"
	"github.com/AitorGuerrero/BadassCity/timedEvents/turnsClock"
)

type TestPayerWithNoMoney struct {
	wallet economy.Wallet
}

func (p TestPayerWithNoMoney) getPayment(economy.MoneyReceiver, economy.Money) (error) {
	return economy.NotEnoughMoney{}
}

func TestWhenPayerHasNotMoneyGangsterShouldDecreaseHisLoyalty(t *t.T) {
	clock := &turnsClock.Clock{}
	g := &gangster{clock: clock}
	m := &market{
		clock: clock,
		gangsters: []*gangster{g},
	}
	err := m.HireGangster(TestPayerWithNoMoney{}, g)
	if (err != nil) {
		t.Error(err)
	}
	for i := 0; i < int(turnsClock.DaysToTurns(daysForPayment)); i++ {
		clock.Next()
	}
	if (g.loyalty != -1) {
		t.Error("Should decrease loyalty")
	}
}
