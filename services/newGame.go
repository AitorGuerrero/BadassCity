package services

import (
	"github.com/AitorGuerrero/BadassCity/game"
	"github.com/AitorGuerrero/BadassCity/game/city"
	"github.com/AitorGuerrero/BadassCity/config"
)

func NewGame(cityType string) (string, error) {
	game := game.New(city.ConstructCityFromConfig(config.Main.CitiesConfigPath + cityType))
	result := game.Id()
	return string(result), nil
}
