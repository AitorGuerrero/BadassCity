package services

import (
	"github.com/AitorGuerrero/BadassCity/game"
	"github.com/AitorGuerrero/BadassCity/game/player"
	"github.com/AitorGuerrero/BadassCity/user/queries/isValid"

	"errors"
)

func AddNewPlayer (gameId, userId string, q isValid.Query) (error) {
	isValid, _ := q.Execute(userId)
	if !isValid {
		return errors.New("User is not valid")
	}
	p, _ := player.New(userId)
	game.Get(gameId).AddPlayer(p)
	return nil
}
