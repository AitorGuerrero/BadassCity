package BadassCity

import t "testing"

func TestGivenALocalWithNotBusinessShouldThrow(t *t.T) {
	l := local{}
	err := l.ImproveBusiness()

	if _, ok := err.(localDoesNotHaveABusiness); !ok {
		t.Error("should throw a localDoesNotHaveABusiness. Thrown: ", err)
	}
}

func TestGivenATotallyImprovedBussinessShouldThrow(t *t.T) {
	l := local{
		business: business{
			level: businessLevel(2),
			model: businessModel{
				maxLevel: businessLevel(2),
			},
		},
		hasBusiness: true,
	}
	err := l.ImproveBusiness()
	if _, ok := err.(totallyImprovedBusiness); !ok {
		t.Error("should throw a localDoesNotHaveABusiness. Thrown: ", err)
	}
}

func TestGivenAOwnerWithNotEnoughMoneyForImproveShouldThrow(t *t.T) {
	pricesForImprovePerRoom := make(map[businessLevel]money)
	pricesForImprovePerRoom[businessLevel(1)] = 10
	l := local{
		room: localRoom(2),
		business: business{
			room: localRoom(2),
			level: businessLevel(1),
			model: businessModel{
				pricesForImprovePerRoom: pricesForImprovePerRoom,
				maxLevel: businessLevel(2),
			},
		},
		hasBusiness: true,
		o: &merchant{
			wallet: wallet{
			},
		},
	}
	err := l.ImproveBusiness()
	if _, ok := err.(notEnoughMoney); !ok {
		t.Error("should throw a notEnoughMoney. Thrown: ", err)
	}
}

func TestShouldImproveTheBusiness(t *t.T) {
	pricesForImprovePerRoom := make(map[businessLevel]money)
	pricesForImprovePerRoom[businessLevel(1)] = 10
	l := local{
		room: localRoom(2),
		business: business{
			room: localRoom(2),
			level: businessLevel(1),
			model: businessModel{
				pricesForImprovePerRoom: pricesForImprovePerRoom,
				maxLevel: businessLevel(2),
			},
		},
		hasBusiness: true,
		o: &merchant{
			wallet: wallet{
				transactions: []transaction{transaction{money(100)}},
			},
		},
	}
	err := l.ImproveBusiness()
	if err != nil {
		t.Error(err)
	}
	if l.business.level != 2 {
		t.Error("Should have level 2, has: ", l.business.level)
	}
}