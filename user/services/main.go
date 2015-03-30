package main

import (
	"github.com/AitorGuerrero/BadassCity/user/services/newUser"
	"github.com/AitorGuerrero/BadassCity/user/services/isValid"

	"github.com/koding/kite"
)

var port = 3635

func main() {
	k := kite.New("BadassCity.user", "0.0.1")

	k.HandleFunc("new", func (r *kite.Request) (interface{}, error) {
		args, _ := r.Args.Map()
		name := args["name"].MustString()
		email := args["email"].MustString()
		password := args["password"].MustString()
		return newUser.Service()
	}).DisableAuthentication()

	k.HandleFunc("is-valid", func (r *kite.Request) (interface{}, error) {
			id := r.Args.One().MustString()
			isValid.Service(id)
		}).DisableAuthentication()

	k.Config.Port = Port()
	k.Run()
}

func Port() int {
	return port
}
