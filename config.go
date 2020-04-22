package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

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

func (conf *SteamConfig) GetAll() []PlayerInfo {
	if time.Since(conf.lastUpdate) > 6*time.Minute {
		conf.Refresh()
	}
	return conf.players
}
