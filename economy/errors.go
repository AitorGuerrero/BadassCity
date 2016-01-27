package economy

type NotEnoughMoney struct {}
func (NotEnoughMoney) Error() string {
	return "Not enough money"
}
