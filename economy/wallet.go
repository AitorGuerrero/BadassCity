package economy

type Wallet struct {
	transactions []Transaction
}

type NotEnoughMoney struct {}
func (NotEnoughMoney) Error() string {
	return "Not enough money"
}

func (w *Wallet) AddTransaction(t Transaction) (err error) {
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

func (w *Wallet) GetTransaction(a Money) (err error, t Transaction) {
	if !w.HasEnoughMoney(a) {
		err = NotEnoughMoney{}
		return
	}
	w.AddTransaction(Transaction{-a})
	t = Transaction{a}
	return
}

func (from *Wallet) TransferTo (to *Wallet, amount Money) (err error) {
	err, t := from.GetTransaction(amount)
	if err != nil {
		return
	}
	to.AddTransaction(t)
	return
}