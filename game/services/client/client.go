package client

import (
	"github.com/AitorGuerrero/BadassCity/game/services"
)

func Get() {
	k := kite.New("BadassCity.game.client", "1.0.0")

	client := k.NewClient("http://localhost:" + services.Port() + "/kite")
	client.Dial()
}
