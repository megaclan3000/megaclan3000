package main

import (
	// "github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// func handlerAPI(w http.ResponseWriter, r *http.Request) {
// 	byt := apiHandler(mux.Vars(r))
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(byt)
// }

func handlerAPIExample(w http.ResponseWriter, r *http.Request) {
	log.Debug("API request to:", r.RequestURI)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`
[
   [ "2020-08-14T14:58:59.279Z", 2.654857302871205 ],
   [ "2020-08-14T14:59:00.279Z", 9.660990312675485 ],
   [ "2020-08-14T14:59:01.279Z", 9.867134407785812 ],
   [ "2020-08-14T14:59:02.279Z", 9.693173367060712 ],
   [ "2020-08-14T14:59:03.279Z", 5.795384070587739 ],
   [ "2020-08-14T14:59:04.279Z", 6.545144917075523 ],
   [ "2020-08-14T14:59:05.279Z", 1.610337192023139 ],
   [ "2020-08-14T14:59:06.279Z", 4.059514075791508 ],
   [ "2020-08-14T14:59:07.279Z", 6.258157166881933 ],
   [ "2020-08-14T14:59:08.279Z", 4.678033662200467 ]
]`))
}
