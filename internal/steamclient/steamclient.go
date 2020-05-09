package steamclient

import (
	log "github.com/sirupsen/logrus"
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
		log.Debugf("Fetching data for ID: %v", v)
		if pi, err := sc.getPlayerInfo(v); err == nil {
			log.Println("adding player", v)
			players = append(players, pi)
		} else {
			log.Println("skipping player", v)
			log.Warningf("Failed to get data for ID: %v", v)
		}
	}
	return players
}
