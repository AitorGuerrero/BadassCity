package isValid

import(
	"github.com/AitorGuerrero/BadassCity/user/queries"
)

type Query interface {
	Execute(uId string) (bool, error)
}

type isValid struct {}

func New () Query {
	return &isValid{}
}

func (q *isValid) Execute(uId string) (bool, error) {
	r, _ := queries.Client().Tell("is-valid", uId)
	return r.MustBool(), nil
}
