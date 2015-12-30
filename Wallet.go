package BadassCity
import "github.com/AitorGuerrero/BadassCity/economy"

type wallet struct {
	transactions []economy.Transaction
}

type notEnoughMoney struct {}
func (notEnoughMoney) Error() string {
	return "Not enough money"
}

func (w *wallet) totalAmount() (amount economy.Money) {
	for _, t := range w.transactions {
		amount += t.Amount()
	}
	return
}
