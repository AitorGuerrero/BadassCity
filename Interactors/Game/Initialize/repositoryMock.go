package Initialize;

import (
	"time"
	"github.com/AitorGuerrero/BadassCity/Domain/Entities/Game"
)
type repositoryMock struct {
	SavedGame Game.Game
	SavedCalled bool
}
func (this repositoryMock) Find(id int) Game.Game {
	return Game.New(id, time.Now());
}
func (this *repositoryMock) Persist(game Game.Game) {
	this.SavedCalled = true;
	this.SavedGame = game;
}
