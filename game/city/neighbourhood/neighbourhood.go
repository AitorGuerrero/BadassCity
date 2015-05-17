package neighbourhood

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/AitorGuerrero/BadassCity/game/city/neighbourhood/local"
)

type Neighbourhood interface {}
type id uuid.UUID

type neighbourhood struct {
	id id
	name string
	locals []local.Local
}

func New (i id, n string, ls []local.Local) *neighbourhood {
	nb := neighbourhood{i, n, ls}
	return &nb
}

func NewId () id {
	return id(uuid.NewUUID())
}
