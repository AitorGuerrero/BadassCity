package player

import (
	"code.google.com/p/go-uuid/uuid"
)

type Player interface {}
type id uuid.UUID
type userId string

type player struct {
	id id
	userId userId
}

func New (i id, ui string) (player) {
	return player{i, userId(ui)}
}

func NewId () {
	return id(uuid.NewUUID())
}
