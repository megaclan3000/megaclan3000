package steamclient

import (
	"errors"

	log "github.com/sirupsen/logrus"
)

// PlayerInfo contains the information to be shown of a given player
type PlayerInfo struct {
	PlayerSummary       PlayerSummary
	UserStatsForGame    UserStatsForGame
	RecentlyPlayedGames RecentlyPlayedGames
	PlayerHistory       PlayerHistory
}

func (sc *SteamClient) getPlayerInfo(steamID string) (PlayerInfo, error) {

	if len(steamID) <= 1 {
		panic("Tried to get playerInfo for empty ID")
	}

	info := PlayerInfo{}
	var err error
	var url string

	//PlayerSummary
	summaryData := playerSummariesData{}
	url = "https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v2/?key=" + sc.Config.SteamAPIKey + "&steamids=" + steamID

	if err := getJSON(url, &summaryData); err != nil {
		log.Warn(err)
		return info, errors.New("Unable to get PlayerSummary for: " + steamID)
	}

	if info.PlayerSummary, err = sc.parsePlayerSummary(summaryData); err != nil {
		return info, err
	}

	//UserStatsForGame
	statsData := userStatsForGameData{}
	url =
		"https://api.steampowered.com/ISteamUserStats/GetUserStatsForGame/v2/?appid=730&key=" +
			sc.Config.SteamAPIKey + "&steamid=" + steamID

	if err := getJSON(url, &statsData); err != nil {
		log.Warn(err)
		return info, errors.New("Unable to get UserStatsForGame for: " + steamID)
	}

	if info.UserStatsForGame, err = sc.parseUserStatsForGame(statsData); err != nil {
		return info, err
	}

	// RecentlyPlayedGames
	recentData := recentlyPlayedGamesData{}
	url = "https://api.steampowered.com/IPlayerService/GetRecentlyPlayedGames/v0001/?key=" + sc.Config.SteamAPIKey + "&steamid=" + steamID

	if err := getJSON(url, &recentData); err != nil {
		log.Warn(err)
		return info, errors.New("Unable to get RecentlyPlayedGames for: " + steamID)
	}

	if info.RecentlyPlayedGames, err = sc.parseRecentlyPlayedGames(recentData, steamID); err != nil {
		info.RecentlyPlayedGames.SteamID = steamID
		log.Warnf("Unable to parse RecentlyPlayedGames for: %v Might have not played in the last two weeks.", steamID)
	}

	return info, nil
}
