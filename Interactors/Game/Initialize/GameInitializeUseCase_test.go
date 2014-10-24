package Initialize;

import (
	"testing"
)

func TestGameInitialized (t *testing.T) {
	request := Request{}
	request.Id = 5
	repository := repositoryMock{}
	useCase := New(&repository);
	useCase.Execute(request);
	if repository.SavedCalled == false {
		t.Error("The use case should save the game");
	}
	if repository.SavedGame.Id() != 5 {
		t.Error("Saving id incorrect");
	}
	if repository.SavedGame.Initialized() != true {
		t.Error("Should have been initialized");
	}
}
