package cities

type Size string

type Price float32

const SizeS Size = "S"
const SizeM Size = "M"
const SizeL Size = "L"
const SizeXL Size = "XL"

type Local struct {
	Price Price
	Size Size
}

type NeighbourhoodConfig struct {
	Name string
	Locals []Local
}

type CityConfig struct {
	Name string
	Neighborhoods []NeighbourhoodConfig
}
