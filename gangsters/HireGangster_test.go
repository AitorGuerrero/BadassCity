package gangsters
import (
	"testing"
	"github.com/AitorGuerrero/BadassCity/economy"
	"github.com/AitorGuerrero/BadassCity/timedEvents/turnsClock"
)

type TestPayer struct {}
func (p TestPayer) getPayment(economy.Money) (error, economy.Transaction) {
	return nil, economy.Transaction{}
}

func TestWhenGangsterIsNotInTheMarketShouldThrow(t *testing.T) {
	m := market{}
	err := m.HireGangster(TestPayer{}, &Gangster{})
	if err == nil {
		t.Error("Should return error")
	}
	if _, ok := err.(GangsterNotInTheMarket); !ok {
		t.Error("Should return GangsterNotInTheMarket. Returned: ", err)
	}
}

func TestShouldRemoveGangsterFromMarket(t *testing.T) {
	clock := &turnsClock.Clock{}
	g := &Gangster{
		clock: clock,
	}
	m := market{
		gangsters: []*Gangster{g},
		clock: clock,
	}
	err := m.HireGangster(TestPayer{}, g)
	if (err != nil) {
		t.Error(err)
	}
	if err, _ = m.gangsterIndex(g); err == nil {
		t.Error("Gangster is in the market")
	}
}