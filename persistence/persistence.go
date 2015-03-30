package persistence

type serializable interface {
	serialize() interface{}
}

var persistLayer *persistLayer

type persistLayer interface {
	Persist (*interface{})
	Flush ()
}

func SetPersistLayer(aPersistLayer *persistLayer) {
	persistLayer = aPersistLayer
}

func Persist(entity serializable) {
	persistLayer.Persist(entity.serialize())
}

func Flush() {
	persistLayer.Flush()
}

func FindCount(id string) {

}
