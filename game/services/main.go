package main

import (
	"github.com/AitorGuerrero/BadassCity/game/services/new"
	"github.com/AitorGuerrero/BadassCity/game/services/player/addNewPlayer"

	"github.com/koding/kite"
)

var port = 3636

func main() {
	k := kite.New("BadassCity.game", "0.0.1")
	k.HandleFunc("new", new.Service).DisableAuthentication()
	k.HandleFunc("add-new-player", addNewPlayer.Service).DisableAuthentication()

	k.Config.Port = Port()
	k.Run()
}

func Port() int {
	return port
}
