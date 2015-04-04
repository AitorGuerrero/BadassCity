package services

import (
	t "testing"
	"errors"
)

type queryFake struct {
	exists bool
	error error
}

func (q queryFake) Execute(userId string) (bool, error) {
	return q.exists, q.error
}

func (q queryFake) Returns(e bool, er error) {
	q.exists = e
	q.error = er
}

func TestWhenUserIsNotValidShouldReturnAError (t *t.T) {
	gameId := "gameId"
	userId := "userId"
	q := queryFake{false, nil}
	error := AddNewPlayer(gameId, userId, q)
	if error == nil {
		t.Error("Should return a error")
	}
}

func TestWhenUserIsValidShouldReturnNil (t *t.T) {
	gameId := "gameId"
	userId := "userId"
	q := queryFake{true, nil}
	error := AddNewPlayer(gameId, userId, q)
	if error != nil {
		t.Error("Should return nil as error")
	}
}

func TestWhenServiceReturnsAErrorShouldReturnServerError (t *t.T) {
	gameId := "gameId"
	userId := "userId"
	expectedError := "Server error"
	q := queryFake{true, errors.New("aError")}
	error := AddNewPlayer(gameId, userId, q)
	if error.Error() != expectedError {
		t.Error("Should return correct error message")
	}
}
