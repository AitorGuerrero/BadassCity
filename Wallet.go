package BadassCity
import "github.com/AitorGuerrero/BadassCity/economy"

type wallet struct {
	transactions []transaction
}

type notEnoughMoney struct {}
func (notEnoughMoney) Error() string {
	return "Not enough money"
}

type merchant struct {
	wallet wallet
}

func (pl merchant) hasEnoughMoney(a economy.Money) bool {
	return pl.availableMoney() >= a
}

func (m *merchant) getTransaction(t transaction) {
	m.wallet.transactions = append(m.wallet.transactions, t)
}

func (m *merchant) availableMoney() economy.Money {
	return economy.Money(m.wallet.totalAmount())
}

func (m *merchant) giveTransaction(a economy.Money) (err error, t transaction) {
	if !m.hasEnoughMoney(a) {
		err = notEnoughMoney{}
		return
	}
	m.getTransaction(transaction{-a})
	t = transaction{a}
	return
}

func (w *wallet) totalAmount() (amount economy.Money) {
	for _, t := range w.transactions {
		amount += t.amount()
	}
	return
}
