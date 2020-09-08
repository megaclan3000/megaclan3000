package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerAPIClaninfo(w http.ResponseWriter, r *http.Request) {

	var byt []byte
	var err error

	vars := mux.Vars(r)

	switch vars["endpoint"] {
	case "matches":
		if byt, err = json.Marshal(datastorage.GetMatches()); err != nil {
			panic(err)
		}
	case "players":
		if byt, err = json.Marshal(datastorage.GetPlayers()); err != nil {
			panic(err)
		}
	case "updates":
		if byt, err = json.Marshal(datastorage.GetUpdates()); err != nil {
			panic(err)
		}
	default:
		if byt, err = json.Marshal(
			struct {
				Matches interface{} `json:"matches"`
				Players interface{} `json:"players"`
				Updates interface{} `json:"updates"`
			}{
				Matches: datastorage.GetMatches(),
				Players: datastorage.GetPlayers(),
				Updates: datastorage.GetUpdates(),
			}); err != nil {
			panic(err)
		}
	}
	// byt, err = json.Marshal(datastorage.GetMatches())

	w.Header().Set("Content-Type", "application/json")
	w.Write(byt)
}
