package main

import (
	"github.com/AitorGuerrero/BadassCity/game/services/new"
	"github.com/AitorGuerrero/BadassCity/game/services/player/addNewPlayer"
	"github.com/AitorGuerrero/BadassCity/persistence"
	"github.com/AitorGuerrero/BadassCity/persistence/gorm"

	"github.com/koding/kite"
)

var port = 3636
var dbUserName = ""
var dbPassword = ""
var dbName = ""

func main() {
	persistence.SetPersistLayer(gorm.Get(dbUserName, dbPassword, dbName))
	k := kite.New("BadassCity.game", "0.0.1")

	k.HandleFunc("new-game", func (r *kite.Request) (interface{}, error) {
		return new.Service(r.Args.One().MustString())
	}).DisableAuthentication()

	k.HandleFunc("add-new-player", func (r *kite.Request) (interface{}, error) {
		args := r.Args.MustSliceOfLength(2)
		gameId := args[0].MustString()
		userId := args[1].MustString()
		return addNewPlayer.Service(gameId, userId)
	}).DisableAuthentication()

	k.Config.Port = Port()
	k.Run()
}

func Port() int {
	return port
}
