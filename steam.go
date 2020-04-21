package main

import (
	"encoding/json"
	"fmt"

	"log"
	"net/http"
	"strconv"
	"time"
)

// PlayerInfo contains the information to be shown of a given player
type PlayerInfo struct {
	ID                       string
	Steamid                  string
	Communityvisibilitystate string
	Profilestate             string
	Personaname              string
	Profileurl               string
	Avatar                   string
	Avatarmedium             string
	Realname                 string
	Avatarfull               string
	Lastlogoff               string
	Personastate             string
	Primaryclanid            string
	Timecreated              string
	Personastateflags        string
	Gameextrainfo            string
	Gameid                   string
	Loccountrycode           string

	Stats    map[string]string
	Playtime map[string]string

	TotalKD     string
	PlayedHours string
	LastMatchKD string
	HitRatio    string
}

func getPlayerInfo(steamID string) PlayerInfo {

	url := "https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v2/?key=" + steamAPIKey + "&steamids=" + steamID

	people1 := PlayerSummariesData{}
	getJson(url, &people1)

	info := PlayerInfo{}

	info.Realname = people1.Response.Players[0].Realname
	info.Steamid = people1.Response.Players[0].Steamid
	info.Gameid = people1.Response.Players[0].Gameid
	info.Loccountrycode = people1.Response.Players[0].Loccountrycode
	info.Profileurl = people1.Response.Players[0].Profileurl
	info.Avatar = people1.Response.Players[0].Avatar
	info.Avatarmedium = people1.Response.Players[0].Avatarmedium
	info.Primaryclanid = people1.Response.Players[0].Primaryclanid
	info.Avatarfull = people1.Response.Players[0].Avatarfull
	info.Gameextrainfo = people1.Response.Players[0].Gameextrainfo
	info.Personaname = people1.Response.Players[0].Personaname
	info.Lastlogoff = strconv.Itoa(people1.Response.Players[0].Lastlogoff)
	info.Personastate = strconv.Itoa(people1.Response.Players[0].Personastate)
	info.Timecreated = strconv.Itoa(people1.Response.Players[0].Timecreated)
	info.Communityvisibilitystate = strconv.Itoa(people1.Response.Players[0].Communityvisibilitystate)
	info.Profilestate = strconv.Itoa(people1.Response.Players[0].Profilestate)

	info.Stats = getPlayerStats(steamID)
	info.Playtime = getPlaytime(steamID)

	if total_deaths, err := strconv.ParseFloat(info.Stats["total_deaths"], 64); err == nil {
		if total_kills, err := strconv.ParseFloat(info.Stats["total_kills"], 64); err == nil {
			info.TotalKD = fmt.Sprintf("%f", total_kills/total_deaths)
		}
	}

	if last_deaths, err := strconv.ParseFloat(info.Stats["last_match_deaths"], 64); err == nil {
		if last_kills, err := strconv.ParseFloat(info.Stats["last_match_kills"], 64); err == nil {
			info.LastMatchKD = fmt.Sprintf("%f", last_kills/last_deaths)
		}
	}

	if total_shots_fired, err := strconv.ParseFloat(info.Stats["total_shots_fired"], 64); err == nil {
		if total_shots_hit, err := strconv.ParseFloat(info.Stats["total_shots_hit"], 64); err == nil {
			info.HitRatio = fmt.Sprintf("%f", total_shots_hit/total_shots_fired)
		}
	}

	if secI, err := strconv.Atoi(info.Stats["total_time_played"]); err == nil {
		info.PlayedHours = strconv.Itoa(secI / 3600)
	}

	return info
}

func getPlaytime(steamID string) map[string]string {

	url := "https://api.steampowered.com/IPlayerService/GetRecentlyPlayedGames/v0001/?key=" + steamAPIKey + "&steamid=" + steamID

	people1 := RecentlyPlayedGamesData{}
	getJson(url, &people1)
	stats := make(map[string]string)

	for _, v := range people1.Response.Games {
		if v.Appid == 730 {
			stats["PlaytimeForever"] = strconv.Itoa(v.PlaytimeForever / 60)
			stats["Playtime2Weeks"] = strconv.Itoa(v.Playtime2Weeks / 60)
		}
	}
	return stats
}

func getPlayerStats(steamID string) map[string]string {

	url := "https://api.steampowered.com/ISteamUserStats/GetUserStatsForGame/v2/?appid=730&key=" + steamAPIKey + "&steamid=" + steamID

	people1 := UserStatsForGameData{}
	getJson(url, &people1)
	stats := make(map[string]string)

	for _, v := range people1.Playerstats.Stats {
		// log.Println(v.Name, ":\t\t", v.Value)
		stats[v.Name] = strconv.Itoa(v.Value)
	}

	return stats
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Getting: ", url, "\n--------------")
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
