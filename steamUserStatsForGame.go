package main

// https://developer.valvesoftware.com/wiki/Steam_Web_API#GetUserStatsForGame_.28v0002.29
// Returns a list of achievements for this user by app id

type UserStatsForGameData struct {
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
