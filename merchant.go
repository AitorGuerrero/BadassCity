package BadassCity
import "github.com/AitorGuerrero/BadassCity/economy"

type merchant struct {
	wallet economy.Wallet
}

func trespassTransaction (amount economy.Money, giver, receiver *merchant) (err error) {
	err, t := giver.wallet.GetTransaction(amount)
	if err != nil {
		return
	}
	receiver.wallet.AddTransaction(t)
	return
}