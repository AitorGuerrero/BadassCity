package gangsters
import t "testing"

func TestShouldCreateInitialGangsters(t *t.T) {
	m := NewMarket()
	if (len(m.Gangsters()) != initialGangstersAmount) {
		t.Error("Should create 20 gangsters. Created: ", len(m.Gangsters()))
	}
}
