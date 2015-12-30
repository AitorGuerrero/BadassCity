package BadassCity

type businessLevel int

type notEnoughRoom struct {}
func (notEnoughRoom) Error() string {
	return "Not enough room"
}
type totallyImprovedBusiness struct {}
func (totallyImprovedBusiness) Error() string {
	return "Totally improved business"
}

func (l businessLevel) increase() businessLevel {
	return businessLevel(int(l)+1)
}

func (l businessLevel) isGreaterOrEqualThan(b businessLevel) bool {
	return int(l) >= int(b)
}

type business struct {
	room localRoom
	model businessModel
	level businessLevel
}

func (b business) isTotallyImproved() bool {
	return b.level.isGreaterOrEqualThan(b.model.maxLevel)
}

func (b business) priceForImprovePerRoom() money {
	return b.model.pricesForImprovePerRoom[b.level]
}

func (b business) benefitsPerRoom() money {
	return b.model.revenueByLevel[b.level]
}

func (b business) benefits() money {
	return b.benefitsPerRoom().multiply(float32(b.room))
}

func (b business) priceForImprove() money {
	return b.priceForImprovePerRoom().multiply(float32(b.room))
}

func (b *business) improve() error {
	if b.isTotallyImproved() {
		return totallyImprovedBusiness{}
	}
	b.level = b.level.increase()

	return nil
}