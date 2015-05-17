package addNewPlayer

import (
	t "testing"
	"github.com/AitorGuerrero/BadassCity/game"
)

type userId struct {}
func (userId) toString() string {
	return "aaa"
}

type gameId struct {}
func (gameId) toString() string {
	return "aaa"
}

type userIsValidQuery struct {}
func (userIsValidQuery) Execute(userId string) (bool, error) {
	return true, nil
}

type userIsNotValidQuery struct {}
func (userIsNotValidQuery) Execute(userId string) (bool, error) {
	return false, nil
}

func TestWhenUserIsNotValidShouldReturnAError (t *t.T) {
	gameId := gameId{}
	userId := userId{}
	q := userIsNotValidQuery{}
	s := New(q)
	error := s.Execute(gameId, userId)
	if error == nil {
		t.Error("Should return a error")
	}
}

func TestWhenUserIsValidShouldReturnNil (t *t.T) {
	gameId := gameId{}
	userId := userId{}
	q := userIsValidQuery{}
	s := New(q)
	error := s.Execute(gameId, userId)
	if error != nil {
		t.Error("Should return nil as error")
	}
}

func TestWhenAllIsOkTheUserShouldBeInTheGame(t *t.T) {

}
