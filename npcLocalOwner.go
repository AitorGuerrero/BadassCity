package BadassCity
import "errors"

type npcLocalOwner struct {
	w wallet
	ts []transaction
}

func (o *npcLocalOwner) giveTransaction(float32) (err error, t transaction) {
	err = errors.New("npc local owner can't give transactions")
	return
}

func (o *npcLocalOwner) getTransaction (t transaction) {
	o.w.transactions = append(o.w.transactions, t)
}