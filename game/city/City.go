package city

import (
	"github.com/AitorGuerrero/BadassCity/game/city/neighbourhood"
	"code.google.com/p/go-uuid/uuid"
)

type City interface {}

type id uuid.UUID

type city struct {
	id id
	name string
	neighbourhoods []neighbourhood.Neighbourhood
}

func New (id id, name string, neighbourhoods []neighbourhood.Neighbourhood) city {
	return city{
		id: id,
		name: name,
		neighbourhoods: neighbourhoods,
	}
}

func NewId () id {
	return id(uuid.NewUUID())
}
