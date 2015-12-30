package BadassCity
import t "testing"

func TestGivenALocalWithNotEnoughRoomShouldThrow(t *t.T) {
	l := givenALocal()
	l.room = 1
	b := givenABusiness()
	err := l.StartABusiness(b);
	if _, ok := err.(notEnoughRoom); !ok{
		t.Error("should throw notEnoughRoom error. Thrown: ", err)
	}
}

func TestGivenAOwnerWithNotEnoughMoneyShouldThrow(t *t.T) {
	l := givenALocal()
	b := givenABusiness()
	l.o.wallet.transactions = []transaction{transaction{5}}
	err := l.StartABusiness(b);
	if _, ok := err.(notEnoughMoney); !ok{
		t.Error("should throw notEnoughMoney error. Thrown: ", err)
	}
}

func TestOwnerShouldSpendTheMoney(t *t.T) {
	l := givenALocal()
	b := givenABusiness()
	err := l.StartABusiness(b);
	if err != nil{
		t.Error(err)
	}
	if l.o.availableMoney() > 0 {
		t.Error("Should have spent all the money")
	}
}

func givenABusiness() business {
	return business{model: businessModel{
		neededRoom: 2,
		priceForStartPerRoom: 10,
	}}
}

func givenALocal() local {
	return local{room: 2, o: givenAMerchant()}
}

func givenAMerchant() *merchant {
	return &merchant{
		wallet: wallet{
			transactions: []transaction{
				transaction{a: money(20)},
			},
		},
	}
}