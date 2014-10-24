package Create;

import (
	"testing"
	"time"
	"github.com/AitorGuerrero/BadassCity/Domain/Entities/Game"
)

type repositoryStub struct {
	T *testing.T
	SavedCalled bool
	SavedGame Game.Game
}

func (this *repositoryStub) Save(game Game.Game) {
	this.SavedCalled = true;
	this.SavedGame = game;
}

func TestCreateNewGame (t *testing.T) {
	request := Request{}
	request.Id = 5
	request.CreatedTime = time.Now()
	repository := repositoryStub{T: t, SavedCalled: false}
	Init(&repository)
	Execute(request)
	if repository.SavedCalled == false {
		t.Error("The use case should save");
	}
	if repository.SavedGame.Id() != 5 {
		t.Error("Saving id incorrect");
	}
	t.Log(repository.SavedGame.CreatedTime())
	if repository.SavedGame.CreatedTime() != request.CreatedTime {
		t.Error("Should save the correct created time");
	}
}
