package client

import (
	"github.com/AitorGuerrero/BadassCity/user/services"
	commonClient "github.com/AitorGuerrero/BadassCity/common/services/client"

	"github.com/koding/kite"
)

func Get() kite.Client {
	return commonClient.Get("BadassCity.user.client", "0.0.1", services.Port())
}
