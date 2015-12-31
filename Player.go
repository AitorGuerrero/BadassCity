package BadassCity
import "github.com/AitorGuerrero/BadassCity/economy"

type player struct {
	economy.Merchant
}

func NewPlayer() player {
	return player{}
}

func (p *player) BuyLocal(l *local) (err error) {
	err = p.Wallet.TransferTo(&l.owner.Wallet, l.price)
	if (err != nil) {
		return
	}
	l.changeOwner(&(p.Merchant))

	return
}