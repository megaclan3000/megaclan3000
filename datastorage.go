package main

import (
	"database/sql"
	"errors"
	"path/filepath"
	"strconv"
	"time"

	_ "github.com/lib/pq"

	"github.com/megaclan3000/megaclan3000/internal/demoparser"
	"github.com/megaclan3000/megaclan3000/internal/steamclient"
	log "github.com/sirupsen/logrus"
)

// DataStorage is the interface to get and set data retrieved from the steam
// API. It holds the data as in-memory cache to avoid having to pull the data
// when a request is made for better response time
type DataStorage struct {
	Players []steamclient.PlayerInfo
	DB      *sql.DB
}

// UpdateData fetches new data from the steam API for information about
// players. Players that have not played in the last two weeks may report
// incorrect values
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

// NewDataStorage constructs a new interface to the data used throughout the
// application
func NewDataStorage() *DataStorage {

	db, err := sql.Open("postgres", "user=postgres password=megaclan dbname=megadb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	log.Debug("Dropping match table")
	init := "DROP TABLE IF EXISTS matches;"
	_, err = db.Exec(init)
	if err != nil {
		log.Fatal(err)
	}

	schema := `
	CREATE TABLE matches(
		id SERIAL PRIMARY KEY,
		match JSONB
	);
	`

	log.Debug("Creating matches table")
	_, err = db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}

	ds := &DataStorage{
		Players: steamClient.GetPlayers(),
		DB:      db,
	}

	//Find all matches form import folder
	files, err := filepath.Glob("demo-import/*.dem")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		log.Info("Found match demo: ", f)

		demoInfoFromDem, err := demoparser.GetMatchInfo(f, steamClient)
		if err != nil {
			panic(err)
		}

		log.Info("Parsed Match with ID: ", demoInfoFromDem.MatchID)

		err = ds.upload(demoInfoFromDem)
		if err != nil {
			log.Warning("Not uploading ", f, err)
		}
	}

	go ds.UpdateData()
	return ds
}

func (ds *DataStorage) upload(match demoparser.InfoStruct) error {

	// Check if the match is a valid megaclan3000 match
	if !match.MatchValid {
		return errors.New("Match not valid")
	}

	log.Debug("Checking if match ID is already present: ", match.MatchID)

	// Check if this match is already present
	rows, err := ds.DB.Query("SELECT match -> 'match_id' FROM matches where match ->> 'match_id' = $1;", match.MatchID)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		return errors.New("Match ID already present")
	}

	// The database driver will call the Value() method and and marshall the
	// attrs struct to JSON before the INSERT.
	_, err = ds.DB.Exec("INSERT INTO matches (match) VALUES($1)", match)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Inserted match with ID: ", match.MatchID)

	return nil
}

// GetMatchByID fetches a match from the database and creates a Infostruct form it
func (ds *DataStorage) GetMatchByID(id string) (demoparser.InfoStruct, error) {

	row := ds.DB.QueryRow("SELECT match FROM matches WHERE match ->> 'match_id' = $1;", id)
	var r demoparser.InfoStruct
	err := row.Scan(&r)
	return r, err
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

// GetMatches returns a list of all matches in the database
func (ds DataStorage) GetMatches() interface{} {

	type Matchplayer struct {
		PlayerName string `json:"player_name"`
		Avatar     string `json:"avatar"`
	}
	type data struct {
		MapName     string        `json:"map"`         // Name of the map
		ScoreClan   int           `json:"score_clan"`  // Points clan
		ScoreEnemy  int           `json:"score_enemy"` // Points enemy
		Time        time.Time     `json:"time"`        // Time it was played/uploaded
		Result      int           `json:"result"`      // Resunt: 1=won, 0=draw, -1=lost
		MatchID     string        `json:"matchid"`     // ID of the match, for links
		ClanPlayers []Matchplayer `json:"clan_players"`
	}
	var matches []data

	rows, err := ds.DB.Query("SELECT match FROM matches;")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var m demoparser.InfoStruct
		err := rows.Scan(&m)
		if err != nil {
			log.Fatal(err)
		}
		var result int

		if m.General.ScoreClan > m.General.ScoreEnemy {
			result = 1
		} else if m.General.ScoreClan < m.General.ScoreEnemy {
			result = -1
		} else {
			result = 0
		}

		var players []Matchplayer

		for _, v := range m.Players.Clan() {
			//TODO isclanmember reports all players form the  clan's team as
			//members. There should be a differentiation between actual members
			//and players that just where in our team
			if v.IsClanMember && v.Clantag == "megaclan3000" {
				players = append(players, Matchplayer{
					PlayerName: v.Name,
					Avatar:     v.AvatarURL,
				})
			}
		}

		matches = append(matches,
			data{
				MapName:     m.General.MapName,
				ScoreClan:   m.General.ScoreClan,
				ScoreEnemy:  m.General.ScoreEnemy,
				Time:        m.General.UploadTime,
				Result:      result,
				MatchID:     m.MatchID,
				ClanPlayers: players,
			})
	}

	return matches
}

// GetPlayers returns a list of all players in the db
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

// GetAwards returns a list of all award statuses in the db
func (ds DataStorage) GetAwards() interface{} {

	type Award struct {
		PlayerName  string    `json:"player_name"`
		Avatar      string    `json:"avatar"`
		SteamID64   string    `json:"steamid"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Time        time.Time `json:"time"`
	}

	ret := []Award{
		{
			PlayerName:  "randolf",
			Avatar:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
			SteamID64:   "111",
			Title:       "Robin Hood",
			Description: "Most donated weapons",
			Time:        time.Now(),
		},
		{
			PlayerName:  "randolf",
			Avatar:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
			SteamID64:   "111",
			Title:       "Mayfly",
			Description: "Fastest Death in Round",
			Time:        time.Now(),
		},
		{
			PlayerName:  "randolf",
			Avatar:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
			SteamID64:   "111",
			Title:       "Ground-Aimer",
			Description: "Most Shots to Legs",
			Time:        time.Now(),
		},
		{
			PlayerName:  "randolf",
			Avatar:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
			SteamID64:   "111",
			Title:       "Quick-Reloader",
			Description: "Most times reloaded",
			Time:        time.Now(),
		},
		{
			PlayerName:  "randolf",
			Avatar:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
			SteamID64:   "111",
			Title:       "BOT-Instructor",
			Description: "Most BOTs taken  over",
			Time:        time.Now(),
		},
		{
			PlayerName:  "randolf",
			Avatar:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
			SteamID64:   "111",
			Title:       "Talkshow host",
			Description: "Most words written in chat",
			Time:        time.Now(),
		},
		{
			PlayerName:  "randolf",
			Avatar:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
			SteamID64:   "111",
			Title:       "Rowdy",
			Description: "Most decoy-grenades thrown",
			Time:        time.Now(),
		},
		{
			PlayerName:  "randolf",
			Avatar:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
			SteamID64:   "111",
			Title:       "Pacifist",
			Description: "Most rounds without frags",
			Time:        time.Now(),
		},
		{
			PlayerName:  "randolf",
			Avatar:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
			SteamID64:   "111",
			Title:       "Partystarter",
			Description: "Most Entry-Frags",
			Time:        time.Now(),
		},
		{
			PlayerName:  "randolf",
			Avatar:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
			SteamID64:   "111",
			Title:       "Dont be looser, buy Defuser",
			Description: "Most defuse kits bought",
			Time:        time.Now(),
		},
		{
			PlayerName:  "randolf",
			Avatar:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
			SteamID64:   "111",
			Title:       "Jeff Bezos",
			Description: "Most money saved on last Round",
			Time:        time.Now(),
		},
		{
			PlayerName:  "randolf",
			Avatar:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
			SteamID64:   "111",
			Title:       "Economist",
			Description: "Least money spend",
			Time:        time.Now(),
		},
		{
			PlayerName:  "randolf",
			Avatar:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
			SteamID64:   "111",
			Title:       "",
			Description: "",
			Time:        time.Now(),
		},
		{
			PlayerName:  "randolf",
			Avatar:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
			SteamID64:   "111",
			Title:       "Freelancer",
			Description: "Most Aces",
			Time:        time.Now(),
		},
		{
			PlayerName:  "randolf",
			Avatar:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
			SteamID64:   "111",
			Title:       "Kangoroo",
			Description: "Most jumps in a single Match",
			Time:        time.Now(),
		},
		{
			PlayerName:  "randolf",
			Avatar:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
			SteamID64:   "111",
			Title:       "Forrest Gump",
			Description: "Most steps in a single Match",
			Time:        time.Now(),
		},
		{
			PlayerName:  "randolf",
			Avatar:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
			SteamID64:   "111",
			Title:       "Bomb-Protector",
			Description: "Most deaths to bomb Explosion",
			Time:        time.Now(),
		},
		{
			PlayerName:  "randolf",
			Avatar:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
			SteamID64:   "111",
			Title:       "Fake-Defuser",
			Description: "Most faked bomb-defuses",
			Time:        time.Now(),
		},
		{
			PlayerName:  "randolf",
			Avatar:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
			SteamID64:   "111",
			Title:       "Photographer",
			Description: "Most teammates blinded",
			Time:        time.Now(),
		},
		{
			PlayerName:  "randolf",
			Avatar:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
			SteamID64:   "111",
			Title:       "Chicken-Chef",
			Description: "Most chickens killed",
			Time:        time.Now(),
		},
		{
			PlayerName:  "randolf",
			Avatar:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
			SteamID64:   "111",
			Title:       "Hip-Shooter",
			Description: "Most no-scope kills with AWP",
			Time:        time.Now(),
		},
		{
			PlayerName:  "randolf",
			Avatar:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
			SteamID64:   "111",
			Title:       "Teamkiller",
			Description: "Most Damage to Teammates",
			Time:        time.Now(),
		},
		{
			PlayerName:  "randolf",
			Avatar:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
			SteamID64:   "111",
			Title:       "Electrician",
			Description: "Most enemies tasered",
			Time:        time.Now(),
		},
		{
			PlayerName:  "randolf",
			Avatar:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
			SteamID64:   "111",
			Title:       "Master Slicer",
			Description: "Most Knive Kills",
			Time:        time.Now(),
		},
		{
			PlayerName:  "randolf",
			Avatar:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
			SteamID64:   "111",
			Title:       "Warmup King",
			Description: "Most Kills during Warmup",
			Time:        time.Now(),
		},
		{
			PlayerName:  "randolf",
			Avatar:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/06/06e1eec83d05fd0823728381fcbe27c0d8318510_full.jpg",
			SteamID64:   "111",
			Title:       "Nade Inspector",
			Description: "Most Deaths while not holding a Weapon",
			Time:        time.Now(),
		},
	}
	return ret
}

// GetUpdates returns all updates form the db
func (ds DataStorage) GetUpdates() interface{} {

	type UpdateType int

	const (
	// Award UpdateType = 0
	// Rank                 = 1
	// MatchWon  = 2
	// MatchLost = 3
	// MatchDraw = 4
	)

	type Update struct {
		Time  time.Time
		Type  UpdateType
		Text1 string
		Text2 string
	}

	ret := []Update{
		//TODO add real data
		{
			Type:  0,
			Time:  time.Now(),
			Text1: "randolf received award",
			Text2: "Warmup-Killer",
		},
		{
			Type:  1,
			Time:  time.Now(),
			Text1: "salatkopf new rank",
			Text2: "4",
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
