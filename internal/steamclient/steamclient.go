package steamclient

import (
	"log"
)

type SteamClient struct {
	config SteamConfig
}

func NewSteamClient() *SteamClient {

	var config SteamConfig

	config = readConfig()
	return &SteamClient{config}
}

func (sc SteamClient) GetPlayers() []PlayerInfo {

	players := []PlayerInfo{}

	for _, v := range sc.config.SteamIDs {
		log.Println("Fetching data for ID:", v)
		players = append(players, sc.getPlayerInfo(v))
	}
	return players
}
