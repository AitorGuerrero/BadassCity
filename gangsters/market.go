package gangsters

import (
	"github.com/AitorGuerrero/BadassCity/timedEvents"
)

const (
	initialGangstersAmount = 20
	maxGangsterAmountInMarket = 20
	gangsterCreationTimeRate = timedEvents.Day * 3
)

type market struct {
	gangsters []*Gangster
	clock timedEvents.Clock
}

type GangsterNotInTheMarket struct {}
func (GangsterNotInTheMarket) Error () string {
	return "The gangster is not in the market"
}

func NewMarket(clock timedEvents.Clock) *market {
	m := &market{clock: clock}
	m.createInitialGangsters()
	m.initGangsterCreationTicker()
	return m
}

func (m market) Gangsters () []*Gangster {
	return m.gangsters
}

func (m *market) HireGangster(p payer, g *Gangster) error {
	err := m.detachGangster(g)
	if (err != nil) {
		return err
	}
	g.changePayer(p)

	return nil
}

func (m *market) createInitialGangsters () {
	for i := 0; i < initialGangstersAmount ; i++ {
		m.createGangster()
	}
}

func (m *market) createGangster() {
	m.gangsters = append(m.gangsters, &Gangster{clock: m.clock})
}

func (m market) gangsterIndex(g *Gangster) (err error, matchingIndex int) {
	matchingIndex = -1
	for index, sg := range m.gangsters {
		if (sg == g) {
			matchingIndex = index
			break;
		}
	}
	if (matchingIndex == -1) {
		err = GangsterNotInTheMarket{}
	}
	return
}

func (m *market) detachGangster(g *Gangster) error {
	err, gangsterIndex := m.gangsterIndex(g)
	if err != nil {
		return err
	}
	m.gangsters = append(m.gangsters[:gangsterIndex], m.gangsters[gangsterIndex + 1:]...)

	return nil
}

func (m *market) initGangsterCreationTicker() {
	m.clock.AddTicker(gangsterCreationTimeRate, func() {
		if (len(m.gangsters) < maxGangsterAmountInMarket) {
			m.createGangster()
		}
	})
}