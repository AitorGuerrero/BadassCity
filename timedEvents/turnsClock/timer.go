package turnsClock

type timer struct {
	remainingTurns turn
	callback func()
}

func (t *timer) next() (expired bool) {
	t.remainingTurns--
	if (t.remainingTurns <= 0) {
		t.callback()
		expired = true
	}

	return
}