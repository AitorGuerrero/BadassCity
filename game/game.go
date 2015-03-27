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
		id: uuid.NewUUID(),
		initiatedAt: time.Now(),
		city: aCity,
		running: false,
	}
}

func Get (id id) game {
	return game {
		id: id,
	}
}

type game struct {
	id uuid.UUID
	initiatedAt time.Time
	startedTime time.Time
	running bool
	players []player
	city city
}

func (aGame *game) AddPlayer(aPlayer player) {
	aGame.players = append(aGame.players, aPlayer)
}

func (aGame game) Id() uuid.UUID {
	return aGame.id
}
