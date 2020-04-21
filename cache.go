package main

import (
	"log"
	"time"
)

type SteamPlayerInfo struct {
	players    []PlayerInfo
	lastUpdate time.Time
	ids        []string
}

func NewSteamPlayerInfo() *SteamPlayerInfo {
	players := []PlayerInfo{}

	ids := []string{
		"76561197978015984", //kapo
		"76561198092006615", //pablo
		"76561198104947907", //felix
		"76561197962156894", //alex
		"76561197967611281", //manu
		"76561198217140904", //bene
		"76561198962966497", //silvarse
		"76561198881047143", //lukas
		"76561197978562286", //enrico

	}

	// "76561198242556348", //sonarse

	info := SteamPlayerInfo{
		players:    players,
		ids:        ids,
		lastUpdate: time.Now(),
	}

	info.Refresh()

	return &info
}

func (spi *SteamPlayerInfo) Refresh() {
	log.Println("Cache outdated, refreshing...")
	// log.Println("time.Since(spi.lastUpdate))

	players := []PlayerInfo{}

	for _, v := range spi.ids {
		players = append(players, getPlayerInfo(v))
	}
	spi.players = players
}

func (spi *SteamPlayerInfo) GetAll() []PlayerInfo {
	if time.Since(spi.lastUpdate) > 6*time.Minute {
		spi.Refresh()
	}
	return spi.players
}
