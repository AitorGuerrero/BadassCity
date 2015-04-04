package main

import (
	"github.com/AitorGuerrero/BadassCity/config"

	"github.com/koding/kite"
)

func main() {
	initServices(config.Kite)
}

func initServices(c config.KiteServiceConfig) {
	k := kite.New(c.Name, c.Version)

	InitNewService(k)
	initAddNewPlayerService(k)

	k.Config.Port = c.Port
	k.Run()
}
