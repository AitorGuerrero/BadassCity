package newUser

import (
	"github.com/AitorGuerrero/BadassCity/user"

	"github.com/koding/kite"
)

func Service(r *kite.Request) (interface{}, error) {
	args, _ := r.Args.Map()
	name := args["name"].MustString()
	email := args["email"].MustString()
	password := args["password"].MustString()
	aUser := user.New(name, email, password)
	result := aUser.Id()
	return result, nil
}
