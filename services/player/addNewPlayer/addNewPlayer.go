package addNewPlayer

import (
	"github.com/AitorGuerrero/BadassCity/game"
	"github.com/AitorGuerrero/BadassCity/game/player"
	userQueries "github.com/AitorGuerrero/BadassCity/user/queries"
)

func Service (gameId, userId string) (interface{}, error) {
	isValid, _ := userQueries.IsValid().Execute(userId)
	if isValid {
		game.Get(gameId).AddPlayer(player.New(userId))
	}
	return nil, nil
}
