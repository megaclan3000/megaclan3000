package main

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetAllPlayers() []PlayerInfo {

	//TODO add IDs
	IDs := []string{
		"76561198092006615",
		"76561198092006615",
	}
	var players []PlayerInfo

	for _, v := range IDs {
		players = append(players, getPlayerInfo(v))
	}

	return players
}

// PlayerInfo contains the information to be shown of a given player
type PlayerInfo struct {
	ID    string
	Name  string
	Stats map[string]string
}

func getPlayerInfo(steamID string) PlayerInfo {

	info := PlayerInfo{}
	info.ID = steamID
	info.Name = getPlayerName(steamID)
	info.Stats = getPlayerStats(steamID)

	return info
}

func getPlayerName(steamID string) string {
	return "randolf"
}

func getPlayerStats(steamID string) map[string]string {

	url := "http://api.steampowered.com/ISteamUserStats/GetUserStatsForGame/v0002/?appid=730&key=" + steamAPIKey + "&steamid=" + steamID

	client := http.Client{Timeout: time.Second * 2}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	people1 := SteamData{}
	jsonErr := json.Unmarshal(body, &people1)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	stats := make(map[string]string)

	for _, v := range people1.Playerstats.Stats {
		stats[v.Name] = strconv.Itoa(v.Value)
	}

	return stats

}

//SteamData holds the data returned by the steam web API for a player's stats
type SteamData struct {
	Playerstats struct {
		SteamID  string `json:"steamID"`
		GameName string `json:"gameName"`
		Stats    []struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"stats"`
		Achievements []struct {
			Name     string `json:"name"`
			Achieved int    `json:"achieved"`
		} `json:"achievements"`
	} `json:"playerstats"`
}
