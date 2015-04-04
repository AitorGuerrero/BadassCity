package services

import (
	"github.com/AitorGuerrero/BadassCity/game"
	"github.com/AitorGuerrero/BadassCity/game/player"
	"github.com/AitorGuerrero/BadassCity/user/queries/isValid"

	"errors"
)

func AddNewPlayer (gameId, userId string, q isValid.Query) (error) {
	isValid, error := q.Execute(userId)
	if error != nil {
		return errors.New("Server error")
	}
	if !isValid {
		return errors.New("User is not valid")
	}
	p, _ := player.New(userId)
	game.Get(gameId).AddPlayer(p)
	return nil
}
