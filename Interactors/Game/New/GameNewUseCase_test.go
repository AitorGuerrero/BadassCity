package New;

import (
	"testing"
	"github.com/AitorGuerrero/BadassCity/Domain/Entities/Game"
)

type repositoryStub struct {
	T *testing.T
	SavedCalled bool
}

func (this *repositoryStub) Save(game Game.Game) {
	this.SavedCalled = true;
	if game.Id != 5 {
		this.T.Error("Saving id incorrect");
	}

}

func TestCreateNewGame (t *testing.T) {
	request := Request{}
	request.Id = 5
	repository := repositoryStub{T: t, SavedCalled: false}
	Init(&repository)
	Execute(request)
	if repository.SavedCalled == false {
		t.Error("The use case should save");
	}
}
