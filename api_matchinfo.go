package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func handlerAPIMatchinfo(w http.ResponseWriter, r *http.Request) {
	log.Debug("API request to:", r.RequestURI)

	var byt []byte
	var err error

	vars := mux.Vars(r)

	matchInfo, err := datastorage.GetMatchByID(vars["matchid"])
	if err != nil {
		log.Warning(err)
		return
	}
	switch vars["endpoint"] {

	case "scoreboard":
		if byt, err = json.Marshal(matchInfo.GetScoreboard()); err != nil {
			log.Warning(err)
			return
		}
	case "rounds":
		if byt, err = json.Marshal(matchInfo.Rounds); err != nil {
			log.Warning(err)
			return
		}

	case "weapons":
		if byt, err = json.Marshal(matchInfo.Weapons()); err != nil {
			log.Warning(err)
			return
		}
	case "duels":

		if byt, err = json.Marshal(matchInfo.Damages()); err != nil {
			log.Warning(err)
			return
		}
	case "heatmaps":
		//TODO
	case "megacoins":
		//TODO
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(byt)
	if err != nil {
		log.Error(err)
	}
}
