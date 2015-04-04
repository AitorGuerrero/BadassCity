package cities

var TestCityConfig = CityConfig {
	"TestCity",
	[]NeighbourhoodConfig {
		NeighbourhoodConfig{
			"n1",
			[]Local {
				Local{1000, SizeS},
				Local{1500, SizeL},
			},
		},
		NeighbourhoodConfig{
			"n2",
			[]Local {
				Local{1500, SizeM},
				Local{2000, SizeXL},
			},
		},
	},
}
