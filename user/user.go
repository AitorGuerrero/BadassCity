package user

import (
	"github.com/AitorGuerrero/BadassCity/persistence"

	"code.google.com/p/go-uuid/uuid"
)

type authHash string

type user struct {
	id uuid.UUID
	name string
	email string
	authHash authHash
}

func New (name string, email string, password string) user {
	return user{
		name: name,
		email: email,
		authHash: generateAuthHash(name, password),
	}
}

func (aUser user) Id() uuid.UUID {
	return aUser.id
}

func generateAuthHash (name string, password string) authHash {
	return authHash(name + password)
}

func ExistsUserWithId (id string) bool {
	entity :=
	return persistence.findCount(id) > 1
}
