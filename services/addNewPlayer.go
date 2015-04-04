package services

import (
	"github.com/AitorGuerrero/BadassCity/game"
	"github.com/AitorGuerrero/BadassCity/game/player"
	"github.com/AitorGuerrero/BadassCity/user/queries/isValid"
)

func AddNewPlayer (gameId, userId string, q isValid.Query) (error) {
	isValid, _ := q.Execute(userId)
	if isValid {
		p, _ := player.New(userId)
		game.Get(gameId).AddPlayer(p)
	}
	return nil
}
