package addNewPlayer

import (
	"github.com/AitorGuerrero/BadassCity/game"
	"github.com/AitorGuerrero/BadassCity/game/player"
	userServiceClient "github.com/AitorGuerrero/BadassCity/user/service/client"

	"github.com/koding/kite"
)

func Service (r *kite.Request) (interface{}, error) {
	args := r.Args.Map()
	gameId := args["gameId"].MustString()
	userId := args["userId"].MustString()
	isValid, _ := userServiceClient.Get().Tell("is-valid", userId)
	if (isValid) {
		game.Get(gameId).AddPlayer(player.New(userId))
	}
	return nil, nil
}
