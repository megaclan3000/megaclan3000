package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// func handlerAPI(w http.ResponseWriter, r *http.Request) {
// 	byt := apiHandler(mux.Vars(r))
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(byt)
// }

func handlerAPIPlayerinfo(w http.ResponseWriter, r *http.Request) {
	log.Debug("API request to:", r.RequestURI)

	var byt []byte

	vars := mux.Vars(r)
	steamID, err := strconv.ParseUint(vars["steamid"], 10, 64)

	if err != nil {
		//TODO handle error
		panic(err)
	}

	switch vars["endpoint"] {
	case "maps":

		info, err := datastorage.GetPlayerInfoBySteamID(steamID)

		//TODO handle error
		if err != nil {
			panic(err)
		}

		byt, err = json.Marshal(info.UserStatsForGame.Stats.MapStats())
	case "weapons":

		info, err := datastorage.GetPlayerInfoBySteamID(steamID)

		//TODO handle error
		if err != nil {
			panic(err)
		}

		byt, err = json.Marshal(info.UserStatsForGame.Stats.WeaponStats())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(byt)
}
