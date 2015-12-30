package BadassCity

type businessModel struct {
	priceForStartPerRoom money
	neededRoom localRoom
	maxLevel businessLevel
	pricesForImprovePerRoom map[businessLevel]money
	revenueByLevel map[businessLevel]money
}
