package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	log.Println("Downloading:", url)
	r, err := myClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	if r.StatusCode != 200 {
		log.Println("Failed to get data from:")
		log.Println("'" + url + "'")
		log.Println(r.StatusCode, http.StatusText(r.StatusCode))
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func divideStringFloats(a, b string) string {

	if aF, err := strconv.ParseFloat(a, 64); err == nil {
		if bF, err := strconv.ParseFloat(b, 64); err == nil {
			return fmt.Sprintf("%f", aF/bF)
		}
	}
	return ""
}
