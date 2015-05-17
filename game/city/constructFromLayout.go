package city

import (
	"github.com/AitorGuerrero/BadassCity/game/city/layout"
	"github.com/AitorGuerrero/BadassCity/game/city/neighbourhood"
	"github.com/AitorGuerrero/BadassCity/game/city/neighbourhood/local"
)

func ConstructFromLayout(l *layout.Layout) *City {
	neighbourhoods := []neighbourhood.Neighbourhood{}
	for _, neighborhoodConfig := range l.Neighborhoods {
		neighbourhoods = append(
			neighbourhoods,
			neighbourhood.New(
				neighbourhood.NewId(),
				neighborhoodConfig.Name,
			),
		)
	}
	i := NewId()
	c := New(i, l.Name, neighbourhoods)
	return &c
}

func constructNeighbourhoods (nsl []layout.Neighbourhood) []neighbourhood.Neighbourhood {
	ns := []neighbourhood.Neighbourhood{}
	for _, nl := range nsl {
		ls := constructLocals(nl.Locals)
		ns = append(
			ns,
			neighbourhood.New(
				neighbourhood.NewId(),
				nl.Name,
				ls,
			),
		)
	}

	return ns
}

func constructLocals (lsl []Layout.Local) []local.Local {
	ls := []local.Local{}
	for _, ll := range lsl {
		ls = append(
			ls,
			local.New(
				local.NewId(),
				ll.Name,
				ll.Size,
				ll.Price,
			),
		)
	}

	return ns
}
