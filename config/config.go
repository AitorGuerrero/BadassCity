package config

import (
	userQueries "github.com/AitorGuerrero/BadassCity/user/queries"
)

func Get() Config {
	return Config {
		userQueries.Config {
			Server: "BadassCity.user.client",
			Version: "0.0.1",
			Port: "3635",
		},
		KiteServiceConfig {
			Name: "BaddassCity.game",
			Version: "0.0.0",
			Port: 3634,
		},
	}
}
