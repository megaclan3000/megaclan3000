package demoparser

import "time"

// Match represents a complete match (both sides) as parsed from a demo file.
// It mainly consists of a slice of rounds that happened during that match.
type Match struct {
	// The ID of the match
	ID string

	// The time and date when the match took place
	Time time.Time

	// Map of the match
	Map string

	// The rounds of the match. This is the main source of information we can
	// use later on.
	Rounds []Round
}

// Round holds the result of parsing all ticks of a match that correspond to a
// single round.
type Round struct {

	// Round number
	Number int

	// Tick at which the round started
	TickStart int

	// Tick at which the round ended
	TickEnd int

	// Team that won the round, either "T" or "CT"
	TeamWon string

	// If the bomb was planted during the round
	BombPlanted bool

	// If the bomb was defused during the round
	BombDefused bool

	// If the bomb exploded during the round
	BombExploded bool

	// List of players that where presend in the round
	Players []Player
}

type Player struct {
	// SteamID of the player
	ID string

	// The the player was on during the round, "T" or "CT"
	Team string

	// If the player won the round
	Won bool

	// If the player received a MVP star for the round
	Mvp bool

	// Reason for the received MVP star
	MvpReason string

	// If the player planted the bomb during the round
	BombPlanted bool

	// If the player defused the bomb during the round
	BombDefused bool

	// Number the player fired during the round
	ShotsFired int

	// Number of shots the player it during the round
	ShotsHit int

	// Number of headshots that the player hit during the round
	Headshots int

	// Number of kills and assists the player made during the round
	Kills []Kill

	// Damages he did during the round, summarized to victims. This slice will
	// have a maximum lenght of the numbers of players of both teams
	// (Competitive: 10) if he damaged all players in the round of both teams.
	Victims []DamageVictim
}

// Kill holds the information about a kill that occured during a round, as
// parsed from the demo file
type Kill struct {

	// Tick at which the kill was made
	Tick int

	// True if it was a assist, false if it was a kill
	Assist bool

	// SteamID of the victim
	VictimID string

	// SteamID of the player he was assisted by, might be nil
	AssistID string

	// If it was an assist, the steamID of the player he assisted. This is nil
	// if it was a kill
	AssistedID string

	// The weapon used by the player at the time of the kill. He might have
	// used other weapons before that contributed damage, the last weapon used
	// is recorded.
	WeaponUsedID int
}

type DamageVictim struct {

	// SteamID of the player he damaged, might be on either team
	DamagedID string

	// The weapon used to inflict the most amount of damage to the this victim.
	MostDamageWeaponID int

	// Amount of damage inflicted to this victim in total
	Amount int
}
