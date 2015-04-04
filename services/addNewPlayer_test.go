package services

import (
	t "testing"
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
	if error != nil {
		t.Error("Should return a error")
	}
}
