package BadassCity

type wallet struct {
	amount float32
	transactions []transaction
}

type merchant interface {
	getTransaction(transaction)
	giveTransaction(float32) (error, transaction)
}

func (w *wallet) totalAmount() (amount float32) {
	for _, t := range w.transactions {
		amount += t.amount()
	}
	return
}
