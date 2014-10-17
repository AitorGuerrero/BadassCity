package Game;

type Game struct {
	Id int
}

func New (id int) Game {
	game := Game{Id : id};
	return game
}
