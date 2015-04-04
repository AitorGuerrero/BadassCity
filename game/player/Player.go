package player

import (
	userQueries "github.com/AitorGuerrero/BadassCity/user/queries"

	"code.google.com/p/go-uuid/uuid"
	"errors"
)

type Id uuid.UUID
type userId string

type player struct {
	id Id
	userId userId
}

func New (aUserId string) (player, error) {
	isValid, _ := userQueries.IsValid().Execute(aUserId)
	if isValid {
		return player{}, errors.New("Invalid user id")
	}
	return player{
		id: Id(uuid.NewUUID()),
		userId: userId(aUserId),
	}, nil
}
