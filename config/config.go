package config

import (
	userQueries "github.com/AitorGuerrero/BadassCity/user/queries"
)

var UserService = userQueries.Config {
	Server: "BadassCity.user.client",
	Version: "0.0.1",
	Port: "3635",
}

var Kite = KiteServiceConfig {
	Name: "BaddassCity.game",
	Version: "0.0.0",
	Port: 3634,
}
