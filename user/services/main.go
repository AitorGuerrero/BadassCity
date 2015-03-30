package main

import (
	"github.com/AitorGuerrero/BadassCity/user/services/newUser"
	"github.com/AitorGuerrero/BadassCity/user/services/isValid"

	"github.com/koding/kite"
)

var port = 3635

func main() {
	k := kite.New("BadassCity.user", "0.0.1")
	k.HandleFunc("new", newUser.Service).DisableAuthentication()
	k.HandleFunc("is-valid", isValid.Service).DisableAuthentication()

	k.Config.Port = Port()
	k.Run()
}

func Port() int {
	return port
}
