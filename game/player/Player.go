package player

import (
	"code.google.com/p/go-uuid/uuid"
)

type userId uuid.UUID

type player struct {
	id uuid.UUID
	userId userId
}

func New (userId userId) player {
	return player{
		id: uuid.NewUUID(),
		userId: userId,
	}
}
