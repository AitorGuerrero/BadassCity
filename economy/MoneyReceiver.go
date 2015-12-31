package economy

func TakeMoney(w *Wallet, m Money) (err error) {
	err, _ = w.generateTransaction(m)
	return
}