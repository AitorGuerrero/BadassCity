package addNewPlayer

import (
	"github.com/AitorGuerrero/BadassCity/game"
	"github.com/AitorGuerrero/BadassCity/game/player"
	userServiceClient "github.com/AitorGuerrero/BadassCity/user/services/client"
)

func Service (gameId, userId string) (interface{}, error) {
	response, _ := userServiceClient.Get().Tell("is-valid", userId)
	isValid := response.MustBool()
	if (isValid) {
		game.Get(gameId).AddPlayer(player.New(userId))
	}
	return nil, nil
}
