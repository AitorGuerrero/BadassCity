package economy

import t "testing"

func TestGivenAMoneyDestroyerWhenGiverHaveNotEnoguhMoneyShouldReturnError(t *t.T) {
	destroyer := MoneyDestroyer{}
	giver := Merchant{}
	giver.receiveTransaction(&transaction{Money(5)})
	err := destroyer.ConsumeMoney(&giver, Money(10))
	if nil == err {
		t.Error("Should return error")
	}
	if _, ok := err.(NotEnoughMoney); !ok {
		t.Error("Error should be NotEnoughMoney")
	}
	if "Not enough money" != err.Error() {
		t.Error("Incorrect error message")
	}
}

func TestGivenAMoneyDestroyerAskingToGetMoneyShouldGetMoney(t *t.T) {
	destroyer := MoneyDestroyer{}
	giver := Merchant{}
	giver.receiveTransaction(&transaction{Money(13)})
	destroyer.ConsumeMoney(&giver, Money(5))
	if giver.wallet.totalAmount() != Money(8) {
		t.Error("Should decrease giver total")
	}
}
