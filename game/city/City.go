package city

import (
	"code.google.com/p/go-uuid/uuid"
)

type City interface {

}

type Neighbourhood interface {

}

type city struct {
	id uuid.UUID
	name string
	neighbourhoods []Neighbourhood
}

func New (name string, neighbourhoods []Neighbourhood) city {
	return city{
		id: uuid.NewUUID(),
		name: name,
		neighbourhoods: neighbourhoods,
	}
}
