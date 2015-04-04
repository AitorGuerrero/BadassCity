package game

import (
	"code.google.com/p/go-uuid/uuid"
	"time"
)

type player interface {}
type city interface {}
type id uuid.UUID
type game struct {
	id id
	initiatedAt time.Time
	startedTime time.Time
	running bool
	players []player
	city city
}

func New (aCity city) *game {
	aGame := &game{
		id: id(uuid.NewUUID()),
		initiatedAt: time.Now(),
		city: aCity,
		running: false,
	}
	return aGame
}

func Get (gameId string) *game {
	return &game {
		id: id(gameId),
	}
}


func (aGame *game) AddPlayer(aPlayer player) {
	aGame.players = append(aGame.players, aPlayer)
}

func (aGame game) Id() id {
	return aGame.id
}
