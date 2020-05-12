package steamclient

// PlayerHistory holds the players history data from the player_history table.
// Stats values that need to be saved over time, are added to this table and
// object.
type PlayerHistory struct {
	SteamID string
	Data    []PlayerHistoryEntry
}

type PlayerHistoryEntry struct {
	Time    string
	TotalKD string
	//ADR string
	//TODO add otheres here, like ADR
}
