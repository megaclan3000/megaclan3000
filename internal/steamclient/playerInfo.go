package steamclient

import "errors"

// PlayerInfo contains the information to be shown of a given player
type PlayerInfo struct {
	PlayerSummary       PlayerSummary
	UserStatsForGame    UserStatsForGame
	RecentlyPlayedGames RecentlyPlayedGames
	PlayerHistory       PlayerHistory
}

func (sc *SteamClient) getPlayerInfo(steamID string) (PlayerInfo, error) {

	info := PlayerInfo{}
	var err error
	var url string

	//PlayerSummary
	summaryData := playerSummariesData{}
	url = "https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v2/?key=" + sc.Config.SteamAPIKey + "&steamids=" + steamID

	if getJSON(url, &summaryData); err != nil {
		return info, errors.New("Unable to get PlayerSummary for: " + steamID)
	}

	if info.PlayerSummary, err = sc.ParsePlayerSummary(summaryData); err != nil {
		return info, err
	}

	//UserStatsForGame
	statsData := userStatsForGameData{}
	url =
		"https://api.steampowered.com/ISteamUserStats/GetUserStatsForGame/v2/?appid=730&key=" +
			sc.Config.SteamAPIKey + "&steamid=" + steamID

	if err := getJSON(url, &statsData); err != nil {
		return info, err
	}

	if info.UserStatsForGame, err = sc.ParseUserStatsForGame(statsData); err != nil {
		return info, err
	}

	// RecentlyPlayedGames
	recentData := recentlyPlayedGamesData{}
	url = "https://api.steampowered.com/IPlayerService/GetRecentlyPlayedGames/v0001/?key=" + sc.Config.SteamAPIKey + "&steamid=" + steamID

	if err := getJSON(url, &recentData); err != nil {
		return info, err
	}

	if info.RecentlyPlayedGames, err = sc.ParseRecentlyPlayedGames(recentData, steamID); err != nil {
		return info, err
	}

	return info, nil
}
