package BadassCity
import "github.com/AitorGuerrero/BadassCity/economy"

type merchant struct {
	wallet wallet
}

func (pl merchant) hasEnoughMoney(a economy.Money) bool {
	return pl.availableMoney() >= a
}

func (m *merchant) getTransaction(t economy.Transaction) {
	m.wallet.transactions = append(m.wallet.transactions, t)
}

func (m *merchant) availableMoney() economy.Money {
	return economy.Money(m.wallet.totalAmount())
}

func (m *merchant) giveTransaction(a economy.Money) (err error, t economy.Transaction) {
	if !m.hasEnoughMoney(a) {
		err = notEnoughMoney{}
		return
	}
	m.getTransaction(economy.NewTransaction(-a))
	t = economy.NewTransaction(a)
	return
}