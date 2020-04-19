package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var steamAPIKey string

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view.html")
	data := GetAllPlayers()
	t.Execute(w, data)
}

func main() {

	content, err := ioutil.ReadFile("steamkey.txt")
	if err != nil {
		log.Fatal(err)
	}

	steamAPIKey = strings.TrimSuffix(string(content), "\n")

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
