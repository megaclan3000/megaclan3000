package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/pinpox/megaclan3000/internal/database"
	"github.com/pinpox/megaclan3000/internal/steamclient"
)

var t *template.Template
var datastorage *database.DataStorage
var steamClient *steamclient.SteamClient

func main() {
	log.Println("main")

	var err error
	// Read config and pull initial data
	steamClient = steamclient.NewSteamClient()

	log.Println("Creating datastorage")
	if datastorage, err = database.NewDataStorage("./data.db"); err != nil {
		log.Fatal("Failed to open database", err)
	}

	r := mux.NewRouter()

	// Serve all static files in public directory
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	// Define routes
	r.HandleFunc("/", handlerIndex)
	r.HandleFunc("/stats", handlerStats)
	r.HandleFunc("/contact", handlerContact)
	r.HandleFunc("/faq", handlerFAQ)
	r.HandleFunc("/player/{id}", handlerDetails)
	r.HandleFunc("/imprint", handlerImprint)

	// Set custom 404 page
	r.NotFoundHandler = http.HandlerFunc(handler404)

	// Parse all templates
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

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "index.html", nil)
}

func handlerStats(w http.ResponseWriter, r *http.Request) {
	//TODO do something with this error
	data, _ := datastorage.GetAllPlayers()
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

	if p, err := datastorage.GetPlayerInfoBySteamID(vars["id"]); err == nil {
		t.ExecuteTemplate(w, "details.html", p)
		return
	}
	t.ExecuteTemplate(w, "404.html", nil)
}

func handlerImprint(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "imprint.html", nil)
}
