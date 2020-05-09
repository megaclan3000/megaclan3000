package steamclient

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

	if info.PlayerSummary, err = sc.GetPlayerSummary(steamID); err != nil {
		return info, err
	}

	if info.UserStatsForGame, err = sc.GetUserStatsForGame(steamID); err != nil {
		return info, err
	}

	if info.RecentlyPlayedGames, err = sc.GetRecentlyPlayedGames(steamID); err != nil {
		return info, err
	}

	return info, nil
}
