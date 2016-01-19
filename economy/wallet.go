package economy

type Wallet struct {
	incomingTransactions []*transaction
	outcomingTransactions []*transaction
}

type NotEnoughMoney struct {}
func (NotEnoughMoney) Error() string {
	return "Not enough money"
}

func (from *Wallet) TransferTo (to *Wallet, amount Money) (err error) {
	err, t := from.generateTransaction(amount)
	if err != nil {
		return
	}
	to.incomingTransactions = append(to.incomingTransactions, t)
	return
}

func (w *Wallet) TotalAmount() (amount Money) {
	var positiveAmount, negativeAmount int
	for _, t := range w.incomingTransactions {
		positiveAmount += int(t.amount)
	}
	for _, t := range w.outcomingTransactions {
		negativeAmount += int(t.amount)
	}
	return Money(positiveAmount - negativeAmount)
}

func (w *Wallet) HasEnoughMoney(m Money) bool {
	return w.TotalAmount() >= m
}

func (w *Wallet) GiveMoney(Money, mr MoneyReceiver) {
}

func (w *Wallet) generateTransaction(a Money) (err error, t *transaction) {
	if !w.HasEnoughMoney(a) {
		err = NotEnoughMoney{}
		return
	}
	t = &transaction{a}
	w.outcomingTransactions = append(w.outcomingTransactions, t)
	return
}