package isValid

type Query interface {
	Execute(uId string) (bool, error)
}

type isValid struct {}

func New () Service {
	return &isValid{}
}

func (q *isValid) Execute(uId string) (bool, error) {
	r, _ := Client().Tell("is-valid", uId)
	return r.MustBool(), nil
}
