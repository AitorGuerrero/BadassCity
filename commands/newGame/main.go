package main

import (
	"github.com/AitorGuerrero/BadassCity/game"
	"github.com/AitorGuerrero/BadassCity/game/city"

	"fmt"
)

var cityConfigFilesPath = "../../resources/configs/cities/test.json"

func main() {
	game := game.New(city.ConstructCityFromConfig(cityConfigFilesPath))

	fmt.Println(game.Id())
}
