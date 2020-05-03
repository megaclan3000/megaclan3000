package steamclient

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// SteamConfig struct holds the values as read from the configuration file
// config.json on startup
type SteamConfig struct {
	SteamAPIKey string   `json:"SteamAPIKey"`
	SteamIDs    []string `json:"SteamIDs"`
	lastUpdate  time.Time
}

func readConfig() SteamConfig {

	conf := SteamConfig{}

	jsonFile, err := os.Open("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &conf)

	return conf
}
