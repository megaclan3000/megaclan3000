package steamclient

import "log"

// PlayerInfo contains the information to be shown of a given player
type PlayerInfo struct {
	PlayerSummary       PlayerSummary
	UserStatsForGame    UserStatsForGame
	RecentlyPlayedGames RecentlyPlayedGames
	PlayerHistory       PlayerHistory
}

func (sc *SteamClient) getPlayerInfo(steamID string) PlayerInfo {

	info := PlayerInfo{}

	info.PlayerSummary = sc.GetPlayerSummary(steamID)
	info.UserStatsForGame = sc.GetUserStatsForGame(steamID)
	log.Println("calculated", info.UserStatsForGame.Extra)
	info.RecentlyPlayedGames = sc.GetRecentlyPlayedGames(steamID)

	return info
}
