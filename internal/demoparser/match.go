package demoparser

import (
	"time"

	// "github.com/golang/geo/r3"
	common "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
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
	TeamWon       bool
	MapName       string
	MapIconURL    string
	UploadTime    time.Time
	MatchDuration time.Duration
	DemoLinkURL   string
	PlayerInfos   map[uint64]*ScoreboardTeamMemberInfo
}

type ScoreboardTeamMemberInfo struct {
	AvatarURL   string
	Name        string
	RankIconURL string
	ClanTag     string
}

type RoundKill struct {
	VictimSteamID64 uint64
	KillerSteamID64 uint64
	KillerWeapon    string
}

type ScoreboardRound struct {
	TeamWon        common.Team
	WinReason      events.RoundEndReason
	Duration       time.Duration
	TeamAKills     []RoundKill
	TeamASurvivors int
	TeamBSurvivors int
	TeamBKills     []RoundKill
}
type ScoreboardLine struct {
	PlayerSteamID64  uint64
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

func GetMatchInfo(id int) (InfoStruct, error) {
	//TODO
	p := NewMyParser()
	var info InfoStruct
	//TODO get correct path for demo file
	err := p.Parse("internal/demoparser/testdata/demo1.dem", &info)
	return info, err
}
