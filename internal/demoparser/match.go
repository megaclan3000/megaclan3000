package demoparser

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	// "github.com/golang/geo/r3"
	common "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
	"github.com/megaclan3000/megaclan3000/internal/steamclient"
	log "github.com/sirupsen/logrus"
)

// InfoStruct holds the data of one demo file in it's parsed form
type InfoStruct struct {
	MatchID    string            `json:"match_id"     db:"match_id"`
	MatchValid bool              `json:"match_valid"  db:"match_valid"`
	General    ScoreboardGeneral `json:"general"      db:"general"`
	Players    ScoreboardPlayers `json:"players"      db:"players"`
	Rounds     []ScoreboardRound `json:"rounds"       db:"rounds"`

	// Duels             [][]int
	// HeatmapsImageURLs []string
	// Megacoins         []MegacoinPlayer
}

// Value : Make the InfoStruct struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (is InfoStruct) Value() (driver.Value, error) {
	return json.Marshal(is)
}

// Scan : Make the InfoStruct struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (is *InfoStruct) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &is)
}

// ScoreboardGeneral holds general information about the match
type ScoreboardGeneral struct {
	ClanWonMatch  bool          `json:"clan_won_round" db:"clan_won_round"`
	ScoreClan     int           `json:"score_clan"     db:"score_clan"`
	ScoreEnemy    int           `json:"score_enemy"    db:"score_enemy"`
	MapName       string        `json:"map_name"       db:"map_name"`
	MapIconURL    string        `json:"map_icon_url"   db:"map_icon_url"`
	UploadTime    time.Time     `json:"upload_time"    db:"upload_time"`
	MatchDuration time.Duration `json:"match_duration" db:"match_duration"`
	DemoLinkURL   string        `json:"demo_link_url"  db:"demo_link_url"`
}

// RoundKill holds information about a kill that happenend during the match
type RoundKill struct {
	Time               time.Duration        `json:"time"                 db:"time"`
	KillerTeamString   string               `json:"killer_team_string"   db:"killer_team_string"`
	VictimTeamString   string               `json:"victim_team_string"   db:"victim_team_string"`
	AssisterTeamString string               `json:"assister_team_string" db:"assister_team_string"`
	IsHeadshot         bool                 `json:"is_headshot"          db:"is_headshot"`
	Victim             *ScoreboardPlayer    `json:"victim"               db:"victim"`
	Killer             *ScoreboardPlayer    `json:"killer"               db:"killer"`
	Assister           *ScoreboardPlayer    `json:"assister"             db:"assister"`
	KillerWeapon       common.EquipmentType `json:"weapon"               db:"weapon"`
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

// NewWeaponstats constructor weaponstats
func NewWeaponstats() WeaponStats {

	return WeaponStats{
		Kills:     make(map[common.EquipmentType]int),
		Headshots: make(map[common.EquipmentType]int),
		Accuracy:  make(map[common.EquipmentType]int),
		Damage:    make(map[common.EquipmentType]int),
		Shots:     make(map[common.EquipmentType]int),
		Hits:      make(map[common.EquipmentType]int),
	}
}

// NewPlayerDamages constructor for player damages struct
func NewPlayerDamages() PlayerDamages {

	return PlayerDamages{
		Damages: make(map[uint64]int),
	}
}

// Damages returns all damages that where dealt during a match
func (is *InfoStruct) Damages() interface{} {

	type pdamage struct {
		Victim string `json:"victim" db:"victim"`
		Amount int    `json:"amount" db:"amount"`
	}

	type pdamages struct {
		PlayerName string    `json:"player"  db:"player"`
		AvatarURL  string    `json:"avatar"  db:"avatar"`
		Damages    []pdamage `json:"damages" db:"damages"`
	}

	ret := struct {
		Clan  map[string]pdamages `json:"clan"  db:"clan"`
		Enemy map[string]pdamages `json:"enemy" db:"enemy"`
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

// Weapons lists all weapons used in the match with it's stats
func (is *InfoStruct) Weapons() interface{} {

	type wlist struct {
		//playername to amount
		Clan  map[string]int `json:"clan"  db:"clan"`
		Enemy map[string]int `json:"enemy" db:"enemy"`
	}

	type weapon struct {
		Name           string `json:"name"            db:"name"`
		TotalKills     int    `json:"total_kills"     db:"total_kills"`
		TotalShots     int    `json:"total_shots"     db:"total_shots"`
		TotalHeadshots int    `json:"total_headshots" db:"total_headshots"`
		TotalAccuracy  int    `json:"total_accuracy"  db:"total_accuracy"`
		TotalDamage    int    `json:"total_damage"    db:"total_damage"`
		TotalHits      int    `json:"total_hits"      db:"total_hits"`
		Kills          wlist  `json:"kills"           db:"kills"`
		Shots          wlist  `json:"shots"           db:"shots"`
		Headshots      wlist  `json:"headshots"       db:"headshots"`
		Accuracy       wlist  `json:"accuracy"        db:"accuracy"`
		Damage         wlist  `json:"damage"          db:"damage"`
		Hits           wlist  `json:"hits"            db:"hits"`
	}

	// Weapons           map[common.EquipmentType]map[*ScoreboardPlayer]WeaponStat
	ret := struct {
		// weaponname to stats
		Weapons map[string]*weapon `json:"weapons" db:"weapons"`
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

			ret.Weapons[v.String()].TotalKills += player.WeaponStats.getKills(v)
			ret.Weapons[v.String()].TotalHeadshots += player.WeaponStats.getHeadshots(v)
			ret.Weapons[v.String()].TotalDamage += player.WeaponStats.getDamage(v)
			ret.Weapons[v.String()].TotalShots += player.WeaponStats.getShots(v)
			ret.Weapons[v.String()].TotalHits += player.WeaponStats.getHits(v)
			// ret.Weapons[v.String()].TotalDamage+= player.WeaponStats.Damage(v)

			if player.IsClanMember {
				ret.Weapons[v.String()].Kills.Clan[player.Name] = player.WeaponStats.getKills(v)
				ret.Weapons[v.String()].Headshots.Clan[player.Name] = player.WeaponStats.getHeadshots(v)
				ret.Weapons[v.String()].Accuracy.Clan[player.Name] = player.WeaponStats.getAccuracy(v)
				ret.Weapons[v.String()].Damage.Clan[player.Name] = player.WeaponStats.getDamage(v)
				ret.Weapons[v.String()].Shots.Clan[player.Name] = player.WeaponStats.getShots(v)
				ret.Weapons[v.String()].Hits.Clan[player.Name] = player.WeaponStats.getHits(v)
			} else {
				ret.Weapons[v.String()].Kills.Enemy[player.Name] = player.WeaponStats.getKills(v)
				ret.Weapons[v.String()].Headshots.Enemy[player.Name] = player.WeaponStats.getHeadshots(v)
				ret.Weapons[v.String()].Accuracy.Enemy[player.Name] = player.WeaponStats.getAccuracy(v)
				ret.Weapons[v.String()].Damage.Enemy[player.Name] = player.WeaponStats.getDamage(v)
				ret.Weapons[v.String()].Shots.Enemy[player.Name] = player.WeaponStats.getShots(v)
				ret.Weapons[v.String()].Hits.Enemy[player.Name] = player.WeaponStats.getHits(v)
			}

		}
	}

	return ret
}

// MarshalJSON marshals a RoundKill struct to json, e.g. for the api
func (rk *RoundKill) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Time               time.Duration        `json:"time"                 db:"time"`
		KillerTeamString   string               `json:"killer_team_string"   db:"killer_team_string"`
		VictimTeamString   string               `json:"victim_team_string"   db:"victim_team_string"`
		AssisterTeamString string               `json:"assister_team_string" db:"assister_team_string"`
		IsHeadshot         bool                 `json:"is_headshot"          db:"is_headshot"`
		Victim             *ScoreboardPlayer    `json:"victim"               db:"victim"`
		Killer             *ScoreboardPlayer    `json:"killer"               db:"killer"`
		Assister           *ScoreboardPlayer    `json:"assister"             db:"assister"`
		KillerWeapon       common.EquipmentType `json:"weapon"               db:"weapon"`
		KillerWeaponName   string               `json:"weapon_name"          db:"weapon_name"`
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

// GetMatchInfo parses a demo file and returns a infostruct containing it's data
func GetMatchInfo(path string, steamClient *steamclient.SteamClient) (InfoStruct, error) {
	//TODO
	p := NewDemoParser(steamClient)
	var info InfoStruct
	err := p.Parse(path, &info)
	return info, err
}

// GetScoreboard returns the scoreboard of match
func (is InfoStruct) GetScoreboard() ScoreboardPlayers {
	return is.Players
}

// PlayerNumByID returns the a players's position in the Players[] list of a
// match
func (sp ScoreboardPlayers) PlayerNumByID(steamID uint64) (int, error) {
	for k, v := range sp.Players {
		if v.Steamid64 == steamID {
			return k, nil
		}
	}
	return 0, errors.New("Player Number not found" + strconv.FormatUint(steamID, 10))
}

// ScoreboardPlayers thin wrapper around the list of players to implment methods
type ScoreboardPlayers struct {
	Players []ScoreboardPlayer `json:"players" db:"players"`
}

// AllWeaponsUsed returns all weapons shot at least once during the match
func (sp *ScoreboardPlayers) AllWeaponsUsed() []common.EquipmentType {
	list := []common.EquipmentType{}

	for _, w := range allWeapons() {
		for _, p := range sp.Players {
			if p.WeaponStats.getShots(w) > 0 {
				list = append(list, w)
			}
		}
	}

	return list
}

func (sp *ScoreboardPlayers) addAssist(steamID uint64) {
	for k := range sp.Players {
		if sp.Players[k].Steamid64 == steamID {
			sp.Players[k].Assists++
		}
	}
}

// Clan returns all players of the clan team
func (sp ScoreboardPlayers) Clan() []ScoreboardPlayer {

	out := []ScoreboardPlayer{}

	for _, p := range sp.Players {
		if p.IsClanMember {
			out = append(out, p)
		}
	}
	return out
}

// Enemy returns all players of the enemy team
func (sp ScoreboardPlayers) Enemy() []ScoreboardPlayer {

	out := []ScoreboardPlayer{}

	for _, p := range sp.Players {
		if !p.IsClanMember {
			out = append(out, p)
		}
	}

	return out
}

func (p *DemoParser) playerByID(player *common.Player) *ScoreboardPlayer {

	if player == nil {
		return nil
	}

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

// WeaponStats maps weapons to the player stats
type WeaponStats struct {

	// Number of Kills
	Kills map[common.EquipmentType]int `json:"kills" db:"kills"`

	// Number of Headshots
	Headshots map[common.EquipmentType]int `json:"headshots" db:"headshots"`

	// Percent shots hit of shots fired
	Accuracy map[common.EquipmentType]int `json:"accuracy" db:"accuracy"`

	// Damage caused
	Damage map[common.EquipmentType]int `json:"damage" db:"damage"`

	// Shots fired
	Shots map[common.EquipmentType]int `json:"shots" db:"shots"`

	// Shots hit
	Hits map[common.EquipmentType]int `json:"hits" db:"hits"`
}

func (ws *WeaponStats) addKill(e events.Kill) {
	ws.Kills[e.Weapon.Type]++
}

func (ws *WeaponStats) addHeadshot(e events.Kill) {
	if e.IsHeadshot {
		ws.Headshots[e.Weapon.Type]++
	}
}

func (ws *WeaponStats) addDamage(e events.PlayerHurt) {
	ws.Damage[e.Weapon.Type] += e.HealthDamage
}

func (ws *WeaponStats) addShot(e events.WeaponFire) {
	ws.Shots[e.Weapon.Type]++
	ws.Accuracy[e.Weapon.Type] = (ws.getHits(e.Weapon.Type) * 100) / ws.getShots(e.Weapon.Type)
}

func (ws *WeaponStats) addHit(e events.PlayerHurt) {
	ws.Hits[e.Weapon.Type]++
}

func (ws WeaponStats) getKills(w common.EquipmentType) int {
	return ws.Kills[w]
}

// getAccuracy returns
func (ws WeaponStats) getAccuracy(w common.EquipmentType) int {
	return ws.Accuracy[w]
}

func (ws WeaponStats) getHeadshots(w common.EquipmentType) int {
	return ws.Headshots[w]
}

func (ws WeaponStats) getDamage(w common.EquipmentType) int {
	return ws.Damage[w]
}

func (ws WeaponStats) getShots(w common.EquipmentType) int {
	return ws.Shots[w]
}

func (ws WeaponStats) getHits(w common.EquipmentType) int {
	return ws.Hits[w]
}

func (sp *ScoreboardPlayer) addDamage(damage int, victim *ScoreboardPlayer) {
	sp.PlayerDamages.Damages[victim.Steamid64] += damage
}

// PlayerDamages holds the damage of a player to each other player
type PlayerDamages struct {
	Damages map[uint64]int `json:"damages" db:"damages"`
}

// ScoreboardPlayer holds the information about the player of a match
type ScoreboardPlayer struct {
	WeaponStats      WeaponStats   `json:"weapon_stats" db:"weapon_stats"`
	PlayerDamages    PlayerDamages `json:"player_damages" db:"player_damages"`
	IsBot            bool          `json:"isbot" db:"isbot"`
	IsClanMember     bool          `json:"isclanmember" db:"isclanmember"`
	Steamid64        uint64        `json:"steamid64" db:"steamid64"`
	Name             string        `json:"name" db:"name"`
	Clantag          string        `json:"clantag" db:"clantag"`
	AvatarURL        string        `json:"avatar_url" db:"avatar_url"`
	Rank             int           `json:"rank" db:"rank"`
	Kills            int           `json:"kills" db:"kills"`
	MVPs             int           `json:"mvps" db:"mvps"`
	Deaths           int           `json:"deaths" db:"deaths"`
	Assists          int           `json:"assists" db:"assists"`
	Kd               float64       `json:"kd" db:"kd"`
	Adr              int           `json:"adr" db:"adr"`
	Headshots        int           `json:"headshots" db:"headshots"`
	Hsprecent        float64       `json:"hsprecent" db:"hsprecent"`
	Firstkills       int           `json:"firstkills" db:"firstkills"`
	Firstdeaths      int           `json:"firstdeaths" db:"firstdeaths"`
	Tradekills       int           `json:"tradekills" db:"tradekills"`
	Tradedeaths      int           `json:"tradedeaths" db:"tradedeaths"`
	Tradefirstkills  int           `json:"tradefirstkills" db:"tradefirstkills"`
	Tradefirstdeaths int           `json:"tradefirstdeaths" db:"tradefirstdeaths"`
	Roundswonv5      int           `json:"roundswonv5" db:"roundswonv5"`
	Roundswonv4      int           `json:"roundswonv4" db:"roundswonv4"`
	Roundswonv3      int           `json:"roundswonv3" db:"roundswonv3"`
	Rounds5K         int           `json:"rounds5k" db:"rounds5k"`
	Rounds4K         int           `json:"rounds4k" db:"rounds4k"`
	Rounds3K         int           `json:"rounds3k" db:"rounds3k"`
}

// ScoreboardRound holds the information about a round in a match
type ScoreboardRound struct {
	ClanWonRound     bool                  `json:"clan_won_round" db:"clan_won_round"`
	Duration         time.Duration         `json:"duration" db:"duration"`
	ClanKills        []RoundKill           `json:"kills_clan" db:"kills_clan"`
	EnemyKills       []RoundKill           `json:"kills_enemy" db:"kills_enemy"`
	ScoreClan        int                   `json:"score_clan" db:"score_clan"`
	ScoreEnemy       int                   `json:"score_enemy" db:"score_enemy"`
	ClanSurvivors    int                   `json:"survivivors_clan" db:"survivivors_clan"`
	EnemySurvivors   int                   `json:"survivors_enemy" db:"survivors_enemy"`
	TeamWon          common.Team           `json:"team_won" db:"team_won"`
	TotalDamageGiven int                   `json:"total_damage_given" db:"total_damage_given"`
	TotalDamageTaken int                   `json:"total_damage_taken" db:"total_damage_taken"`
	WinReason        events.RoundEndReason `json:"win_reason" db:"win_reason"`
	WinnerTeam       common.Team           `json:"winner_team" db:"winner_team"`
}
