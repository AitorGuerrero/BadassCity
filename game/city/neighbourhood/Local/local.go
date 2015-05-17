package local

import (
	"uuid"
)

const SizeS size = "S"
const SizeM size = "M"
const SizeL size = "L"
const SizeXL size = "XL"

type Local interface {}

type size string
type price float32
type id uuid.UUID
type local struct {
	id id
	size size
	price price
}

func NewId() id {
	return id(uuid.NewUUID())
}

func New (i id, s size, p price) Local {
	l := local{i, s, p}
	return &l
}
