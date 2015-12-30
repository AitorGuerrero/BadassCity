package economy

type Transaction struct {
	amount Money
}

func NewTransaction(a Money) Transaction {
	return Transaction{a}
}

func (t Transaction) Amount() Money {
	return t.amount
}