package game

import(
	"time"
)

type serializedGame struct {
	Id string
	InitiatedAt time.Time
	StartedTime time.Time
	Running bool
}

func (aGame *game) serialize () interface{} {
	return &serializedGame {
		Id: string(aGame.id),
		InitiatedAt: aGame.initiatedAt,
		StartedTime: aGame.startedTime,
		Running: aGame.running,
	}
}

func (aGame *serializedGame) unserialize () *game {
	return &game {
		id: id(aGame.Id),
		initiatedAt: aGame.InitiatedAt,
		startedTime: aGame.StartedTime,
		running: aGame.Running,
	}
}
