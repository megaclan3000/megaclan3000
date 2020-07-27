package database

import (
	common "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"

	"time"
)

type RoundInfo struct {

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

	Kills []Kill
}

type Kill struct {
}

type MatchInfo struct {
	ID     uint64
	Time   time.Time
	Map    string
	Rounds []RoundInfo
}
