package economy

type MoneyReceiver struct {
	transactions []transaction
}

func (mr *MoneyReceiver) TakeMoney(w *Wallet, m Money) error {
	err, t := w.getTransaction(m)
	if err != nil {
		return err
	}
	mr.transactions = append(mr.transactions, t)

	return nil
}