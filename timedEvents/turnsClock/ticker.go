package turnsClock

type ticker struct {
	turnsForCallback turn
	remainingTurns turn
	callback func()
}

func (t *ticker) next() {
	t.remainingTurns--
	if (t.remainingTurns <= 0) {
		t.callback()
		t.remainingTurns = t.turnsForCallback
	}
}
