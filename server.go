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

	// Read config and pull initial data
	config = readConfig()
	config.Refresh()

	r := mux.NewRouter()

	// Serve all static files in public directory
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	// Define routes
	r.HandleFunc("/", handlerStats)
	r.HandleFunc("/contact", handlerContact)
	r.HandleFunc("/faq", handlerFAQ)
	r.HandleFunc("/player/{id}", handlerDetails)

	// Set custom 404 page
	r.NotFoundHandler = http.HandlerFunc(handler404)

	// Parse all templates
	var err error
	t, err = template.ParseGlob("./templates/*")
	if err != nil {
		log.Println("Cannot parse templates:", err)
		os.Exit(-1)
	}

	// Set up the HTTP-server
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func handlerStats(w http.ResponseWriter, r *http.Request) {
	data := config.GetAll()
	t.ExecuteTemplate(w, "stats.html", data)
}

func handlerContact(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "contact.html", nil)
}

func handlerFAQ(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "faq.html", nil)
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
