package game

import (
	"github.com/AitorGuerrero/BadassCity/game/city"

	"code.google.com/p/go-uuid/uuid"
)

type Game interface {}
type id uuid.UUID
type game struct {
	id id
	city city.City
	running bool
}

func New (i id, c city.City) *game {
	return &game{
		i,
		c,
		false,
	}
}

func NewId() id {
	i := id(uuid.NewUUID())
	return i
}

func (g *game) Id() *id {
	return &g.id
}
