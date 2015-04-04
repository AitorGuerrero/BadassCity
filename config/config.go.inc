package config

func Get() Config {
	return Config {
		MongoDBConfig {
			Server: "mongodb://user:password@server:port/database",
		},
		KiteServiceConfig {
			Name: "BaddassCity.user",
			Version: "0.0.0",
			Port: 3635,
		},
	}
}
