package demoparser

import (
	"time"

	"github.com/golang/geo/r3"
	common "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Match represents a complete match (both sides) as parsed from a demo file.
// It mainly consists of a slice of rounds that happened during that match.

type Rounds map[int]*Round

type Match struct {
	// The ID of the match. Since there seems to be no real id in the demo
	// files, we use a hash of the header
	Hash uint64

	// The time the demo was uploaded. There seems to be no information about
	// when the match took place in the demo files
	UploadTime time.Time

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

	// PlayerHurt []events.PlayerHurt
}

// func (m Match) MarshalBSON() ([]byte, error) {
// 	return bson.Marshal(NewBsonMatch(m))
// }

type BsonRound struct {
	TimeStart    time.Duration `bson:"timestart"`
	TimeEnd      time.Duration `bson:"timeend"`
	TeamWon      common.Team   `bson:"teamwon"`
	BombPlanted  bool          `bson:"bombplanted"`
	BombDefused  bool          `bson:"bombdefused"`
	BombExploded bool          `bson:"bombexploded"`
	Kills        []BsonKill    `bson:"kills"`
	PlayersCT    []uint64      `bson:"players_ct"`
	PlayersT     []uint64      `bson:"players_t"`
	// PlayerEndStates map[primitive.ObjectID]PlayerState `bson:"playerendstates"`
	// PlayerHurts      []primitive.ObjectID `bson:"playerhurts"`
	//TODO
}

// type PlayerState struct {
// 	Inventory   map[int]*common.Equipment // All weapons / equipment the player is currently carrying. See also Weapons().
// 	AmmoLeft    [32]int                   // Ammo left for special weapons (e.g. grenades), index corresponds Equipment.AmmoType
// 	Team        common.Team               // Team identifier for the player (e.g. TeamTerrorists or TeamCounterTerrorists).
// 	IsConnected bool
// 	IsDefusing  bool
// 	IsPlanting  bool
// 	IsReloading bool
// 	IsUnknown   bool // Used to identify unknown/broken players. see https://github.com/markus-wa/demoinfocs-golang/issues/162
// }

type BsonKill struct {
	ObjectID primitive.ObjectID `bson:"_id"`
	// RoundID              primitive.ObjectID   `bson:"round"`
	WeaponType           common.EquipmentType `bson:"weapontype"`           // The type of weapon which the equipment instantiates.
	WeaponOriginalString string               `bson:"weaponoriginalstring"` // E.g. 'models/weapons/w_rif_m4a1_s.mdl'. Used internally to differentiate alternative weapons (M4A4 / M4A1-S etc.).
	Victim               uint64               `bson:"victim"`
	Killer               uint64               `bson:"killer"`
	Assister             uint64               `bson:"assister"`
	PenetratedObjects    int                  `bson:"penetratedobjects"`
	IsHeadshot           bool                 `bson:"isheadshot"`
}

type BsonPlayer struct {
	ObjectID          primitive.ObjectID `bson:"_id"`               // 64-bit representation of the user's Steam ID. See https://developer.valvesoftware.com/wiki/SteamID
	SteamID64         uint64             `bson:"steamid"`           // 64-bit representation of the user's Steam ID. See https://developer.valvesoftware.com/wiki/SteamID
	LastAlivePosition r3.Vector          `bson:"lastaliveposition"` // The location where the player was last alive. Should be equal to Position if the player is still alive.
	UserID            int                `bson:"userid"`            // Mostly used in game-events to address this player
	Name              string             `bson:"name"`              // Steam / in-game user name
	EntityID          int                `bson:"entityid"`          // Usually the same as Entity.ID() but may be different between player death and re-spawn.
	FlashDuration     float32            `bson:"flashduration"`     // Blindness duration from the flashbang currently affecting the player (seconds)
	FlashTick         int                `bson:"flashtick"`         // In-game tick at which the player was last flashed
	IsBot             bool               `bson:"isbot"`             // True if this is a bot-entity. See also IsControllingBot and ControlledBot().
	//TODO
}

type BsonMatch struct {
	ObjectID   primitive.ObjectID   `bson:"_id"`
	UploadTime time.Time            `bson:"uploadtime"`
	Map        string               `bson:"map"`
	Players    []uint64             `bson:"players"`
	Rounds     []primitive.ObjectID `bson:"rounds"`
	// PlayerHurts []BsonPlayerHurt `bson:"playerhurts"`
}

func NewBsonPlayer(p common.Player) BsonPlayer {
	return BsonPlayer{
		ObjectID:          primitive.NewObjectID(),
		SteamID64:         p.SteamID64,
		LastAlivePosition: p.LastAlivePosition,
		UserID:            p.UserID,
		Name:              p.Name,
		EntityID:          p.EntityID,
		FlashDuration:     p.FlashDuration,
		FlashTick:         p.FlashTick,
		IsBot:             p.IsBot,
	}
}

func NewBsonRound(r Round) BsonRound {
	//TODO

	var kills []BsonKill
	var playersT []uint64
	var playersCT []uint64
	// var bots []uint64

	for _, v := range r.Kills {
		kills = append(kills, NewBsonKill(v))
	}

	for _, p := range r.Players {
		switch p.Team {
		case common.TeamCounterTerrorists:
			playersCT = append(playersCT, p.SteamID64)
		case common.TeamTerrorists:
			playersT = append(playersT, p.SteamID64)
		}
	}

	return BsonRound{
		TimeStart:    r.TimeStart,
		TimeEnd:      r.TimeEnd,
		TeamWon:      r.TeamWon,
		BombPlanted:  r.BombPlanted,
		BombDefused:  r.BombDefused,
		BombExploded: r.BombExploded,
		Kills:        kills,
		PlayersCT:    playersCT,
		PlayersT:     playersT,
		// PlayerEndStates: states,
	}
}

func NewBsonKill(k events.Kill) BsonKill {

	var assister uint64

	if k.Assister != nil {
		assister = k.Assister.SteamID64
	}

	return BsonKill{
		ObjectID:             primitive.NewObjectID(),
		WeaponType:           k.Weapon.Type,
		WeaponOriginalString: k.Weapon.OriginalString,
		Killer:               k.Killer.SteamID64,
		Victim:               k.Victim.SteamID64,
		Assister:             assister,
		PenetratedObjects:    k.PenetratedObjects,
		IsHeadshot:           k.IsHeadshot,
	}
}

// func NewBsonMatch(m Match) BsonMatch {

// 	rounds := []BsonRound{}

// 	for _, r := range m.Rounds {
// 		rounds = append(rounds, NewBsonRound(*r))
// 	}

// 	return BsonMatch{

// 		ObjectID:   primitive.NewObjectID(),
// 		UploadTime: m.UploadTime,
// 		Map:        m.Map,
// 		Rounds:     rounds,
// 		// PlayerHurts []BsonPlayerHurt `bson:"playerhurts"`
// 	}
// }