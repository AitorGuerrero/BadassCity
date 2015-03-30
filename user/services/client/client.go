package client

import (
	commonClient "github.com/AitorGuerrero/BadassCity/common/services/client"

	"github.com/koding/kite"
)

func Get() *kite.Client {
	return commonClient.Get("BadassCity.user.client", "0.0.1", "3635")
}