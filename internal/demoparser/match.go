package demoparser

import (
	"time"

	// "github.com/golang/geo/r3"
	common "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
	log "github.com/sirupsen/logrus"
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
	return is.Players
}

type ScoreboardPlayers struct {
	Players []ScoreboardPlayer
}

func (sp *ScoreboardPlayers) AddKill(steamID uint64) {
	for k := range sp.Players {
		if sp.Players[k].Steamid64 == steamID {
			sp.Players[k].Kills += 1
		}
	}
}

func (sp *ScoreboardPlayers) AddDeath(steamID uint64) {
	for k := range sp.Players {
		if sp.Players[k].Steamid64 == steamID {
			sp.Players[k].Deaths += 1
		}
	}
}

func (sp *ScoreboardPlayers) AddAssist(steamID uint64) {
	for k := range sp.Players {
		if sp.Players[k].Steamid64 == steamID {
			sp.Players[k].Assists += 1
		}
	}
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

func (p *MyParser) PlayerByID(player *common.Player) *ScoreboardPlayer {

	for _, v := range p.Match.Players.Players {
		if v.Steamid64 == player.SteamID64 {
			return &v
		}
	}

	log.Warning("Created new player for ID:", player.SteamID64)
	newplayer := p.NewScoreBoardPlayer(player)
	p.Match.Players.Players = append(p.Match.Players.Players, newplayer)

	return &newplayer
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

type ScoreboardRound struct {
	ClanWonRound     bool
	Duration         time.Duration         `json:"duration"`
	ClanKills        []RoundKill           `json:"kills_clan"`
	EnemyKills       []RoundKill           `json:"kills_enemy"`
	ScoreClan        int                   `json:"score_clan"`
	ScoreEnemy       int                   `json:"score_enemy"`
	ClanSurvivors    int                   `json:"survivivors_clan"`
	EnemySurvivors   int                   `json:"survivors_enemy"`
	TeamWon          common.Team           `json:"team_won"`
	TotalDamageGiven int                   `json:"total_damage_given"`
	TotalDamageTaken int                   `json:"total_damage_taken"`
	WinReason        events.RoundEndReason `json:"win_reason"`
	WinnerTeam       common.Team           `json:"winner_team"`
}

func (is InfoStruct) GetRounds() []ScoreboardRound {

	//TODO implement with real data
	return is.Rounds
	// return []Round{
	// 	{

	// 		ScoreClan:        2,
	// 		ScoreEnemy:       1,
	// 		WinReason:        1,
	// 		TotalDamageTaken: 900,
	// 		TotalDamageGiven: 1200,
	// 		WinnerTeam:       common.TeamCounterTerrorists,
	// 		KillsClan: []RoundKill{
	// 			{
	// 				Killer:             &is.Players.Clan()[0],
	// 				Assister:           &is.Players.Clan()[3],
	// 				Victim:             &is.Players.Enemy()[0],
	// 				KillerTeamString:   "T",
	// 				VictimTeamString:   "CT",
	// 				AssisterTeamString: "CT",
	// 				KillerWeapon:       common.EqAK47,
	// 			},
	// 			{
	// 				Killer:             &is.Players.Clan()[1],
	// 				Assister:           &is.Players.Clan()[3],
	// 				Victim:             &is.Players.Enemy()[0],
	// 				KillerTeamString:   "T",
	// 				VictimTeamString:   "CT",
	// 				AssisterTeamString: "CT",
	// 				KillerWeapon:       common.EqAK47,
	// 			},
	// 		},
	// 		KillsEnemy: []RoundKill{
	// 			{
	// 				Assister:           &is.Players.Enemy()[1],
	// 				Killer:             &is.Players.Enemy()[3],
	// 				Victim:             &is.Players.Clan()[0],
	// 				KillerTeamString:   "T",
	// 				VictimTeamString:   "CT",
	// 				AssisterTeamString: "CT",
	// 				KillerWeapon:       common.EqAK47,
	// 			},
	// 		},
	// 		Duration: time.Hour,
	// 	},
	// }
}
