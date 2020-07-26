package demoparser

import (
	"time"

	common "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

// Match represents a complete match (both sides) as parsed from a demo file.
// It mainly consists of a slice of rounds that happened during that match.

type Rounds map[int]*Round

type Match struct {
	// The ID of the match
	ID string

	// The time and date when the match took place
	Time time.Time

	// Map of the match
	Map string

	// The rounds of the match. This is the main source of information we can
	// use later on.
	Rounds Rounds

	// Kills during warmup are saved here, as there is no round active
	WarmupKills []events.Kill
}

// Round holds the result of parsing all ticks of a match that correspond to a
// single round.
type Round struct {

	// Tick at which the round started
	TimeStart time.Duration

	// Tick at which the round ended
	TimeEnd time.Duration

	// Team that won the round, either "T" or "CT"
	TeamWon common.Team

	// If the bomb was planted during the round
	BombPlanted bool

	// If the bomb was defused during the round
	BombDefused bool

	// If the bomb exploded during the round
	BombExploded bool

	// List of players that where presend in the round
	Players []common.Player

	Kills []events.Kill

	PlayerHurt []*events.PlayerHurt
}
