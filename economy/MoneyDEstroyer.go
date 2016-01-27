package economy

type MoneyDestroyer struct {
	moneyReceiver
}

func (this MoneyDestroyer) ConsumeMoney(from moneyGiver, m Money) (err error) {
	err, t := from.giveTransaction(m)
	if nil != err {
		return
	}
	this.receiveTransaction(t)

	return
}

func (MoneyDestroyer) receiveTransaction(*transaction) {

}