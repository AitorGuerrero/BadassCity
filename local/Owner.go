package local

import "github.com/AitorGuerrero/goConomy"

const benefitsMultiplierForSelling = 5
const priceForRoom = goConomy.Money(100)

type Owner interface {
	goConomy.MoneyGiver
	goConomy.MoneyReceiver
	HasEnoughMoney(goConomy.Money) bool
}

func trespassLocalTo(l local, from Owner, to Owner) (err error) {
	err = to.GiveMoneyTo(calculatePriceForLocal(l), from)
	if nil != err {
		return
	}
	l.owner = to

	return
}

func calculatePriceForLocal(l local) goConomy.Money {
	price := priceForRoom.Multiply(float32(l.room))
	if l.hasBusiness {
		price = price.Add(l.business.benefits().Multiply(benefitsMultiplierForSelling))
	}

	return price
}
