package queries

import (
	commonClient "github.com/AitorGuerrero/BadassCity/common/services/client"
	"github.com/AitorGuerrero/BadassCity/config"

	"github.com/koding/kite"
)

func Client() *kite.Client {
	return commonClient.Get(
		config.UserService.Server,
		config.UserService.Version,
		config.UserService.Port,
	)
}
