package main

import (
	"github.com/AitorGuerrero/BadassCity/services"
	"github.com/AitorGuerrero/BadassCity/config"

	"github.com/koding/kite"
)

func initServices(c config.KiteServiceConfig) {
	k := kite.New(c.Name, c.Version)

	InitNewService(k)
	initAddNewPlayerService(k)

	k.Config.Port = c.Port
	k.Run()
}

func initAddNewPlayerService(k *kite.Kite) {
	k.HandleFunc("add-new-player", func (r *kite.Request) (interface{}, error) {
			args := r.Args.MustSlice()
			gameId := args[0].MustString()
			userId := args[1].MustString()
			return nil, services.AddNewPlayer(gameId, userId)
		}).DisableAuthentication()
}

func InitNewService(k *kite.Kite) {
	k.HandleFunc("new", func (r *kite.Request) (interface{}, error) {
			c := r.Args.One().MustString()
			return services.NewGame(c)
		}).DisableAuthentication()
}
