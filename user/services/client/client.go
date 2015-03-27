package client

import (
	"github.com/AitorGuerrero/BadassCity/user/services"
)

func Get() {
	k := kite.New("BadassCity.user.client", "1.0.0")

	client := k.NewClient("http://localhost:" + services.Port() + "/kite")
	client.Dial()
}
