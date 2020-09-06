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

func allWeapons() []common.EquipmentType {

	return []common.EquipmentType{
		common.EqUnknown,
		common.EqP2000,
		common.EqGlock,
		common.EqP250,
		common.EqDeagle,
		common.EqFiveSeven,
		common.EqDualBerettas,
		common.EqTec9,
		common.EqCZ,
		common.EqUSP,
		common.EqRevolver,
		common.EqMP7,
		common.EqMP9,
		common.EqBizon,
		common.EqMac10,
		common.EqUMP,
		common.EqP90,
		common.EqMP5,
		common.EqSawedOff,
		common.EqNova,
		common.EqMag7,
		common.EqSwag7,
		common.EqXM1014,
		common.EqM249,
		common.EqNegev,
		common.EqGalil,
		common.EqFamas,
		common.EqAK47,
		common.EqM4A4,
		common.EqM4A1,
		common.EqScout,
		common.EqSSG08,
		common.EqSG556,
		common.EqSG553,
		common.EqAUG,
		common.EqAWP,
		common.EqScar20,
		common.EqG3SG1,
		common.EqZeus,
		common.EqKevlar,
		common.EqHelmet,
		common.EqBomb,
		common.EqKnife,
		common.EqDefuseKit,
		common.EqWorld,
		common.EqDecoy,
		common.EqMolotov,
		common.EqIncendiary,
		common.EqFlash,
		common.EqSmoke,
		common.EqHE,
	}
}

func NewWeaponstats() WeaponStats {

	return WeaponStats{
		kills:     make(map[common.EquipmentType]int),
		headshots: make(map[common.EquipmentType]int),
		accuracy:  make(map[common.EquipmentType]int),
		damage:    make(map[common.EquipmentType]int),
		shots:     make(map[common.EquipmentType]int),
		hits:      make(map[common.EquipmentType]int),
	}
}

func NewPlayerDamages() PlayerDamages {

	return PlayerDamages{
		Damages: make(map[uint64]int),
	}
}

func (is *InfoStruct) Damages() interface{} {

	type pdamage struct {
		Victim string `json:"victim"`
		Amount int    `json:"amount"`
	}

	type pdamages struct {
		PlayerName string    `json:"player"`
		AvatarURL  string    `json:"avatar"`
		Damages    []pdamage `json:"damages"`
	}

	ret := struct {
		Clan  map[string]pdamages `json:"clan"`
		Enemy map[string]pdamages `json:"enemy"`
	}{
		Clan:  make(map[string]pdamages),
		Enemy: make(map[string]pdamages),
	}

	for _, player := range is.Players.Players {

		if player.IsBot {
			continue
		}

		// Prefill with zero damage for all players except BOTs
		dams := make(map[string]int)

		for _, p2 := range is.Players.Players {
			if p2.IsBot {
				continue
			}
			dams[p2.Name] = 0
			log.Warning("adding dam for: ", p2.Name)
		}

		for k2, v := range player.PlayerDamages.Damages {
			vicNum, err := is.Players.PlayerNumByID(k2)
			if err != nil {
				panic(err)
			}

			if is.Players.Players[vicNum].IsBot {
				continue
			}

			name := is.Players.Players[vicNum].Name

			log.Warning("Adding damage player", player.Name, "-> ", name, ": ", v)
			dams[name] = v
		}

		tmp := pdamages{
			PlayerName: player.Name,
			AvatarURL:  player.AvatarURL,
		}

		for k, v := range dams {
			tmp.Damages = append(tmp.Damages, pdamage{
				Victim: k,
				Amount: v,
			})
		}
		if player.IsClanMember {
			ret.Clan[player.Name] = tmp
		} else {
			ret.Enemy[player.Name] = tmp
		}
	}

	return ret
}

func (is *InfoStruct) Weapons() interface{} {

	type wlist struct {
		//playername to amount
		Clan  map[string]int `json:"clan"`
		Enemy map[string]int `json:"enemy"`
	}

	type weapon struct {
		Name string `json:"name"`

		TotalKills     int   `json:"total_kills"`
		TotalShots     int   `json:"total_shots"`
		TotalHeadshots int   `json:"total_headshots"`
		TotalAccuracy  int   `json:"total_accuracy"`
		TotalDamage    int   `json:"total_damage"`
		TotalHits      int   `json:"total_hits"`
		Kills          wlist `json:"kills"`
		Shots          wlist `json:"shots"`
		Headshots      wlist `json:"headshots"`
		Accuracy       wlist `json:"accuracy"`
		Damage         wlist `json:"damage"`
		Hits           wlist `json:"hits"`
	}

	// Weapons           map[common.EquipmentType]map[*ScoreboardPlayer]WeaponStat
	ret := struct {
		// weaponname to stats
		Weapons map[string]*weapon `json:"weapons"`
	}{Weapons: make(map[string]*weapon)}

	for _, v := range is.Players.AllWeaponsUsed() {

		// Skip non-weapon classes
		if v.Class() == common.EqClassUnknown || v.Class() == common.EqClassEquipment {
			continue
		}

		ret.Weapons[v.String()] = &weapon{
			Name:      v.String(),
			Kills:     wlist{Clan: make(map[string]int), Enemy: make(map[string]int)},
			Headshots: wlist{Clan: make(map[string]int), Enemy: make(map[string]int)},
			Accuracy:  wlist{Clan: make(map[string]int), Enemy: make(map[string]int)},
			Damage:    wlist{Clan: make(map[string]int), Enemy: make(map[string]int)},
			Shots:     wlist{Clan: make(map[string]int), Enemy: make(map[string]int)},
			Hits:      wlist{Clan: make(map[string]int), Enemy: make(map[string]int)},
		}

		for _, player := range is.Players.Players {

			ret.Weapons[v.String()].TotalKills += player.WeaponStats.Kills(v)
			ret.Weapons[v.String()].TotalHeadshots += player.WeaponStats.Headshots(v)
			ret.Weapons[v.String()].TotalDamage += player.WeaponStats.Damage(v)
			ret.Weapons[v.String()].TotalShots += player.WeaponStats.Shots(v)
			ret.Weapons[v.String()].TotalHits += player.WeaponStats.Hits(v)
			// ret.Weapons[v.String()].TotalDamage+= player.WeaponStats.Damage(v)

			if player.IsClanMember {
				ret.Weapons[v.String()].Kills.Clan[player.Name] = player.WeaponStats.Kills(v)
				ret.Weapons[v.String()].Headshots.Clan[player.Name] = player.WeaponStats.Headshots(v)
				ret.Weapons[v.String()].Accuracy.Clan[player.Name] = player.WeaponStats.Accuracy(v)
				ret.Weapons[v.String()].Damage.Clan[player.Name] = player.WeaponStats.Damage(v)
				ret.Weapons[v.String()].Shots.Clan[player.Name] = player.WeaponStats.Shots(v)
				ret.Weapons[v.String()].Hits.Clan[player.Name] = player.WeaponStats.Hits(v)
			} else {
				ret.Weapons[v.String()].Kills.Enemy[player.Name] = player.WeaponStats.Kills(v)
				ret.Weapons[v.String()].Headshots.Enemy[player.Name] = player.WeaponStats.Headshots(v)
				ret.Weapons[v.String()].Accuracy.Enemy[player.Name] = player.WeaponStats.Accuracy(v)
				ret.Weapons[v.String()].Damage.Enemy[player.Name] = player.WeaponStats.Damage(v)
				ret.Weapons[v.String()].Shots.Enemy[player.Name] = player.WeaponStats.Shots(v)
				ret.Weapons[v.String()].Hits.Enemy[player.Name] = player.WeaponStats.Hits(v)
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

// AllWeaponsUsed returns all weapons shot at least once during the match
func (sp *ScoreboardPlayers) AllWeaponsUsed() []common.EquipmentType {
	list := []common.EquipmentType{}

	for _, w := range allWeapons() {
		for _, p := range sp.Players {
			if p.WeaponStats.Shots(w) > 0 {
				list = append(list, w)
			}
		}
	}
	return list
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
	kills map[common.EquipmentType]int

	// Number of headshots
	headshots map[common.EquipmentType]int

	// Percent shots hit of shots fired
	accuracy map[common.EquipmentType]int

	// Damage caused
	damage map[common.EquipmentType]int

	// Shots fired
	shots map[common.EquipmentType]int

	// Shots hit
	hits map[common.EquipmentType]int
}

func (ws *WeaponStats) AddKill(e events.Kill) {
	ws.kills[e.Weapon.Type]++
}

func (ws *WeaponStats) AddHeadshot(e events.Kill) {
	if e.IsHeadshot {
		ws.headshots[e.Weapon.Type]++
	}
}

func (ws *WeaponStats) AddDamage(e events.PlayerHurt) {
	ws.damage[e.Weapon.Type] += e.HealthDamage
}

func (ws *WeaponStats) AddShot(e events.WeaponFire) {
	ws.shots[e.Weapon.Type]++
	ws.accuracy[e.Weapon.Type] = (ws.Hits(e.Weapon.Type) * 100) / ws.Shots(e.Weapon.Type)
}

func (ws *WeaponStats) AddHit(e events.PlayerHurt) {
	ws.hits[e.Weapon.Type]++
}

func (ws WeaponStats) Kills(w common.EquipmentType) int {
	return ws.kills[w]
}

func (ws WeaponStats) Accuracy(w common.EquipmentType) int {
	return ws.accuracy[w]
}

func (ws WeaponStats) Headshots(w common.EquipmentType) int {
	return ws.headshots[w]
}

func (ws WeaponStats) Damage(w common.EquipmentType) int {
	return ws.damage[w]
}

func (ws WeaponStats) Shots(w common.EquipmentType) int {
	return ws.shots[w]
}

func (ws WeaponStats) Hits(w common.EquipmentType) int {
	return ws.hits[w]
}

func (sp *ScoreboardPlayer) AddDamage(damage int, victim *ScoreboardPlayer) {
	sp.PlayerDamages.Damages[victim.Steamid64] += damage
}

type PlayerDamages struct {
	Damages map[uint64]int
}

type ScoreboardPlayer struct {
	WeaponStats      WeaponStats
	PlayerDamages    PlayerDamages
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
