package BadassCity
import (
	t "testing"
	"github.com/AitorGuerrero/BadassCity/economy"
	"github.com/AitorGuerrero/BadassCity/timedEvents/turnsClock"
)

func TestGivenALocalWithNotEnoughRoomShouldThrow(t *t.T) {
	c := &turnsClock.Clock{}
	l := makeFakeLocal(c)
	l.room = 1
	b := givenABusiness()
	err := l.StartABusiness(b);
	if _, ok := err.(notEnoughRoom); !ok{
		t.Error("should throw notEnoughRoom error. Thrown: ", err)
	}
}

func TestGivenAOwnerWithNotEnoughMoneyShouldThrow(t *t.T) {
	c := &turnsClock.Clock{}
	l := makeFakeLocal(c)
	b := givenABusiness()
	economy.Consume(&l.owner.Wallet, economy.Money(5))
	err := l.StartABusiness(b);
	if _, ok := err.(economy.NotEnoughMoney); !ok{
		t.Error("should throw notEnoughMoney error. Thrown: ", err)
	}
}

func TestOwnerShouldSpendTheMoney(t *t.T) {
	c := &turnsClock.Clock{}
	l := makeFakeLocal(c)
	b := givenABusiness()
	err := l.StartABusiness(b);
	if err != nil{
		t.Error(err)
	}
	if l.owner.Wallet.TotalAmount() > 0 {
		t.Error("Should have spent all the money")
	}
}

func givenABusiness() business {
	return business{model: businessModel{
		neededRoom: 2,
		priceForStartPerRoom: 10,
	}}
}