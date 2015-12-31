package economy

func GenerateMoney(w *Wallet, m Money) {
	w.incomingTransactions = append(w.incomingTransactions, &transaction{m})
}