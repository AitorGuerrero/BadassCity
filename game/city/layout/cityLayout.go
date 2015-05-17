package layout

import (
	"github.com/AitorGuerrero/BadassCity/game/city/layout/neighbourhood"
)

type Size string

type Price float32

const SizeS Size = "S"
const SizeM Size = "M"
const SizeL Size = "L"
const SizeXL Size = "XL"

type City struct {
	Name string
	Neighbourhoods []neighbourhood.NeighbourhoodLayout
}

type Neighbourhood struct {
	Name string
	Locals []Local
}

type Local struct {
	Size Size
	Price Price
}
