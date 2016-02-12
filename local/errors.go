package local

type notEnoughRoom struct {}
func (notEnoughRoom) Error() string {
	return "Not enough room"
}

type totallyImprovedBusiness struct {}
func (totallyImprovedBusiness) Error() string {
	return "Totally improved business"
}

type localDoesNotHaveABusiness struct{}
func (localDoesNotHaveABusiness) Error() string {
	return "Local does not have a business"
}