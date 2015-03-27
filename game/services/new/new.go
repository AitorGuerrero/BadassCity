package new

import (
	"github.com/AitorGuerrero/BadassCity/api/game"

	"github.com/koding/kite"
)

func Service(r *kite.Request) (interface{}, error) {
	cityType := r.Args.One().MustString()
	game := game.New(city.ConstructCityFromConfig(cityConfigFilesPath))
	result := game.Id()
	return result, nil
}
