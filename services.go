package main

import (
	"github.com/AitorGuerrero/BadassCity/services"
	"github.com/AitorGuerrero/BadassCity/user/queries/isValid"

	"github.com/koding/kite"
)

func initAddNewPlayerService(k *kite.Kite) {
	k.HandleFunc("add-new-player", func (r *kite.Request) (interface{}, error) {
			args := r.Args.MustSlice()
			gId := args[0].MustString()
			uId := args[1].MustString()
			q := isValid.New()
			return nil, services.AddNewPlayer(gId, uId, q)
		}).DisableAuthentication()
}

func InitNewService(k *kite.Kite) {
	k.HandleFunc("new", func (r *kite.Request) (interface{}, error) {
			c := r.Args.One().MustString()
			return services.NewGame(c)
		}).DisableAuthentication()
}
