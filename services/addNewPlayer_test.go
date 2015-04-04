package services

import (
	t "testing"
	"errors"
)

type query struct {
	exists bool
	error error
}

func (q query) Execute(userId string) (bool, error) {
	return q.exists, q.error
}

func (q query) Returns(e bool, er error) {
	q.exists = e
	q.error = er
}

func TestWhenUserIsNotValidShouldReturnAError (t *t.T) {
	gameId := "gameId"
	userId := "userId"
	q := query{false, nil}
	error := AddNewPlayer(gameId, userId, q)
	if error == nil {
		t.Error("Should return a error")
	}
}

func TestWhenUserIsNotValidShouldReturnNil (t *t.T) {
	gameId := "gameId"
	userId := "userId"
	q := query{true, nil}
	error := AddNewPlayer(gameId, userId, q)
	if error != nil {
		t.Error("Should return nil as error")
	}
}

func TestWhenServiceReturnsAErrorShouldReturnServerError (t *t.T) {
	gameId := "gameId"
	userId := "userId"
	expectedError := "Server error"
	q := query{true, errors.New("aError")}
	error := AddNewPlayer(gameId, userId, q)
	if error.Error() != expectedError {
		t.Error("Should return correct error message")
	}
}
