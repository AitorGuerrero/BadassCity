package new

import (
	"github.com/AitorGuerrero/BadassCity/game"
	"github.com/AitorGuerrero/BadassCity/game/city"
	"github.com/AitorGuerrero/BadassCity/persistence"

	"code.google.com/p/go-uuid/uuid"
)

var cityConfigFilesPath = "aaa"

func Service(cityType string) (uuid.UUID, error) {
	game := game.New(city.ConstructCityFromConfig(cityConfigFilesPath + cityType))
	result := game.Id()
	persistence.Flush()
	return result, nil
}
