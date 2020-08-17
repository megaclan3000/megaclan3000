package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// func handlerAPI(w http.ResponseWriter, r *http.Request) {
// 	byt := apiHandler(mux.Vars(r))
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(byt)
// }

func handlerAPIPlayerinfo(w http.ResponseWriter, r *http.Request) {
	log.Debug("API request to:", r.RequestURI)

	var byt []byte

	//TODO use correct id

	var steamID uint64 = 76561198092006615
	vars := mux.Vars(r)

	switch vars["endpoint"] {
	case "weapons":

		info, err := datastorage.GetPlayerInfoBySteamID(steamID)

		//TODO handle error
		if err != nil {
			panic(err)
		}

		byt, err = json.Marshal(info.UserStatsForGame.Stats.WeaponStats())

	default:
		byt = []byte(`Here is info about player: ` + vars["steamid"])
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(byt)
}
