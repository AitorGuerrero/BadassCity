package main

import (
	newGameService "github.com/AitorGuerrero/BadassCity/services/new/kite"
	addPlayerToGameService "github.com/AitorGuerrero/BadassCity/services/player/addNewPlayer/kite"
	"github.com/AitorGuerrero/BadassCity/config"

	"github.com/koding/kite"
)

func main() {
	cfg := config.Get()
	initServices(cfg.KiteServiceConfig)
}

func initServices(c config.KiteServiceConfig) {
	k := kite.New(c.Name, c.Version)

	newGameService.AddService(k)
	addPlayerToGameService.AddService(k)

	k.Config.Port = c.Port
	k.Run()
}
