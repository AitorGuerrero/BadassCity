package city

import (
	"github.com/AitorGuerrero/BadassCity/game/city/neighbourhood"
	config "github.com/AitorGuerrero/BadassCity/game/resources/configs/cities"
)

func ConstructCityFromConfig(fileName string) (City) {
	cityConfig, _ :=  config.Factory(fileName)
	return constructCity(cityConfig)
}

func constructCity(cityConfig *config.CityConfig) (City) {
	neighbourhoods := []neighbourhood.Neighbourhood{}
	for _, neighborhoodConfig := range cityConfig.Neighborhoods {
		neighbourhoods = append(
			neighbourhoods,
			neighbourhood.New(neighborhoodConfig.Name),
		)
	}
	return New(cityConfig.Name, neighbourhoods)
}

func constructNeighborhood(neighbourhoodConfig *config.NeighbourhoodConfig) (neighbourhood.Neighbourhood) {
	return neighbourhood.New(neighbourhoodConfig.Name)
}
