package economy

type wallet struct {
	incomingTransactions []*transaction
	outcomingTransactions []*transaction
}

func (this *wallet) receiveTransaction (t *transaction) {
	this.incomingTransactions = append(this.incomingTransactions, t)
	return
}

func (w *wallet) generateTransaction(amount Money) (err error, t *transaction) {
	if !w.hasEnoughMoney(amount) {
		err = NotEnoughMoney{}
		return
	}
	t = &transaction{amount}
	w.outcomingTransactions = append(w.outcomingTransactions, t)
	return
}

func (w *wallet) totalAmount() (amount Money) {
	var positiveAmount, negativeAmount int
	for _, t := range w.incomingTransactions {
		positiveAmount += int(t.amount)
	}
	for _, t := range w.outcomingTransactions {
		negativeAmount += int(t.amount)
	}
	return Money(positiveAmount - negativeAmount)
}

func (w *wallet) hasEnoughMoney(m Money) bool {
	return w.totalAmount() >= m
}
