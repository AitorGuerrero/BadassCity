package BadassCity

import "github.com/AitorGuerrero/BadassCity/economy"

func trespassTransaction (amount economy.Money, giver, receiver *merchant) (err error) {
	err, t := giver.giveTransaction(amount)
	if err != nil {
		return
	}
	receiver.getTransaction(t)
	return
}