package game

import (
	t "testing"
)

func TestWhenCreatingANewGameShouldReturnAGameWithTheProvidedId (t *t.T) {
	i := NewId()
	g := New(i)

	if g.Id() != i {
		t.Error("The Id does not Match")
	}
}

func TestWhenCreatingANewGameShouldNotBeRunning (t *t.T) {
	i := NewId()
	g := New(i)

	if false != g.running {
		t.Error("The game is running")
	}
}
