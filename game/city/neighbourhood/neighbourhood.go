package neighbourhood

import (
	"code.google.com/p/go-uuid/uuid"
)

type Neighbourhood interface {}
type Id uuid.UUID

type neighbourhood struct {
	id Id
	name string
}

func New (name string) neighbourhood {
	return neighbourhood{
		id: Id(uuid.NewUUID()),
		name: name,
	}
}
