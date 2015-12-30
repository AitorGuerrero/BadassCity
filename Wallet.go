package BadassCity

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

func (pl merchant) hasEnoughMoney(a money) bool {
	return pl.availableMoney() >= a
}

func (m *merchant) getTransaction(t transaction) {
	m.wallet.transactions = append(m.wallet.transactions, t)
}

func (m *merchant) availableMoney() money {
	return money(m.wallet.totalAmount())
}

func (m *merchant) giveTransaction(a money) (err error, t transaction) {
	if !m.hasEnoughMoney(a) {
		err = notEnoughMoney{}
		return
	}
	m.getTransaction(transaction{-a})
	t = transaction{a}
	return
}

func (w *wallet) totalAmount() (amount money) {
	for _, t := range w.transactions {
		amount += t.amount()
	}
	return
}
