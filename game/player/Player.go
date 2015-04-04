package player

import (
	"code.google.com/p/go-uuid/uuid"
)

type Id uuid.UUID
type userId string

type player struct {
	id Id
	userId userId
}

func New (aUserId string) (player, error) {
	return player{
		id: Id(uuid.NewUUID()),
		userId: userId(aUserId),
	}, nil
}
