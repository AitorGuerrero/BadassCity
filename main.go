package main

import (
	"github.com/AitorGuerrero/BadassCity/config"
)

func main() {
	cfg := config.Get()
	initServices(cfg.KiteServiceConfig)
}
