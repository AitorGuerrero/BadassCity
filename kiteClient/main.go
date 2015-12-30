package main

//import (
//	"github.com/koding/kite"
//	"net/url"
//	"github.com/AitorGuerrero/BadassCity"
//)

func main() {
//	playersCollection := BadassCity.NewPlayerCollection()
//	city := BadassCity.NewCity()
//
//	k := kite.New("badassCity", "0.0.0")
//	k.Config.Port = 6000
//	k.Config.DisableAuthentication = true
//	k.HandleFunc("addPlayerToGame", func(r *kite.Request) (o interface{}, err error) {
//		userSerializedId := r.Args.One().MustString()
//		BadassCity.NewPlayer(userSerializedId)
//		return
//	})
//	k.HandleFunc("buyLocal", func(r *kite.Request) (o interface{}, err error) {
//		args := r.Args.MustMap()
//		playerId := args["userId"].MustString()
//		localId := args["localId"].MustString()
//		neighbourhoodId := args["neighbourhoodId"].MustString()
//		pl := playersCollection.Get(playerId)
//		return
//	})
//	k.Register(&url.URL{Scheme: "http", Host: "localhost:6000/kite"})
//	k.Run()
//	print("FINISHED")
}