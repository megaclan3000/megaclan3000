package main

import (
	// "github.com/gorilla/mux"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// func handlerAPI(w http.ResponseWriter, r *http.Request) {
// 	byt := apiHandler(mux.Vars(r))
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(byt)
// }

func handlerAPIMatchinfo(w http.ResponseWriter, r *http.Request) {
	log.Debug("API request to:", r.RequestURI)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`
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
]`))
}
