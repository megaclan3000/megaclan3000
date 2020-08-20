package main

import (
	// "github.com/gorilla/mux"
	"encoding/json"
	"net/http"
)

// func handlerAPI(w http.ResponseWriter, r *http.Request) {
// 	byt := apiHandler(mux.Vars(r))
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(byt)
// }

func handlerAPIClaninfo(w http.ResponseWriter, r *http.Request) {

	var byt []byte

	//TODO use correct id

	// vars := mux.Vars(r)

	var err error

	byt, err = json.Marshal(datastorage.GetMatches())

	//TODO handle error
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(byt)
}
