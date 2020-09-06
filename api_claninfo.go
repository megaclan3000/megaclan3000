package main

import (
	"net/http"
)

func handlerAPIClaninfo(w http.ResponseWriter, r *http.Request) {

	var byt []byte

	//TODO query database for vars := mux.Vars(r)
	// byt, err = json.Marshal(datastorage.GetMatches())

	w.Header().Set("Content-Type", "application/json")
	w.Write(byt)
}
