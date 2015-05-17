package newGame

import (
	"github.com/AitorGuerrero/BadassCity/game"
	"github.com/AitorGuerrero/BadassCity/game/city"
	"github.com/AitorGuerrero/BadassCity/game/city/layout"
)

type gameCollection interface {
	Add(aGame game.Game)
}

type cityLayoutRepository interface {
	Get(id string) layout.City
}

func New(gc gameCollection, clr cityLayoutRepository) *service {
	return &service{gc, clr}
}

type service struct {
	gc gameCollection
	clr cityLayoutRepository
}

func (s *service) Execute (cityLayoutId string) {
	cl := s.clr.Get(cityLayoutId)
	c := city.NewFromLayout(
		city.NewId(),
		cl,
	)
	aGame := game.New(c)
	s.gc.Add(aGame)
}
