package Game;

import "time"

type Game interface {
	Id() int
	CreatedTime() time.Time
	Initialize()
	Initialized() bool
}

type game struct {
	id int
	createdTime time.Time
	initialized bool
}

func (this game) Id () int {
	return this.id;
}

func (this game) CreatedTime () time.Time {
	return this.createdTime;
}

func (this game) Initialized () bool {
	return this.initialized;
}

func (game *game) Initialize () {
	game.initialized = true;
}

func New (id int, createdTime time.Time) *game {
	game := &game{
		id : id,
		createdTime: createdTime,
		initialized: false,
	};
	return game
}
