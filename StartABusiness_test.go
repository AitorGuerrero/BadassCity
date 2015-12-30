package BadassCity
import (
	t "testing"
	"github.com/AitorGuerrero/BadassCity/economy"
)

func TestGivenALocalWithNotEnoughRoomShouldThrow(t *t.T) {
	l := makeFakeLocal()
	l.room = 1
	b := givenABusiness()
	err := l.StartABusiness(b);
	if _, ok := err.(notEnoughRoom); !ok{
		t.Error("should throw notEnoughRoom error. Thrown: ", err)
	}
}

func TestGivenAOwnerWithNotEnoughMoneyShouldThrow(t *t.T) {
	l := makeFakeLocal()
	b := givenABusiness()
	l.owner.wallet.AddTransaction(economy.NewTransaction(-5))
	err := l.StartABusiness(b);
	if _, ok := err.(economy.NotEnoughMoney); !ok{
		t.Error("should throw notEnoughMoney error. Thrown: ", err)
	}
}

func TestOwnerShouldSpendTheMoney(t *t.T) {
	l := makeFakeLocal()
	b := givenABusiness()
	err := l.StartABusiness(b);
	if err != nil{
		t.Error(err)
	}
	if l.owner.availableMoney() > 0 {
		t.Error("Should have spent all the money")
	}
}

func givenABusiness() business {
	return business{model: businessModel{
		neededRoom: 2,
		priceForStartPerRoom: 10,
	}}
}