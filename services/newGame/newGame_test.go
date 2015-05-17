package newGame

import (
	t "testing"
	"github.com/AitorGuerrero/BadassCity/game"
)

type testGameRepository struct {
	t *t.T
	c []game.Game
}

func (gr *testGameRepository) count() int {
	return len(gr.c)
}

func (gr *testGameRepository) Add(aGame game.Game) {
	gr.c = append(gr.c, aGame)
}

func TestWhenAllIsOkShouldPersistAGame (t *t.T) {
	r := testGameRepository{t, []game.Game{}}
	clr := testCityLayoutRepository{}
	cli := 1
	s := New(&r, &clr)
	s.Execute(cl1)
	amount := r.count()
	if amount == 0 {
		t.Error("There should be add the game to the collection")
	} else if amount != 1 {
		t.Error("There should be only one item in the collection")
	}
}
