package main

import (
	"flag"
	"sort"
	"strconv"
	"text/template"

	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/megaclan3000/megaclan3000/internal/steamclient"
)

var t *template.Template
var datastorage *DataStorage
var steamClient *steamclient.SteamClient

func main() {

	// -verbose flag to set logging level to DebugLevel
	flagVerbose := flag.Bool("verbose", false, "Enable verbose output")
	flag.Parse()

	if *flagVerbose {
		log.SetLevel(log.DebugLevel)
	}

	// Output to stdout instead of the default stderr
	// log.SetOutput(os.Stdout)

	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)

	// Read config and pull initial data
	steamClient = steamclient.NewSteamClient("./config.json")

	log.Info("Creating datastorage and getting initial values")
	datastorage = &DataStorage{Players: steamClient.GetPlayers()}

	r := mux.NewRouter()

	// Serve all static files in public directory
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	// Define routes
	r.HandleFunc("/", parseTemplates(handlerIndex))
	r.HandleFunc("/stats", parseTemplates(handlerStats))
	r.HandleFunc("/contact", parseTemplates(handlerContact))
	r.HandleFunc("/faq", parseTemplates(handlerFAQ))
	r.HandleFunc("/player/{id}", parseTemplates(handlerDetails))
	r.HandleFunc("/imprint", parseTemplates(handlerImprint))

	// Set custom 404 page
	r.NotFoundHandler = http.HandlerFunc(handler404)

	// Set up the HTTP-server
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	//start updating data every 5 minutes asynchroniusly
	go updateData()
	log.Fatal(srv.ListenAndServe())
}

func parseTemplates(h http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		// Parse all templates
		t, err = template.ParseGlob("./templates/*")
		if err != nil {
			log.Panic("Cannot parse templates", err)
		}

		h(w, r)
	}
}

func updateData() {

	// Get PlayerInfo for all players periodically and store/cache in
	// memory so we don't have to wait when retrieving it in the fronend

	for {
		log.Debug("Updating player information")
		datastorage.Players = steamClient.GetPlayers()

		// Sleep for a predefined duration (in minutes)
		time.Sleep(time.Duration(steamClient.Config.UpdateInterval) * time.Minute)
	}
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	if err := t.ExecuteTemplate(w, "index.html", nil); err != nil {
		log.Warn(err)
	}
}

func handlerStats(w http.ResponseWriter, r *http.Request) {

	// Stort players by Personastate (online status)
	sort.Slice(datastorage.Players, func(i, j int) bool {
		return datastorage.Players[i].PlayerSummary.Personastate > datastorage.Players[j].PlayerSummary.Personastate
	})

	if err := t.ExecuteTemplate(w, "stats.html", datastorage.Players); err != nil {
		log.Warn(err)
	}
}

func handlerContact(w http.ResponseWriter, r *http.Request) {
	if err := t.ExecuteTemplate(w, "contact.html", nil); err != nil {
		log.Warn(err)
	}
}

func handlerFAQ(w http.ResponseWriter, r *http.Request) {
	if err := t.ExecuteTemplate(w, "faq.html", nil); err != nil {
		log.Warn(err)
	}
}

func handler404(w http.ResponseWriter, r *http.Request) {
	if err := t.ExecuteTemplate(w, "404.html", nil); err != nil {
		log.Warn(err)
	}
}

func handlerDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var id uint64
	var err error

	if id, err = strconv.ParseUint(vars["id"], 10, 64); err == nil {
		if p, err := datastorage.GetPlayerInfoBySteamID(id); err == nil {
			if err := t.ExecuteTemplate(w, "details.html", p); err != nil {
				log.Warn(err)
			}
			return
		}
	}

	if err := t.ExecuteTemplate(w, "404.html", nil); err != nil {
		log.Warn(err)
	}
}

func handlerImprint(w http.ResponseWriter, r *http.Request) {
	if err := t.ExecuteTemplate(w, "imprint.html", nil); err != nil {
		log.Warn(err)
	}
}
