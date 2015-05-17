package addNewPlayer

import (
	"github.com/AitorGuerrero/BadassCity/game"
	"github.com/AitorGuerrero/BadassCity/game/player"

	"errors"
)

type Service interface {
	Execute(gameId GameId, userId UserId) (error)
}

type UserIsValidQuery interface {
	Execute(userId interface{}) (bool, error)
}

type addNewPlayer struct {
	q UserIsValidQuery
}

func New (q UserIsValidQuery) Service {
	s := addNewPlayer{q: q}
	return s
}

func (np addNewPlayer) Execute (gameId string, userId string) error {
	error := assertUserIsValid(userId.toString(), np.q)
	if nil != error {
		p, _ := player.New(userId.toString())
		game.Get(gameId.toString()).AddPlayer(p)
	}
	return error
}

func assertUserIsValid (userId string, q UserIsValidQuery) error {
	isValid, error := q.Execute(userId)
	if error != nil {
		return errors.New("Server error")
	}
	if !isValid {
		return errors.New("User is not valid")
	}
	return nil
}
