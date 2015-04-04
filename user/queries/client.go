package queries

import (
	commonClient "github.com/AitorGuerrero/BadassCity/common/services/client"

	"github.com/koding/kite"
)

type Config struct {
	Server string
	Version string
	Port string
}

func Client() *kite.Client {
	return commonClient.Get("BadassCity.user.client", "0.0.1", "3635")
}
