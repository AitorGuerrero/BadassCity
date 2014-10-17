package Game;

type Game interface {
	Id() int
}

type game struct {
	id int
}

func (this game) Id () int {
	return this.id;
}

func New (id int) game {
	game := game{id : id};
	return game
}
