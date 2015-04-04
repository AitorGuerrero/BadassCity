package kite

import (
	"github.com/AitorGuerrero/BadassCity/services/new"

	"github.com/koding/kite"
)

func AddService(k *kite.Kite) {
	k.HandleFunc("new", func (r *kite.Request) (interface{}, error) {
			c := r.Args.One().MustString()
			return new.Service(c)
		}).DisableAuthentication()
}
