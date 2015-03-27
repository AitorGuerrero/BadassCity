package city

import (
	"io/ioutil"
	"encoding/json"
	"log"

	"github.com/AitorGuerrero/BadassCity/game/city/neighbourhood"
)

type neighbourhoodConfig struct {
	name string
}
type cityConfig struct {
	name string
	neighborhoods []neighbourhoodConfig
}

func ConstructCityFromConfig(fileName string) (City) {
	file, e := ioutil.ReadFile(fileName)
	if e != nil {
		log.Fatal(e)
	}
	var cityConfig cityConfig
	e = json.Unmarshal(file, &cityConfig)
	if e != nil {
		log.Fatal(e)
	}

	return constructCity(cityConfig)
}

func constructCity(cityConfig cityConfig) (City) {
	neighbourhoods := []Neighbourhood{}
	for _, neighborhoodConfig := range cityConfig.neighborhoods {
		neighbourhoods = append(
			neighbourhoods,
			neighbourhood.New(neighborhoodConfig.name),
		)
	}
	aCity := New(cityConfig.name, neighbourhoods)
	return aCity
}

func constructNeighborhood(neighbourhoodConfig neighbourhoodConfig) (Neighbourhood) {
	return neighbourhood.New(neighbourhoodConfig.name)
}
