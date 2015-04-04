package queries

type IsValidInterface interface {
	Execute(uId string) (bool, error)
}

type isValid struct {}

func IsValid () IsValidInterface {
	return &isValid{}
}

func (q *isValid) Execute(uId string) (bool, error) {
	r, _ := Client().Tell("is-valid", uId)
	return r.MustBool(), nil
}
