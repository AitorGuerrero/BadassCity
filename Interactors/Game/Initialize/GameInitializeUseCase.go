package Initialize;

import "github.com/AitorGuerrero/BadassCity/Domain/Entities/Game"

type Request struct {
	Id int
}

type UseCase struct {
	repository Repository
}

type Repository interface {
	Find (id int) Game.Game
	Persist (game Game.Game)
}

func (this UseCase) Execute (request Request) {
	game := this.repository.Find(request.Id);
	game.Initialize();
	this.repository.Persist(game);
}

func New(repository Repository) *UseCase {
	return &UseCase{repository: repository}
}
