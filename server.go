package main

import (
	"flag"
	"text/template"

	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/pinpox/megaclan3000/internal/database"
	"github.com/pinpox/megaclan3000/internal/steamclient"
)

var t *template.Template
var datastorage *database.DataStorage
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

	var err error
	// Read config and pull initial data
	steamClient = steamclient.NewSteamClient("./config.json")

	log.Info("Creating datastorage")
	if datastorage, err = database.NewDataStorage("./data.db", "./schema.sql"); err != nil {
		log.Fatal("Failed to open database:", err)
	}

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
	var err error
	for {

		// Get PlayerInfo for all players
		players := steamClient.GetPlayers()

		// Save to db
		for _, v := range players {

			log.Infof("Updating data for %v (%v)", v.PlayerSummary.Personaname, v.PlayerSummary.SteamID)
			if err = datastorage.UpdatePlayerInfo(v); err != nil {
				log.Fatal(err)
			}

			// get latest timestamp
			var lastUpdateTime time.Time
			if lastUpdateTime, err = datastorage.GetPlayerHistoryLatestTime(v.PlayerSummary.SteamID); err != nil {
				log.Fatal(err)
			}

			// if part threshold, update
			log.Println("updatetime", time.Now().Sub(lastUpdateTime).Minutes())
			if time.Now().Sub(lastUpdateTime).Minutes() > float64(steamClient.Config.HistoryInterval) {
				log.Infof("Updating history for %v (%v)", v.PlayerSummary.Personaname, v.PlayerSummary.SteamID)
				err = datastorage.UpdatePlayerHistory(v)
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		// Sleep for a predefined duration (in minutes)
		time.Sleep(time.Duration(steamClient.Config.UpdateInterval) * time.Minute)
	}
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "index.html", nil)
}

func handlerStats(w http.ResponseWriter, r *http.Request) {

	var players []steamclient.PlayerInfo
	var err error

	if players, err = datastorage.GetAllPlayers(); err != nil {
		log.Error("Error getting stats from database:", err)
		t.ExecuteTemplate(w, "404.html", nil)
		return
	}
	t.ExecuteTemplate(w, "stats.html", players)
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
