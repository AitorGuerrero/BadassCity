package New;

import "github.com/AitorGuerrero/BadassCity/Domain/Entities/Game"

type Request struct {
	Id int
}

type Repository interface {
	Save (Game.Game)
}

var repository Repository;

func Init (comingRepository Repository) {
	repository = comingRepository
}
func Execute (request Request) {
	var game Game.Game;
	game = Game.Game{Id: request.Id};
	repository.Save(game);
}
