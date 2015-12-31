package economy

type Wallet struct {
	transactions []transaction
}

type NotEnoughMoney struct {}
func (NotEnoughMoney) Error() string {
	return "Not enough money"
}

func (from *Wallet) TransferTo (to *Wallet, amount Money) (err error) {
	err, t := from.getTransaction(amount)
	if err != nil {
		return
	}
	to.addTransaction(t)
	return
}

func (w *Wallet) addTransaction(t transaction) (err error) {
	if w.TotalAmount() + t.Amount < 0 {
		err = NotEnoughMoney{}
		return
	}
	w.transactions = append(w.transactions, t)
	return
}

func (w *Wallet) TotalAmount() (amount Money) {
	for _, t := range w.transactions {
		amount += t.Amount
	}
	return
}

func (w *Wallet) HasEnoughMoney(m Money) bool {
	return w.TotalAmount() >= m
}

func (w *Wallet) getTransaction(a Money) (err error, t transaction) {
	if !w.HasEnoughMoney(a) {
		err = NotEnoughMoney{}
		return
	}
	w.addTransaction(transaction{-a})
	t = transaction{a}
	return
}