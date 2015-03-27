package player

import (
	"code.google.com/p/go-uuid/uuid"
)

type player struct {
	id uuid.UUID
}

func New () player {
	return player{id: uuid.NewUUID()}
}
