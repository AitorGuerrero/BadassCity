package main

import (
	"github.com/AitorGuerrero/BadassCity/config"
)

func main() {
	initServices(config.Kite)
}

func initServices(c config.Kite) {
	k := kite.New(c.Name, c.Version)

	InitNewService(k)
	initAddNewPlayerService(k)

	k.Config.Port = c.Port
	k.Run()
}
