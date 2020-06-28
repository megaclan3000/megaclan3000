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
	SteamID string `db:"steamid"`

	// The time when the entry was saved
	Time time.Time `db:"time"`

	// Total kill/death ratio
	TotalKD string `db:"total_kd"`

	// Total avarage damage per round
	TotalADR string `db:"total_adr"`

	// Last match avarage damage per round
	LastMatchADR string `db:"last_match_adr"`

	// Total kills
	TotalKills string `db:"total_kills"`

	// Total kills with headshot
	TotalKillsHeadshot string `db:"total_kills_headshot"`

	// Total shots hit
	TotalShotsHit string `db:"total_shots_hit"`

	// Total shots fired
	TotalShotsFired string `db:"total_shots_fired"`

	// Contribution score in last match
	LastMatchContributionScore string `db:"last_match_contribution_score"`

	// Damage dealt in last match
	LastMatchDamage string `db:"last_match_damage"`

	// Death count in last match
	LastMatchDeaths string `db:"last_match_deaths"`

	// Kills in last match
	LastMatchKills string `db:"last_match_kills"`

	// Number of round of last match
	LastMatchRounds string `db:"last_match_rounds"`

	// Las match kill/death ratio
	LastMatchKD string `db:"last_match_kd"`

	// Total hit ratio
	HitRatio string `db:"hit_ratio"`

	// Platime in the last 2 weeks
	Playtime2Weeks string `db:"playtime_2_weeks"`
}
