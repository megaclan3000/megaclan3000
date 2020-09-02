package demoparser

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	// "github.com/golang/geo/r3"
	common "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
	"github.com/megaclan3000/megaclan3000/internal/steamclient"
	log "github.com/sirupsen/logrus"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type InfoStruct struct {
	General           ScoreboardGeneral
	Players           ScoreboardPlayers
	Rounds            []ScoreboardRound
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

type RoundKill struct {
	Time               time.Duration
	KillerTeamString   string
	VictimTeamString   string
	AssisterTeamString string
	IsHeadshot         bool
	Victim             *ScoreboardPlayer
	Killer             *ScoreboardPlayer
	Assister           *ScoreboardPlayer
	KillerWeapon       common.EquipmentType
}

func (is *InfoStruct) Weapons() interface{} {

	type wlist struct {
		//playername to amount
		Clan  map[string]int `json:"clan"`
		Enemy map[string]int `json:"enemy"`
	}

	type weapon struct {
		Kills     wlist `json:"kills"`
		Shots     wlist `json:"shots"`
		Headshots wlist `json:"headshots"`
		Accuracy  wlist `json:"accuracy"`
		Damage    wlist `json:"damage"`
		Hits      wlist `json:"hits"`
	}

	// Weapons           map[common.EquipmentType]map[*ScoreboardPlayer]WeaponStat
	ret := struct {
		// weaponname to stats
		Weapons map[string]*weapon `json:"weapons"`
	}{Weapons: make(map[string]*weapon)}

	var weapons []string
	for _, player := range is.Players.Players {
		for v := range player.WeaponStats.Kills {
			weapons = append(weapons, v.String())
		}
		for v := range player.WeaponStats.Headshots {
			weapons = append(weapons, v.String())
		}
		for v := range player.WeaponStats.Damage {
			weapons = append(weapons, v.String())
		}
		for v := range player.WeaponStats.Accuracy {
			weapons = append(weapons, v.String())
		}
		for v := range player.WeaponStats.Shots {
			weapons = append(weapons, v.String())
		}
		for v := range player.WeaponStats.Hits {
			weapons = append(weapons, v.String())
		}
	}

	for _, v := range weapons {
		ret.Weapons[v] = &weapon{
			Kills:     wlist{Clan: make(map[string]int), Enemy: make(map[string]int)},
			Headshots: wlist{Clan: make(map[string]int), Enemy: make(map[string]int)},
			Accuracy:  wlist{Clan: make(map[string]int), Enemy: make(map[string]int)},
			Damage:    wlist{Clan: make(map[string]int), Enemy: make(map[string]int)},
			Shots:     wlist{Clan: make(map[string]int), Enemy: make(map[string]int)},
			Hits:      wlist{Clan: make(map[string]int), Enemy: make(map[string]int)},
		}
	}

	for _, player := range is.Players.Players {

		for wep, kills := range player.WeaponStats.Kills {
			if player.IsClanMember {
				ret.Weapons[wep.String()].Kills.Clan[player.Name] = kills
			} else {
				ret.Weapons[wep.String()].Kills.Enemy[player.Name] = kills
			}
		}

		for weapon, hs := range player.WeaponStats.Headshots {
			if player.IsClanMember {
				ret.Weapons[weapon.String()].Headshots.Clan[player.Name] = hs
			} else {
				ret.Weapons[weapon.String()].Headshots.Enemy[player.Name] = hs
			}
		}

		for weapon, accuracy := range player.WeaponStats.Accuracy {
			if player.IsClanMember {
				ret.Weapons[weapon.String()].Accuracy.Clan[player.Name] = accuracy
			} else {
				ret.Weapons[weapon.String()].Accuracy.Enemy[player.Name] = accuracy
			}
		}

		for weapon, damage := range player.WeaponStats.Damage {
			if player.IsClanMember {
				ret.Weapons[weapon.String()].Damage.Clan[player.Name] = damage
			} else {
				ret.Weapons[weapon.String()].Damage.Enemy[player.Name] = damage
			}
		}

		for weapon, shots := range player.WeaponStats.Shots {
			if player.IsClanMember {
				ret.Weapons[weapon.String()].Shots.Clan[player.Name] = shots
			} else {
				ret.Weapons[weapon.String()].Shots.Enemy[player.Name] = shots
			}
		}

		for weapon, hits := range player.WeaponStats.Hits {
			if player.IsClanMember {
				ret.Weapons[weapon.String()].Hits.Clan[player.Name] = hits
			} else {
				ret.Weapons[weapon.String()].Hits.Enemy[player.Name] = hits
			}
		}
	}

	return ret
}

func (rk *RoundKill) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Time               time.Duration        `json:"time"`
		KillerTeamString   string               `json:"killer_team_string"`
		VictimTeamString   string               `json:"victim_team_string"`
		AssisterTeamString string               `json:"assister_team_string"`
		IsHeadshot         bool                 `json:"is_headshot"`
		Victim             *ScoreboardPlayer    `json:"victim"`
		Killer             *ScoreboardPlayer    `json:"killer"`
		Assister           *ScoreboardPlayer    `json:"assister"`
		KillerWeapon       common.EquipmentType `json:"weapon"`
		KillerWeaponName   string               `json:"weapon_name"`
	}{

		Time:               rk.Time,
		KillerTeamString:   rk.KillerTeamString,
		VictimTeamString:   rk.VictimTeamString,
		AssisterTeamString: rk.AssisterTeamString,
		IsHeadshot:         rk.IsHeadshot,
		Victim:             rk.Victim,
		Killer:             rk.Killer,
		Assister:           rk.Assister,
		KillerWeapon:       rk.KillerWeapon,
		KillerWeaponName:   rk.KillerWeapon.String(),
	})
}

type WeaponUser struct {
	Kills     int
	HSPercent int
	Accuracy  int
	Damage    int
}

type MegacoinPlayer struct {
	//TODO
	ForCriteriaA int
	ForCriteriaB int
	ForCriteriaC int
}

func GetMatchInfo(id string, steamClient *steamclient.SteamClient) (InfoStruct, error) {
	//TODO
	p := NewMyParser(steamClient)
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

func (sp ScoreboardPlayers) PlayerNumByID(steamID uint64) (int, error) {
	for k, v := range sp.Players {
		if v.Steamid64 == steamID {
			return k, nil
		}
	}
	return 0, errors.New("Player Number not found" + strconv.FormatUint(steamID, 10))
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

type WeaponStats struct {

	// Number of kills
	Kills map[common.EquipmentType]int

	// Number of headshots
	Headshots map[common.EquipmentType]int

	// Percent shots hit of shots fired
	Accuracy map[common.EquipmentType]int

	// Damage caused
	Damage map[common.EquipmentType]int

	// Shots fired
	Shots map[common.EquipmentType]int

	// Shots hit
	Hits map[common.EquipmentType]int
}

type ScoreboardPlayer struct {
	WeaponStats WeaponStats

	IsBot            bool    `json:"isbot"`
	IsClanMember     bool    `json:"isclanmember"`
	Steamid64        uint64  `json:"steamid64"`
	Name             string  `json:"name"`
	Clantag          string  `json:"clantag"`
	AvatarURL        string  `json:"avatar_url"`
	Rank             int     `json:"rank"`
	Kills            int     `json:"kills"`
	MVPs             int     `json:"mvps"`
	Deaths           int     `json:"deaths"`
	Assists          int     `json:"assists"`
	Kd               float64 `json:"kd"`
	Adr              int     `json:"adr"`
	Headshots        int     `json:"headshots"`
	Hsprecent        float64 `json:"hsprecent"`
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
