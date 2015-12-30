package BadassCity
import "github.com/AitorGuerrero/BadassCity/economy"

type businessModel struct {
	priceForStartPerRoom economy.Money
	neededRoom localRoom
	maxLevel businessLevel
	pricesForImprovePerRoom map[businessLevel]economy.Money
	revenueByLevel map[businessLevel]economy.Money
}
