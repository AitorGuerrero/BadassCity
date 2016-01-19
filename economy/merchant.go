package economy

type Merchant struct {
	Wallet Wallet
}

type MoneyReceiver interface {
	ReceiveMoney(Money, *Wallet)
}

type MoneyGiver interface {
	GiveMoney(Money, *Wallet)
}