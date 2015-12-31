package economy

type MoneyGenerator struct {

}

func (mg MoneyGenerator) GenerateMoney(w *Wallet, m Money) {
	w.addTransaction(transaction{m})
}