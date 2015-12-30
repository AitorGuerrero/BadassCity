package BadassCity
import "github.com/AitorGuerrero/BadassCity/economy"

type merchant struct {
	wallet economy.Wallet
}

func (pl merchant) hasEnoughMoney(a economy.Money) bool {
	return pl.availableMoney() >= a
}

func (m *merchant) getTransaction(t economy.Transaction) {
	m.wallet.AddTransaction(t)
}

func (m *merchant) availableMoney() economy.Money {
	return economy.Money(m.wallet.TotalAmount())
}

func (m *merchant) giveTransaction(a economy.Money) (err error, t economy.Transaction) {
	if !m.hasEnoughMoney(a) {
		err = economy.NotEnoughMoney{}
		return
	}
	m.getTransaction(economy.NewTransaction(-a))
	t = economy.NewTransaction(a)
	return
}

func trespassTransaction (amount economy.Money, giver, receiver *merchant) (err error) {
	err, t := giver.giveTransaction(amount)
	if err != nil {
		return
	}
	receiver.getTransaction(t)
	return
}