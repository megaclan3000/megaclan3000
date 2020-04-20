package main

import (
	"encoding/json"
	// "fmt"
	// "io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetAllPlayers() []PlayerInfo {

	//TODO add IDs
	IDs := []string{
		"76561197978015984", //kapo
		"76561198092006615", //pablo
		"76561198104947907", //felix
		"76561197962156894", //alex
		"76561197967611281", //manu
		"76561198217140904", //bene
		"76561198242556348", //sonarse
		"76561198881047143", //lukas
		"76561197978562286", //enrico

	}
	var players []PlayerInfo

	for _, v := range IDs {
		players = append(players, getPlayerInfo(v))
	}

	return players
}

// PlayerInfo contains the information to be shown of a given player
type PlayerInfo struct {
	ID                       string
	Steamid                  string
	Name                     string
	Communityvisibilitystate string
	Profilestate             string
	Personaname              string
	Profileurl               string
	Avatar                   string
	Avatarmedium             string
	Avatarfull               string
	Personastate             string
	Realname                 string
	Primaryclanid            string
	Timecreated              string
	Personastateflags        string
	Loccountrycode           string
	Locstatecode             string
	Loccityid                string
	Stats                    map[string]string
}

func getPlayerInfo(steamID string) PlayerInfo {

	url := "https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v2/?key=" + steamAPIKey + "&steamids=" + steamID

	people1 := AutoGenerated{}
	getJson(url, &people1)

	info := PlayerInfo{}
	info.Steamid = people1.Response.Players[0].Steamid
	info.Personaname = people1.Response.Players[0].Personaname
	info.Profileurl = people1.Response.Players[0].Profileurl
	info.Avatar = people1.Response.Players[0].Avatar
	info.Avatarmedium = people1.Response.Players[0].Avatarmedium
	info.Avatarfull = people1.Response.Players[0].Avatarfull
	info.Personastate = strconv.Itoa(people1.Response.Players[0].Personastate)
	info.Communityvisibilitystate = strconv.Itoa(people1.Response.Players[0].Communityvisibilitystate)
	info.Profilestate = strconv.Itoa(people1.Response.Players[0].Profilestate)
	info.Loccityid = strconv.Itoa(people1.Response.Players[0].Loccityid)
	info.Timecreated = strconv.Itoa(people1.Response.Players[0].Timecreated)
	info.Personastateflags = strconv.Itoa(people1.Response.Players[0].Personastateflags)
	info.Realname = people1.Response.Players[0].Realname
	info.Primaryclanid = people1.Response.Players[0].Primaryclanid
	info.Loccountrycode = people1.Response.Players[0].Loccountrycode
	info.Locstatecode = people1.Response.Players[0].Locstatecode

	// info.Name = getPlayerName(steamID)
	info.Stats = getPlayerStats(steamID)

	return info
}

func getPlayerStats(steamID string) map[string]string {

	// url := "http://api.steampowered.com/ISteamUserStats/GetUserStatsForGame/v0002/?appid=730&key=" + steamAPIKey + "&steamid=" + steamID
	url := "https://api.steampowered.com/ISteamUserStats/GetUserStatsForGame/v2/?appid=730&key=" + steamAPIKey + "&steamid=" + steamID
	// url := "https://api.steampowered.com/ISteamUserStats/GetUserStatsForGame/v2/

	people1 := SteamData{}
	getJson(url, &people1)
	stats := make(map[string]string)

	for _, v := range people1.Playerstats.Stats {
		// log.Println(v.Name, ":\t\t", v.Value)
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

type AutoGenerated struct {
	Response struct {
		Players []struct {
			Steamid                  string `json:"steamid"`
			Communityvisibilitystate int    `json:"communityvisibilitystate"`
			Profilestate             int    `json:"profilestate"`
			Personaname              string `json:"personaname"`
			Profileurl               string `json:"profileurl"`
			Avatar                   string `json:"avatar"`
			Avatarmedium             string `json:"avatarmedium"`
			Avatarfull               string `json:"avatarfull"`
			Personastate             int    `json:"personastate"`
			Realname                 string `json:"realname"`
			Primaryclanid            string `json:"primaryclanid"`
			Timecreated              int    `json:"timecreated"`
			Personastateflags        int    `json:"personastateflags"`
			Loccountrycode           string `json:"loccountrycode"`
			Locstatecode             string `json:"locstatecode"`
			Loccityid                int    `json:"loccityid"`
		} `json:"players"`
	} `json:"response"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		log.Fatal(err)
		// return err
	}

	log.Println("Getting: ", url, " --------------")
	// log.Println(url)
	// log.Println(r.Body)
	// log.Println("--------------")
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
