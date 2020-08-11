package steamclient

import (
	"strconv"

	log "github.com/sirupsen/logrus"
)

// SteamClient acts as main interface to interact with the steam API and gather
// data
type SteamClient struct {
	Config SteamConfig
}

// NewSteamClient returrns a new SteamClient
func NewSteamClient(configPath string) *SteamClient {

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

	for _, v := range sc.Config.SteamIDs {
		log.Debugf("Fetching data for ID: %v", v)
		if pi, err := sc.getPlayerInfo(v); err == nil {
			players = append(players, pi)
		} else {
			log.Warningf("Failed to get data for ID: %v", v)
			log.Warn(err)
		}
	}
	return players
}

func (sc SteamClient) GetAvatarUrl(id uint64) string {

	//PlayerSummary
	summaryData := playerSummariesData{}
	url := "https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v2/?key=" + sc.Config.SteamAPIKey + "&steamids=" + strconv.FormatUint(id, 10)

	if err := getJSON(url, &summaryData); err != nil {
		log.Warn(err)

		return "/public/img/avatars/other.jpg"
	}

	if summary, err := sc.parsePlayerSummary(summaryData); err != nil {
		log.Warn(err)
		return "/public/img/avatars/other.jpg"
	} else {
		return summary.Avatarfull
	}

}
