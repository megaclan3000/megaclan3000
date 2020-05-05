package steamclient

import (
	"log"
)

// SteamClient acts as main interface to interact with the steam API and gather
// data
type SteamClient struct {
	config SteamConfig
}

// NewSteamClient returrns a new SteamClient
func NewSteamClient(configPath string) *SteamClient {

	// var config SteamConfig{
	// 	configPath: configconfigPath,
	// }
	config, err := NewSteamConfig(configPath)
	if err != nil {
		panic(err)
	}

	return &SteamClient{config}
}

// GetPlayers returns info for all players fetched from the API using the
// steamIDs in it's config
func (sc SteamClient) GetPlayers() []PlayerInfo {

	players := []PlayerInfo{}

	for _, v := range sc.config.SteamIDs {
		log.Println("Fetching data for ID:", v)
		players = append(players, sc.getPlayerInfo(v))
	}
	return players
}
