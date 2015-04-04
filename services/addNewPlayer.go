package services

import (
	"github.com/AitorGuerrero/BadassCity/game"
	"github.com/AitorGuerrero/BadassCity/game/player"
	userQueries "github.com/AitorGuerrero/BadassCity/user/queries"
)

func AddNewPlayer (gameId, userId string) (error) {
	isValid, _ := userQueries.IsValid().Execute(userId)
	if isValid {
		p, _ := player.New(userId)
		game.Get(gameId).AddPlayer(p)
	}
	return nil
}
