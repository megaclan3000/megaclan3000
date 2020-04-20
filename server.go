package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var steamAPIKey string
var playerInfo *SteamPlayerInfo

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view.html")
	data := playerInfo.GetAll()
	t.Execute(w, data)
}

func main() {

	content, err := ioutil.ReadFile("steamkey.txt")
	if err != nil {
		log.Fatal(err)
	}

	steamAPIKey = strings.TrimSuffix(string(content), "\n")

	playerInfo = NewSteamPlayerInfo()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
