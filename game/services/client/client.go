package client

import (
	"github.com/AitorGuerrero/BadassCity/game/services"
	commonClient "github.com/AitorGuerrero/BadassCity/common/services/client"

	"github.com/koding/kite"
)

func Get() kite.Client {
	return commonClient.Get("BadassCity.game.client", "0.0.1", services.Port())
}
