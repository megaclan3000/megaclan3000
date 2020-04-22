package main

import (
	"html/template"
	"log"
	"net/http"
	// "github.com/davecgh/go-spew/spew"
)

var config SteamConfig

func main() {

	config = readConfig()
	config.Refresh()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/stats.html")
	data := config.GetAll()
	t.Execute(w, data)
}
