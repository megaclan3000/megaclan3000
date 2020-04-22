package main

import (
	"fmt"
	"strconv"
)

// https://developer.valvesoftware.com/wiki/Steam_Web_API#GetUserStatsForGame_.28v0002.29
// Returns a list of achievements for this user by app id

type userStatsForGameData struct {
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

type UserStatsForGame struct {
	SteamID      string
	GameName     string
	Stats        map[string]string
	Archivements map[string]string
}

func getUserStatsForGame(steamID string) UserStatsForGame {

	url := "https://api.steampowered.com/ISteamUserStats/GetUserStatsForGame/v2/?appid=730&key=" + config.SteamAPIKey + "&steamid=" + steamID
	data := userStatsForGameData{}
	getJson(url, &data)

	stats := make(map[string]string)
	archivements := make(map[string]string)

	for _, v := range data.Playerstats.Stats {
		stats[v.Name] = strconv.Itoa(v.Value)
	}

	if total_deaths, err := strconv.ParseFloat(stats["total_deaths"], 64); err == nil {
		if total_kills, err := strconv.ParseFloat(stats["total_kills"], 64); err == nil {
			stats["total_kd"] = fmt.Sprintf("%f", total_kills/total_deaths)
		}
	}

	if last_deaths, err := strconv.ParseFloat(stats["last_match_deaths"], 64); err == nil {
		if last_kills, err := strconv.ParseFloat(stats["last_match_kills"], 64); err == nil {
			stats["last_match_kd"] = fmt.Sprintf("%f", last_kills/last_deaths)
		}
	}

	if total_shots_fired, err := strconv.ParseFloat(stats["total_shots_fired"], 64); err == nil {
		if total_shots_hit, err := strconv.ParseFloat(stats["total_shots_hit"], 64); err == nil {
			stats["hit_ratio"] = fmt.Sprintf("%f", total_shots_hit/total_shots_fired)
		}
	}

	if secI, err := strconv.Atoi(stats["total_time_played"]); err == nil {
		stats["played_hours"] = strconv.Itoa(secI / 3600)
	}

	for _, v := range data.Playerstats.Achievements {
		archivements[v.Name] = strconv.Itoa(v.Achieved)
	}

	return UserStatsForGame{
		SteamID:      data.Playerstats.SteamID,
		GameName:     data.Playerstats.GameName,
		Stats:        stats,
		Archivements: archivements,
	}
}
