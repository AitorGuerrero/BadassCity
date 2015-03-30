package client

import (
	"github.com/koding/kite"
)

func Get(name, version, port string) *kite.Client {
	k := kite.New(name, version)

	client := k.NewClient("http://localhost:" + port + "/kite")
	client.Dial()

	return client
}
