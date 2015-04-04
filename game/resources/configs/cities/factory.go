package cities

import (
	"errors"
)

func Factory(name string) (*CityConfig, error) {
	switch name {
	case "test":
		return &TestCityConfig, nil
	default:
		return CityConfig{}, errors.New("City type does not exists")
	}
}
