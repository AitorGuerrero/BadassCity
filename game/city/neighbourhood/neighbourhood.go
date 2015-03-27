package neighbourhood

import (
	"code.google.com/p/go-uuid/uuid"
)

type neighbourhood struct {
	id uuid.UUID
	name string
}

func New (name string) neighbourhood {
	return neighbourhood{
		id: uuid.NewUUID(),
		name: name,
	}
}
