package BadassCity

type local struct {
	p float32
	o localOwner
}

func (l local) price() float32 {
	return l.p
}

func (l local) owner() localOwner {
	return l.o
}

func (l *local) changeOwner(o localOwner) {
	l.o = o
}