package BadassCity
import "github.com/AitorGuerrero/BadassCity/economy"

type transaction struct {
	a economy.Money
}

func (t transaction) amount() economy.Money {
	return t.a
}

func trespassTransaction (amount economy.Money, giver, receiver *merchant) (err error) {
	err, t := giver.giveTransaction(amount)
	if err != nil {
		return
	}
	receiver.getTransaction(t)
	return
}