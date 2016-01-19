package economy

func Consume(w *Wallet, m Money) (err error) {
	err, _ = w.generateTransaction(m)
	return
}