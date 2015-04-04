package game

import (
	"code.google.com/p/go-uuid/uuid"
	"time"
)

type player interface {}
type city interface {}
type Id uuid.UUID
type game struct {
	id Id
	initiatedAt time.Time
	startedTime time.Time
	running bool
	players []player
	city city
}

func New (aCity city) *game {
	aGame := &game{
		id: Id(uuid.NewUUID()),
		initiatedAt: time.Now(),
		city: aCity,
		running: false,
	}
	return aGame
}

func Get (gameId string) *game {
	return &game {
		id: Id(gameId),
	}
}


func (aGame *game) AddPlayer(aPlayer player) {
	aGame.players = append(aGame.players, aPlayer)
}

func (aGame game) Id() Id {
	return aGame.id
}
