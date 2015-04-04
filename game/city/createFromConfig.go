package city

import (
	"github.com/AitorGuerrero/BadassCity/game/city"
	"github.com/AitorGuerrero/BadassCity/game/city/neighbourhood"
	config "github.com/AitorGuerrero/BadassCity/game/resources/configs/cities"
)

func ConstructCityFromConfig(fileName string) (City) {
	return constructCity(config.Factory(fileName))
}

func constructCity(cityConfig config.CityConfig) (City) {
	neighbourhoods := []neighbourhood.Neighbourhood{}
	for _, neighborhoodConfig := range cityConfig.neighborhoods {
		neighbourhoods = append(
			neighbourhoods,
			neighbourhood.New(neighborhoodConfig.name),
		)
	}
	return city.New(cityConfig.name, neighbourhoods)
}

func constructNeighborhood(neighbourhoodConfig config.NeighbourhoodConfig) (neighbourhood.Neighbourhood) {
	return neighbourhood.New(neighbourhoodConfig.name)
}
