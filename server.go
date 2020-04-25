package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var config SteamConfig
var t *template.Template

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

	// Parse all templates
	var err error
	t, err = template.ParseGlob("./templates/*")
	if err != nil {
		log.Println("Cannot parse templates:", err)
		os.Exit(-1)
	}

	log.Fatal(srv.ListenAndServe())
}

func handlerStats(w http.ResponseWriter, r *http.Request) {
	data := config.GetAll()
	t.ExecuteTemplate(w, "stats.html", data)
}

func handler404(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "404.html", nil)
}

func handlerDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	players := config.GetAll()

	for _, p := range players {
		if vars["id"] == p.PlayerSummary.Steamid {
			t.ExecuteTemplate(w, "details.html", p)
			return
		}
	}
	t.ExecuteTemplate(w, "404.html", nil)
}
