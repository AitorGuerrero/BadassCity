package config

import (
	userQueries "github.com/AitorGuerrero/BadassCity/user/queries"
)

type KiteServiceConfig struct {
	Name string
	Version string
	Port int
}

type MongoDBConfig struct {
	Server string
}

type Config struct {
	userQueries.Config
	KiteServiceConfig
}
