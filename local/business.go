package local

import "github.com/AitorGuerrero/goConomy"

type businessLevel int

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

func (b business) priceForImprovePerRoom() goConomy.Money {
	return b.model.pricesForImprovePerRoom[businessLevel(int(b.level) + 1)]
}

func (b business) benefitsPerRoom() goConomy.Money {
	return b.model.revenueByLevel[b.level]
}

func (b business) benefits() goConomy.Money {
	return b.benefitsPerRoom().Multiply(float32(b.room))
}

func (b business) priceForImprove() goConomy.Money {
	return b.priceForImprovePerRoom().Multiply(float32(b.room))
}

func (b *business) improve() {
	b.level = b.level.increase()
}