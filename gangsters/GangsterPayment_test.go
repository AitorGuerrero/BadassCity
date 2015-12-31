package gangsters

import (
	t "testing"
	"github.com/AitorGuerrero/BadassCity/economy"
	"errors"
	"github.com/AitorGuerrero/BadassCity/timedEvents/turnsClock"
)

type TestPayerWithNoMoney struct {}
func (p TestPayerWithNoMoney) getPayment(economy.Money) (error, economy.Transaction) {
	return errors.New("Not enoguh money"), economy.Transaction{}
}

func TestWhenPayerHasNotMoneyGangsterShouldDecreaseHisLoyalty(t *t.T) {
	clock := &turnsClock.Clock{}
	gangster := &Gangster{clock: clock}
	m := &market{
		clock: clock,
		gangsters: []*Gangster{gangster},
	}
	err := m.HireGangster(TestPayerWithNoMoney{}, gangster)
	if (err != nil) {
		t.Error(err)
	}
	for i := 0; i < int(turnsClock.DaysToTurns(daysForPayment)); i++ {
		clock.Next()
	}
	if (gangster.loyalty != -1) {
		t.Error("Should decrease loyalty")
	}
}
