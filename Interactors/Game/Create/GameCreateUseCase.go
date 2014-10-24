package Create;

import (
	"time"
	"github.com/AitorGuerrero/BadassCity/Domain/Entities/Game"
)

type Request struct {
	Id int
	CreatedTime time.Time
}

type Repository interface {
	Save (Game.Game)
}

var repository Repository;

func Init (comingRepository Repository) {
	repository = comingRepository
}
func Execute (request Request) {
	game := Game.New(request.Id, request.CreatedTime);
	repository.Save(game);
}
