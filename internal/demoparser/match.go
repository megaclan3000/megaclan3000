package demoparser

import (
	"errors"
	"log"
	"strconv"
	"time"

	// "github.com/golang/geo/r3"
	common "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
	// "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type InfoStruct struct {
	General           ScoreboardGeneral
	Players           ScoreboardPlayers
	Rounds            []ScoreboardRound
	Weapons           []ScoreboardWeaponLine
	Duels             [][]int
	HeatmapsImageURLs []string
	Megacoins         []MegacoinPlayer
}

type ScoreboardGeneral struct {
	ClanWonMatch  bool
	ScoreClan     int
	ScoreEnemy    int
	MapName       string
	MapIconURL    string
	UploadTime    time.Time
	MatchDuration time.Duration
	DemoLinkURL   string
}

func (sg ScoreboardGeneral) TopWeaponsByKills() []common.EquipmentType {
	//TODO use real data
	return []common.EquipmentType{
		common.EqAK47,
		common.EqAWP,
		common.EqM4A1,
		common.EqBizon,
	}
}

type WeaponStat struct {
	Kills     int
	Headshots int
	Accuracy  int
	Damage    int
}

type RoundKill struct {
	KillerTeamString   string
	VictimTeamString   string
	AssisterTeamString string
	Victim             *ScoreboardPlayer
	Killer             *ScoreboardPlayer
	Assister           *ScoreboardPlayer
	KillerWeapon       common.EquipmentType
}

type ScoreboardRound struct {
	TeamWon            common.Team
	ClanWonRound       bool
	WinReason          events.RoundEndReason
	Duration           time.Duration
	TeamClanSurvivors  int
	TeamEnemySurvivors int
	ScoreClan          int
	ScoreEnemy         int
	TeamEnemyKills     []RoundKill
	TeamClanKills      []RoundKill
}

// type ScoreboardLine struct {
// 	PlayerSteamID64  uint64
// 	Kills            int
// 	Deaths           int
// 	Assists          int
// 	KDDiff           int
// 	KD               float64
// 	ADR              int
// 	HSPrecent        int
// 	FirstKills       int
// 	FirstDeaths      int
// 	TradeKills       int
// 	TradeDeaths      int
// 	TradeFirstKills  int
// 	TradeFirstDeaths int
// 	RoundsWonV5      int
// 	RoundsWonV4      int
// 	RoundsWonV3      int
// 	Rounds5k         int
// 	Rounds4k         int
// 	Rounds3k         int
// }

type WeaponUser struct {
	Kills     int
	HSPercent int
	Accuracy  int
	Damage    int
}

type ScoreboardWeaponLine struct {
	WeaponName    string
	WeaponIconURL string
	TeamClan      []WeaponUser
	TeamEnemy     []WeaponUser
}

type MegacoinPlayer struct {
	//TODO
	ForCriteriaA int
	ForCriteriaB int
	ForCriteriaC int
}

func GetMatchInfo(id string) (InfoStruct, error) {
	//TODO
	p := NewMyParser()
	var info InfoStruct
	//TODO get correct path for demo file
	err := p.Parse("internal/demoparser/testdata/demo"+id+".dem", &info)
	return info, err
}

// API methods
// Scoreboard
func (is InfoStruct) GetScoreboard() ScoreboardPlayers {
	//TODO implement
	return is.Players
	// out := Scoreboard{
	// 	Clan: []ScoreboardPlayer{
	// 		{
	// 			Name:             "randolf",
	// 			Clantag:          "megaclan3000",
	// 			AvatarURL:        "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
	// 			Rank:             7,
	// 			Steamid64:        1,
	// 			Kills:            2,
	// 			Deaths:           3,
	// 			Assists:          4,
	// 			Kddiff:           5,
	// 			Kd:               6,
	// 			Adr:              7,
	// 			Hsprecent:        8,
	// 			Firstkills:       9,
	// 			Firstdeaths:      10,
	// 			Tradekills:       11,
	// 			Tradedeaths:      12,
	// 			Tradefirstkills:  13,
	// 			Tradefirstdeaths: 14,
	// 			Roundswonv5:      15,
	// 			Roundswonv4:      16,
	// 			Roundswonv3:      17,
	// 			Rounds5K:         18,
	// 			Rounds4K:         19,
	// 			Rounds3K:         20,
	// 		},
	// 		{
	// 			Name:             "Player 2",
	// 			Clantag:          "megaclan3000",
	// 			AvatarURL:        "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
	// 			Rank:             7,
	// 			Steamid64:        1,
	// 			Kills:            2,
	// 			Deaths:           3,
	// 			Assists:          4,
	// 			Kddiff:           5,
	// 			Kd:               6,
	// 			Adr:              7,
	// 			Hsprecent:        8,
	// 			Firstkills:       9,
	// 			Firstdeaths:      10,
	// 			Tradekills:       11,
	// 			Tradedeaths:      12,
	// 			Tradefirstkills:  13,
	// 			Tradefirstdeaths: 14,
	// 			Roundswonv5:      15,
	// 			Roundswonv4:      16,
	// 			Roundswonv3:      17,
	// 			Rounds5K:         18,
	// 			Rounds4K:         19,
	// 			Rounds3K:         20,
	// 		},
	// 	},

	// 	Enemy: []ScoreboardPlayer{
	// 		{
	// 			Name:             "Player 3",
	// 			Clantag:          "megaclan3000",
	// 			AvatarURL:        "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
	// 			Rank:             7,
	// 			Steamid64:        1,
	// 			Kills:            2,
	// 			Deaths:           3,
	// 			Assists:          4,
	// 			Kddiff:           5,
	// 			Kd:               6,
	// 			Adr:              7,
	// 			Hsprecent:        8,
	// 			Firstkills:       9,
	// 			Firstdeaths:      10,
	// 			Tradekills:       11,
	// 			Tradedeaths:      12,
	// 			Tradefirstkills:  13,
	// 			Tradefirstdeaths: 14,
	// 			Roundswonv5:      15,
	// 			Roundswonv4:      16,
	// 			Roundswonv3:      17,
	// 			Rounds5K:         18,
	// 			Rounds4K:         19,
	// 			Rounds3K:         20,
	// 		},
	// 		{
	// 			Name:             "Player 4",
	// 			Clantag:          "megaclan3000",
	// 			AvatarURL:        "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
	// 			Rank:             7,
	// 			Steamid64:        1,
	// 			Kills:            2,
	// 			Deaths:           3,
	// 			Assists:          4,
	// 			Kddiff:           5,
	// 			Kd:               6,
	// 			Adr:              7,
	// 			Hsprecent:        8,
	// 			Firstkills:       9,
	// 			Firstdeaths:      10,
	// 			Tradekills:       11,
	// 			Tradedeaths:      12,
	// 			Tradefirstkills:  13,
	// 			Tradefirstdeaths: 14,
	// 			Roundswonv5:      15,
	// 			Roundswonv4:      16,
	// 			Roundswonv3:      17,
	// 			Rounds5K:         18,
	// 			Rounds4K:         19,
	// 			Rounds3K:         20,
	// 		},
	// 	},
	// }
	// return out
}

type ScoreboardPlayers struct {
	Players []ScoreboardPlayer
}

func (sp ScoreboardPlayers) Clan() []ScoreboardPlayer {

	out := []ScoreboardPlayer{}

	for _, p := range sp.Players {
		if p.IsClanMember {
			out = append(out, p)
		}
	}
	return out
}

func (sp ScoreboardPlayers) Enemy() []ScoreboardPlayer {

	out := []ScoreboardPlayer{}

	for _, p := range sp.Players {
		if !p.IsClanMember {
			out = append(out, p)
		}
	}

	return out
}

func (sb ScoreboardPlayers) PlayerByID(steamID uint64) (*ScoreboardPlayer, error) {

	for k := range sb.Players {
		if sb.Players[k].Steamid64 == steamID {
			return &sb.Players[k], nil
		}
	}

	return nil, errors.New("Could not find player by ID: " + strconv.FormatUint(steamID, 10))
}

type ScoreboardPlayer struct {
	WeaponStats map[common.EquipmentType]WeaponStat

	IsClanMember     bool    `json:"isclanmember"`
	Steamid64        uint64  `json:"steamid64"`
	Name             string  `json:"name"`
	Clantag          string  `json:"clantag"`
	AvatarURL        string  `json:"avatar_url"`
	Rank             int     `json:"rank"`
	Kills            int     `json:"kills"`
	Deaths           int     `json:"deaths"`
	Assists          int     `json:"assists"`
	Kddiff           int     `json:"kddiff"`
	Kd               float64 `json:"kd"`
	Adr              int     `json:"adr"`
	Hsprecent        int     `json:"hsprecent"`
	Firstkills       int     `json:"firstkills"`
	Firstdeaths      int     `json:"firstdeaths"`
	Tradekills       int     `json:"tradekills"`
	Tradedeaths      int     `json:"tradedeaths"`
	Tradefirstkills  int     `json:"tradefirstkills"`
	Tradefirstdeaths int     `json:"tradefirstdeaths"`
	Roundswonv5      int     `json:"roundswonv5"`
	Roundswonv4      int     `json:"roundswonv4"`
	Roundswonv3      int     `json:"roundswonv3"`
	Rounds5K         int     `json:"rounds5k"`
	Rounds4K         int     `json:"rounds4k"`
	Rounds3K         int     `json:"rounds3k"`
}

type Round struct {
	ScoreClan        int           `json:"score_clan"`
	ScoreEnemy       int           `json:"score_enemy"`
	WinReason        int           `json:"win_reason"`
	TotalDamageTaken int           `json:"total_damage_taken"`
	TotalDamageGiven int           `json:"total_damage_given"`
	WinnerTeam       common.Team   `json:"winner_team"`
	KillsClan        []RoundKill   `json:"kills_clan"`
	KillsEnemy       []RoundKill   `json:"kills_enemy"`
	Duration         time.Duration `json:"duration"`
}

func (is InfoStruct) GetRounds() []Round {

	if len(is.Players.Clan()) == 0 {
		log.Fatal("no players found for getrounds", is.Players.Clan(), is.Players.Enemy())
	}
	//TODO implement with real data
	return []Round{
		{

			ScoreClan:        2,
			ScoreEnemy:       1,
			WinReason:        1,
			TotalDamageTaken: 900,
			TotalDamageGiven: 1200,
			WinnerTeam:       common.TeamCounterTerrorists,
			KillsClan: []RoundKill{
				{
					Killer:             &is.Players.Clan()[0],
					Assister:           &is.Players.Clan()[3],
					Victim:             &is.Players.Enemy()[0],
					KillerTeamString:   "T",
					VictimTeamString:   "CT",
					AssisterTeamString: "CT",
					KillerWeapon:       common.EqAK47,
				},
				{
					Killer:             &is.Players.Clan()[1],
					Assister:           &is.Players.Clan()[3],
					Victim:             &is.Players.Enemy()[0],
					KillerTeamString:   "T",
					VictimTeamString:   "CT",
					AssisterTeamString: "CT",
					KillerWeapon:       common.EqAK47,
				},
			},
			KillsEnemy: []RoundKill{
				{
					Assister:           &is.Players.Enemy()[1],
					Killer:             &is.Players.Enemy()[3],
					Victim:             &is.Players.Clan()[0],
					KillerTeamString:   "T",
					VictimTeamString:   "CT",
					AssisterTeamString: "CT",
					KillerWeapon:       common.EqAK47,
				},
			},
			Duration: time.Hour,
		},
	}
}
