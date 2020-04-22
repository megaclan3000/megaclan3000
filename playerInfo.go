package main

// PlayerInfo contains the information to be shown of a given player
type PlayerInfo struct {
	PlayerSummary       PlayerSummary
	UserStatsForGame    UserStatsForGame
	RecentlyPlayedGames RecentlyPlayedGames
}

func getPlayerInfo(steamID string) PlayerInfo {

	info := PlayerInfo{}

	info.PlayerSummary = getPlayerSummary(steamID)
	info.UserStatsForGame = getUserStatsForGame(steamID)
	info.RecentlyPlayedGames = getRecentlyPlayedGames(steamID)

	// spew.Dump(info)

	return info
}
