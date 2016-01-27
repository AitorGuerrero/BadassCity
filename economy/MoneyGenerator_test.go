package economy

import t "testing"

func TestGivenAMoneyGeneratorAskingToGiveMoneyShouldGetMoney(t *t.T) {
	generator := MoneyGenerator{}
	receiver := Merchant{}
	generator.GenerateMoney(&receiver, Money(5))
	if receiver.wallet.totalAmount() != Money(5) {
		t.Error("Should increase giver total")
	}
}
