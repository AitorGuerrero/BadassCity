package config

type MainConfig struct {
	CitiesConfigPath string
}

type KiteServiceConfig struct {
	Name string
	Version string
	Port int
}

type MongoDBConfig struct {
	Server string
}

type UserServiceConfig struct {
	Server string
	Version string
	Port string
}
