package steamclient

import (
	"time"
)

// PlayerHistory holds the players history data from the player_history table.
// Stats values that need to be saved over time, are added to this table and
// object.
type PlayerHistory struct {

	//SteamID of the player
	SteamID string

	// Data array containing entries to the history with time and values
	Data []PlayerHistoryEntry
}

// PlayerHistoryEntry holds on snapshot of the values that are tracked over
// time. The steamid field is no primary key in this case, since every player
// will have multiple entries
type PlayerHistoryEntry struct {

	// SteamID of the player
	SteamID string

	// The time when the entry was saved
	Time time.Time

	// Total kill/death ratio
	TotalKD string

	// Total avarage damage per round
	TotalADR string

	// Last match avarage damage per round
	LastMatchADR string

	// Total kills
	TotalKills string

	// Total kills with headshot
	TotalKillsHeadshot string

	// Total shots hit
	TotalShotsHit string

	// Total shots fired
	TotalShotsFired string

	// Contribution score in last match
	LastMatchContributionScore string

	// Damage dealt in last match
	LastMatchDamage string

	// Death count in last match
	LastMatchDeaths string

	// Kills in last match
	LastMatchKills string

	// Number of round of last match
	LastMatchRounds string

	// Las match kill/death ratio
	LastMatchKD string

	// Total hit ratio
	HitRatio string

	// Platime in the last 2 weeks
	Playtime2Weeks string
}
