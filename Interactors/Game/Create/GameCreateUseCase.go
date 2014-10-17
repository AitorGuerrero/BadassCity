package Create;

import (
	"github.com/AitorGuerrero/BadassCity/Domain/Entities/Game"
)

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
	game := Game.New(request.Id);
	repository.Save(game);
}
