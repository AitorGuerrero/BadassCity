package BadassCity

type localOwner interface {
	giveTransaction (transaction)
}

type npcLocalOwner struct {
	w wallet
	ts []transaction
}

func (o *player) giveTransaction (t transaction) {
	o.w.transactions = append(o.w.transactions, t)
}

func (o *npcLocalOwner) giveTransaction (t transaction) {
	o.w.transactions = append(o.w.transactions, t)
}