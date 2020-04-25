package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	// "github.com/davecgh/go-spew/spew"
)

var config SteamConfig

func main() {

	config = readConfig()
	config.Refresh()

	r := mux.NewRouter()
	r.HandleFunc("/", handlerStats)
	r.HandleFunc("/player/{id}", handlerDetails)
	r.NotFoundHandler = http.HandlerFunc(handler404)

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}

func handlerStats(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/stats.html")
	data := config.GetAll()
	t.Execute(w, data)
}

func handler404(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/404.html")
}

func handlerDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	players := config.GetAll()

	for _, p := range players {
		if vars["id"] == p.PlayerSummary.Steamid {
			t, _ := template.ParseFiles("templates/details.html")
			t.Execute(w, p)
			return
		}
	}
	http.ServeFile(w, r, "templates/404.html")
}
