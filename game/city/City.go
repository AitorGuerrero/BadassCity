package city

import (
	"github.com/AitorGuerrero/BadassCity/game/city/neighbourhood"
	"code.google.com/p/go-uuid/uuid"
)

type City interface {}

type Id uuid.UUID

type city struct {
	id Id
	name string
	neighbourhoods []neighbourhood.Neighbourhood
}

func New (name string, neighbourhoods []neighbourhood.Neighbourhood) city {
	return city{
		id: Id(uuid.NewUUID()),
		name: name,
		neighbourhoods: neighbourhoods,
	}
}
