package economy

type Wallet struct {
	transactions []Transaction
}

func NewWallet() Wallet {
	return Wallet{}
}

type NotEnoughMoney struct {}
func (NotEnoughMoney) Error() string {
	return "Not enough money"
}

func (w *Wallet) AddTransaction(t Transaction) {
	w.transactions = append(w.transactions, t)
}

func (w *Wallet) TotalAmount() (amount Money) {
	for _, t := range w.transactions {
		amount += t.Amount()
	}
	return
}
