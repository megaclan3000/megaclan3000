package main

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

	for _, v := range config.SteamIDs {
		log.Println("")
		log.Println("Fetching data for ID:", v)

		players = append(players, getPlayerInfo(v))
	}
	conf.lastUpdate = time.Now()
	conf.players = players
}

// GetAll returns all the player objects. After a specified time these will be
// fetched from the steam API, otherwise returned as already in memory
func (conf *SteamConfig) GetAll() []PlayerInfo {
	if time.Since(conf.lastUpdate) > 6*time.Minute {
		conf.Refresh()
	}
	return conf.players
}
