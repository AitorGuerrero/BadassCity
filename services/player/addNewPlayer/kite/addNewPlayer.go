package kite

import (
	"github.com/AitorGuerrero/BadassCity/services/player/addNewPlayer"

	"github.com/koding/kite"
)

func AddService(k *kite.Kite) {
	k.HandleFunc("add-new-player", func (r *kite.Request) (interface{}, error) {
			args := r.Args.MustSlice()
			gameId := args[0].MustString()
			userId := args[1].MustString()
			return addNewPlayer.Service(gameId, userId)
		}).DisableAuthentication()
}
