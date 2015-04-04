package services

import (
	"github.com/AitorGuerrero/BadassCity/game"
	"github.com/AitorGuerrero/BadassCity/game/city"
)

var cityConfigFilesPath = "aaa"

func NewGame(cityType string) (string, error) {
	game := game.New(city.ConstructCityFromConfig(cityConfigFilesPath + cityType))
	result := game.Id()
	return string(result), nil
}
