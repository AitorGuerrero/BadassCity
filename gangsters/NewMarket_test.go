package gangsters
import (
	t "testing"
	"github.com/AitorGuerrero/BadassCity/timedEvents/turnsClock"
)

func TestShouldCreateInitialGangsters(t *t.T) {
	clock := &turnsClock.Clock{}
	m := NewMarket(clock)
	if (len(m.Gangsters()) != initialGangstersAmount) {
		t.Error("Should create 20 gangsters. Created: ", len(m.Gangsters()))
	}
}

func TestShouldCreateGangsterInTime(t *t.T) {
	clock := &turnsClock.Clock{}
	m := NewMarket(clock)
	m.gangsters = []*gangster{}
	for i := 0; i < int(gangsterCreationTimeRate) * 2; i++ {
		clock.Next()
	}
	if (len(m.Gangsters()) != 2) {
		t.Error("Should create a gangster after turns")
	}
}