package game

import (
	"time"

	"code.google.com/p/go-uuid/uuid"
)

type player interface {}
type city interface {}
type id uuid.UUID

func New (aCity city) game {
	return game{
		id: id(uuid.NewUUID()),
		initiatedAt: time.Now(),
		city: aCity,
		running: false,
	}
}

func Get (gameId string) *game {
	return &game {
		id: id(gameId),
	}
}

type game struct {
	id id
	initiatedAt time.Time
	startedTime time.Time
	running bool
	players []player
	city city
}

func (aGame *game) AddPlayer(aPlayer player) {
	aGame.players = append(aGame.players, aPlayer)
}

func (aGame game) Id() id {
	return aGame.id
}
