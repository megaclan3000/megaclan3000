package main

import (
	"errors"
	"strconv"
	"time"

	// "github.com/megaclan3000/megaclan3000/internal/demoparser"
	"github.com/megaclan3000/megaclan3000/internal/demoparser"
	"github.com/megaclan3000/megaclan3000/internal/steamclient"
	log "github.com/sirupsen/logrus"
)

// DataStorage is the interface to get and set data retrieved from the steam
// API. It holds the data as in-memory cache to avoid having to pull the data
// when a request is made for better response time
type DataStorage struct {
	Players []steamclient.PlayerInfo
}

func (ds *DataStorage) UpdateData() {

	// Get PlayerInfo for all players periodically and store/cache in
	// memory so we don't have to wait when retrieving it in the fronend
	for {
		log.Debug("Updating player information")
		ds.Players = steamClient.GetPlayers()

		// Sleep for a predefined duration (in minutes)
		time.Sleep(time.Duration(steamClient.Config.UpdateInterval) * time.Minute)
	}
}

func NewDataStorage() *DataStorage {
	ds := &DataStorage{Players: steamClient.GetPlayers()}
	go ds.UpdateData()
	return ds
}

func (ds *DataStorage) Upload(match demoparser.InfoStruct) error {
	return nil
}

func (ds *DataStorage) GetMatchByID(id string) (demoparser.InfoStruct, error) {

	demoInfoFromDem, err := demoparser.GetMatchInfo("1", steamClient)
	if err != nil {
		panic(err)
	}

	return demoInfoFromDem, nil
}

// GetPlayerInfoBySteamID returns the PlayerInfo object for a given steamID
func (ds DataStorage) GetPlayerInfoBySteamID(steamID uint64) (steamclient.PlayerInfo, error) {

	for _, v := range ds.Players {
		if v.PlayerSummary.SteamID == steamID {
			return v, nil
		}
	}
	return steamclient.PlayerInfo{}, errors.New("Player not found")
}

func (ds DataStorage) GetMatches() interface{} {
	//TODO implement real data
	ret := []struct {
		MapName    string    `json:"map"`         // Name of the map
		ScoreClan  int       `json:"score_clan"`  // Points clan
		ScoreEnemy int       `json:"score_enemy"` // Points enemy
		Time       time.Time `json:"time"`        // Time it was played/uploaded
		Result     int       `json:"result"`      // Resunt: 1=won, 0=draw, -1=lost
		MatchID    string    `json:"matchid"`     // ID of the match, for links
	}{
		{
			MapName:    "de_dust2",
			ScoreClan:  16,
			ScoreEnemy: 4,
			Time:       time.Now(),
			Result:     1,
			MatchID:    "1",
		},
		{
			MapName:    "de_mirage",
			ScoreClan:  16,
			ScoreEnemy: 5,
			Time:       time.Now(),
			Result:     1,
			MatchID:    "1",
		},
		{
			MapName:    "de_inferno",
			ScoreClan:  5,
			ScoreEnemy: 16,
			Time:       time.Now(),
			Result:     -1,
			MatchID:    "1",
		},
		{
			MapName:    "de_cache",
			ScoreClan:  15,
			ScoreEnemy: 15,
			Time:       time.Now(),
			Result:     0,
			MatchID:    "1",
		},
	}

	return ret
}

func (ds DataStorage) GetPlayers() interface{} {

	type Player struct {
		PlayerName string `json:"player_name"`
		SteamID64  string `json:"steamid"`
		Avatar     string `json:"avatar"`
		Matches    int    `json:"matches"`

		Kills  int `json:"kills"`
		Deaths int `json:"deaths"`
		Hits   int `json:"hits"`
		Shots  int `json:"shots"`

		Hours  string `json:"hours"`
		Wins   int    `json:"wins"`
		Points int    `json:"points"`
		Status string `json:"status"`
	}

	var ret []Player

	for _, v := range ds.Players {
		ret = append(ret, Player{
			PlayerName: v.PlayerSummary.Personaname,
			SteamID64:  strconv.FormatUint(v.UserStatsForGame.SteamID, 10),
			Avatar:     v.PlayerSummary.Avatarfull,
			Matches:    0, //TODO

			Kills:  v.UserStatsForGame.Stats.TotalKills,
			Deaths: v.UserStatsForGame.Stats.TotalDeaths,
			Hits:   v.UserStatsForGame.Stats.TotalShotsHit,
			Shots:  v.UserStatsForGame.Stats.TotalShotsFired,

			Hours:  v.RecentlyPlayedGames.PlaytimeForever,
			Wins:   v.UserStatsForGame.Stats.TotalMatchesWon,
			Points: 0, //TODO
			Status: v.PlayerSummary.Personastate,
		})
	}
	return ret
}

func (ds DataStorage) GetUpdates() interface{} {

	type UpdateType int

	const (
		Award     UpdateType = 0
		Rank                 = 1
		MatchWon             = 2
		MatchLost            = 3
		MatchDraw            = 4
	)

	type Update struct {
		Time  time.Time
		Type  UpdateType
		Text1 string
		Text2 string
	}

	ret := []Update{
		{
			Type:  0,
			Time:  time.Now(),
			Text1: "randolf received award",
			Text2: "Warmup-Killer",
		},
		{
			Type:  1,
			Time:  time.Now(),
			Text1: "salatkopf ranked up",
			Text2: "4 to 5",
		},
		{
			Type:  2,
			Time:  time.Now(),
			Text1: "Clan won match",
			Text2: "de_dust2",
		},
		{
			Type:  3,
			Time:  time.Now(),
			Text1: "Clan lost match",
			Text2: "de_mirage",
		},
		{
			Type:  4,
			Time:  time.Now(),
			Text1: "Clan draw match",
			Text2: "de_inferno",
		},
	}

	return ret
}
