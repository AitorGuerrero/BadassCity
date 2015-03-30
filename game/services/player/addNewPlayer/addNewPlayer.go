package addNewPlayer

import (
	"github.com/AitorGuerrero/BadassCity/game"
	"github.com/AitorGuerrero/BadassCity/game/player"
	userServiceClient "github.com/AitorGuerrero/BadassCity/user/services/client"

	"github.com/koding/kite"
)

func Service (r *kite.Request) (interface{}, error) {
	args  := r.Args.MustSliceOfLength(2)
	gameId := args[0].MustString()
	userId := args[1].MustString()
	response, _ := userServiceClient.Get().Tell("is-valid", userId)
	isValid := response.MustBool()
	if (isValid) {
		game.Get(gameId).AddPlayer(player.New(userId))
	}
	return nil, nil
}
