package economy

import t "testing"

func TestGivenAMerchantWheAskingForMoneyButHasNotEnoughMoneyShouldReturnError(t *t.T) {
	receiver := MoneyDestroyer{}
	merchant := Merchant{}
	merchant.receiveTransaction(&transaction{Money(10)})
	err := merchant.GiveMoneyTo(Money(15), receiver)
	if nil == err {
		t.Error("Should return error")
	}
	if _, ok := err.(NotEnoughMoney); !ok {
		t.Error("Should be of type NotEnoughMoney")
	}
}

func TestGivenAMerchantAskingForMoneyShouldReduceHisTotal(t *t.T) {
	initialTotal := Money(13)
	transactionAmount := Money(5)
	expectedFinalTotal := Money(8)
	receiver := MoneyDestroyer{}
	merchant := Merchant{}
	merchant.receiveTransaction(&transaction{initialTotal})
	err := merchant.GiveMoneyTo(transactionAmount, receiver)
	if nil != err {
		t.Error(err)
	}
	if merchant.wallet.totalAmount() != expectedFinalTotal {
		t.Errorf("The total should be %v. Is %v", expectedFinalTotal, merchant.wallet.totalAmount())
	}
}
