package main

import (
	// "github.com/gorilla/mux"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// func handlerAPI(w http.ResponseWriter, r *http.Request) {
// 	byt := apiHandler(mux.Vars(r))
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(byt)
// }

func handlerAPIMatchinfo(w http.ResponseWriter, r *http.Request) {
	log.Debug("API request to:", r.RequestURI)

	var byt []byte
	var err error

	vars := mux.Vars(r)

	matchInfo, err := datastorage.GetMatchByID(vars["matchid"])

	if err != nil {
		panic(err)
	}
	switch vars["endpoint"] {

	case "scoreboard":
		if byt, err = json.Marshal(matchInfo.GetScoreboard()); err != nil {
			panic(err)
		}

	case "rounds":
		if byt, err = json.Marshal(matchInfo.Rounds); err != nil {
			panic(err)
		}

	// TODO
	case "weapons":
		if byt, err = json.Marshal(matchInfo.Weapons()); err != nil {
			panic(err)
		}
	// TODO
	case "duels":

		if byt, err = json.Marshal(matchInfo.Damages()); err != nil {
			panic(err)
		}
		// byt = []byte(`
		// [
	// "PlayerClan1": [
	// {
	// 	"PlayerEnemy1": "40",
	// 	"PlayerEnemy2": "41",
	// 	"PlayerEnemy3": "42",
	// 	"PlayerEnemy4": "43",
	// 	"PlayerEnemy5": "44"
	// },
	// "PlayerClan2": [
	// {
	// 	"PlayerEnemy1": "40",
	// 	"PlayerEnemy2": "41",
	// 	"PlayerEnemy3": "42",
	// 	"PlayerEnemy4": "43",
	// 	"PlayerEnemy5": "44"
	// },
	// "PlayerClan3": [
	// {
	// 	"PlayerEnemy1": "40",
	// 	"PlayerEnemy2": "41",
	// 	"PlayerEnemy3": "42",
	// 	"PlayerEnemy4": "43",
	// 	"PlayerEnemy5": "44"
	// },
	// "PlayerClan4": [
	// {
	// 	"PlayerEnemy1": "40",
	// 	"PlayerEnemy2": "41",
	// 	"PlayerEnemy3": "42",
	// 	"PlayerEnemy4": "43",
	// 	"PlayerEnemy5": "44"
	// },
	// "PlayerClan5": [
	// {
	// 	"PlayerEnemy1": "40",
	// 	"PlayerEnemy2": "41",
	// 	"PlayerEnemy3": "42",
	// 	"PlayerEnemy4": "43",
	// 	"PlayerEnemy5": "44"
	// }
	// ]
	// `)
	//TODO
	case "heatmaps":
		//TODO
	case "megacoins":
		//TODO
	default:
		byt = []byte(`
[
	["PlayerA1", "PlayerB1", 5],
	["PlayerA1", "PlayerB2", 5],
	["PlayerA1", "PlayerB3", 5],
	["PlayerA1", "PlayerB4", 5],
	["PlayerA1", "PlayerB5", 5],
	["PlayerA2", "PlayerB1", 5],
	["PlayerA2", "PlayerB2", 5],
	["PlayerA2", "PlayerB3", 5],
	["PlayerA2", "PlayerB4", 5],
	["PlayerA2", "PlayerB5", 5],
	["PlayerA3", "PlayerB1", 5],
	["PlayerA3", "PlayerB2", 5],
	["PlayerA3", "PlayerB3", 5],
	["PlayerA3", "PlayerB4", 5],
	["PlayerA3", "PlayerB5", 5],
	["PlayerA4", "PlayerB1", 5],
	["PlayerA4", "PlayerB2", 5],
	["PlayerA4", "PlayerB3", 5],
	["PlayerA4", "PlayerB4", 5],
	["PlayerA4", "PlayerB5", 5],
	["PlayerA5", "PlayerB1", 5],
	["PlayerA5", "PlayerB2", 5],
	["PlayerA5", "PlayerB3", 5],
	["PlayerA5", "PlayerB4", 5],
	["PlayerA5", "PlayerB5", 5]
]`)

	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(byt)
}
