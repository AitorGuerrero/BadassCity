package economy

type moneyReceiver interface {
	receiveTransaction(*transaction)
}

type moneyGiver interface {
	giveTransaction(Money) (error, *transaction)
}

type Merchant struct {
	moneyReceiver
	moneyGiver
	wallet wallet
}

func (this *Merchant) GiveMoneyTo(amount Money, receiver moneyReceiver) (err error) {
	err, t := this.giveTransaction(amount)
	if (nil != err) {
		return;
	}
	receiver.receiveTransaction(t)

	return
}

func (this *Merchant) giveTransaction(amount Money) (err error, t *transaction) {
	err, t = this.wallet.generateTransaction(amount)
	return
}

func (this *Merchant) receiveTransaction(t *transaction) {
	this.wallet.receiveTransaction(t)
}