package user

type serializedUser struct {
	Id string
	Name string
	Email string
	AuthHash string
}

func (aUser *user) serialize () interface{} {
	return &serializedUser {
		Id: aUser.id,
		Name: aUser.name,
		Email: aUser.email,
		AuthHash: string(aUser.authHash),
	}
}

func (aUser *serializedUser) unserialize () *user {
	return &user {
		id: id(aUser.Id),
		name: aUser.Name,
		email: aUser.Email,
		authHash: authHash(aUser.AuthHash),
	}
}
