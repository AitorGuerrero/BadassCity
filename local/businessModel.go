package local

import "github.com/AitorGuerrero/goConomy"

type businessModel struct {
	priceForStartPerRoom goConomy.Money
	neededRoom localRoom
	maxLevel businessLevel
	pricesForImprovePerRoom map[businessLevel]goConomy.Money
	revenueByLevel map[businessLevel]goConomy.Money
}
