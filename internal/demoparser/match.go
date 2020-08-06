package demoparser

import (
	"time"

	// "github.com/golang/geo/r3"
	common "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
	// "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type InfoStruct struct {
	General    ScoreboardGeneral
	Scoreboard struct {
		TeamA []ScoreboardLine
		TeamB []ScoreboardLine
	}
	Rounds            []ScoreboardRound
	Weapons           []ScoreboardWeaponLine
	Duels             [][]int
	HeatmapsImageURLs []string
	Megacoins         []MegacoinPlayer
}

type ScoreboardGeneral struct {
	MapName     string
	MapIconURL  string
	UploadTime  time.Time
	DemoLinkURL string
}

type RoundKill struct {
	KillerName      string
	KillerAvatarURL string
	KillerWeapon    string
	VictimName      string
}

type ScoreboardRound struct {
	TeamWon        common.Team
	WinReason      common.WinReason
	Duration       time.Duration
	TeamAKills     []RoundKill
	TeamASurvivors int
	TeamBSurvivors int
	TeamBKills     []RoundKill
}
type ScoreboardLine struct {
	AvatarURL        string
	Name             string
	RankIconURL      string
	Kills            int
	Deaths           int
	Assists          int
	KDDiff           int
	KD               float64
	ADR              int
	HSPrecent        int
	FirstKills       int
	FirstDeaths      int
	TradeKills       int
	TradeDeaths      int
	TradeFirstKills  int
	TradeFirstDeaths int
	RoundsWonV5      int
	RoundsWonV4      int
	RoundsWonV3      int
	RoundsWonV2      int
	RoundsWonV1      int
	Rounds5k         int
	Rounds4k         int
	Rounds3k         int
	Rounds2k         int
	Rounds1k         int
	KAST             int
	HLTV             int
}

type WeaponUser struct {
	Kills     int
	HSPercent int
	Accuracy  int
	Damage    int
}

type ScoreboardWeaponLine struct {
	WeaponName    string
	WeaponIconURL string
	TeamA         []WeaponUser
	TeamB         []WeaponUser
}

type MegacoinPlayer struct {
	//TODO
	ForCriteriaA int
	ForCriteriaB int
	ForCriteriaC int
}

func GetMatchInfo(id int) InfoStruct {
	//TODO
	p := NewMyParser()
	var info InfoStruct
	p.Parse("", &info)
	return info
}
