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
	players     []PlayerInfo
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

// Refresh will get new data from the steam API and update the in-memory
// players information
func (conf *SteamConfig) Refresh() {
	log.Println("Cache outdated, refreshing...")

	players := []PlayerInfo{}

	for _, v := range conf.SteamIDs {
		log.Println("")
		log.Println("Fetching data for ID:", v)

		//TODO FIX THIS
		// players = append(players, getPlayerInfo(v))
	}
	conf.lastUpdate = time.Now()
	conf.players = players
}
