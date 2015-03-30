package new

import (
	"github.com/AitorGuerrero/BadassCity/game"
	"github.com/AitorGuerrero/BadassCity/game/city"

	"github.com/koding/kite"
)

var cityConfigFilesPath = "aaa"

func Service(r *kite.Request) (interface{}, error) {
	cityType := r.Args.One().MustString()
	game := game.New(city.ConstructCityFromConfig(cityConfigFilesPath + cityType))
	result := game.Id()
	return result, nil
}
