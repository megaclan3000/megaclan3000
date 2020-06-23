package database

import (
	"database/sql"
	"github.com/jmoiron/modl"
	log "github.com/sirupsen/logrus"
	"io/ioutil"

	// Use sqlite backend
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pinpox/megaclan3000/internal/steamclient"
)

// DataStorage is the main interface to the saved data. It provides methods for
// retrieval as well as methods to ingress new data from the API or update
// existing values
type DataStorage struct {
	db         *sqlx.DB
	statements map[string]*sql.Stmt
}

// GetPlayerInfoBySteamID returns a PlayerInfo from a steamID. It will try to
// get the needed values from the database and return an error if steamID
// cannot be found in it.
func (ds *DataStorage) GetPlayerInfoBySteamID(steamID string) (steamclient.PlayerInfo, error) {

	info := steamclient.PlayerInfo{}
	var err error

	if err = ds.db.Get(&info.PlayerSummary, "SELECT * FROM player_summary WHERE steamid=? LIMIT 1", steamID); err != nil {
		log.Warn("Error retrieving player_summary for steamID:", steamID)
	}

	if err = ds.db.Get(&info.RecentlyPlayedGames, "SELECT * FROM recently_played WHERE steamid=? LIMIT 1", steamID); err != nil {
		log.Warn("Error retrieving recently_played for steamID:", steamID)
	}

	if err = ds.db.Get(&info.UserStatsForGame.Extra, "SELECT * FROM player_extra WHERE steamid=? LIMIT 1", steamID); err != nil {
		log.Warn("Error retrieving player_extra for steamID:", steamID)
	}

	if err = ds.db.Get(&info.UserStatsForGame.Stats, "SELECT * FROM player_stats WHERE steamid=? LIMIT 1", steamID); err != nil {
		log.Warn("Error retrieving player_stats for steamID:", steamID)
	}

	entries := []steamclient.PlayerHistoryEntry{}

	if err = ds.db.Select(&entries, "SELECT * FROM player_history WHERE steamid=? ORDER BY time LIMIT 10", steamID); err != nil {
		log.Warn(err)
		log.Warn("Error retrieving player_history for steamID:", steamID)
	}

	info.PlayerHistory.SteamID = steamID
	info.PlayerHistory.Data = entries

	info.UserStatsForGame.SteamID = steamID

	return info, nil
}

// NewDataStorage creates a new DataStorage for a given sqlite database filepath
func NewDataStorage(pathStorage, pathSchema string) (*DataStorage, error) {
	var err error

	// Initialize database
	storage := new(DataStorage)
	storage.statements = make(map[string]*sql.Stmt)

	// Connect to database
	log.Debugf("Reading %v", pathStorage)
	if storage.db, err = sqlx.Open("sqlite3", pathStorage); err != nil {
		log.Fatal("Failed to open sqlite file", err)
	}

	// Read and execute schema from schema.sql
	schema, err := ioutil.ReadFile(pathSchema)
	if err != nil {
		log.Fatal(err)
	}

	storage.db.MustExec(string(schema))

	return storage, nil
}

// GetAllPlayers returns a PlayerInfo object for all players known to the
// database
func (ds *DataStorage) GetAllPlayers() ([]steamclient.PlayerInfo, error) {

	var players []steamclient.PlayerInfo
	var Ids []string

	if err := ds.db.Select(&Ids, "SELECT steamid FROM player_stats"); err == nil {
		for _, v := range Ids {
			log.Debugf("Got ID from database: %v", v)
			if pi, err := ds.GetPlayerInfoBySteamID(v); err == nil {
				players = append(players, pi)
			} else {
				log.Fatal(err)
			}
		}
	}

	return players, nil
}

// UpdatePlayerInfo receives a PlayerInfo object and updates the database entry
// for it's steamID
func (ds *DataStorage) UpdatePlayerInfo(pi steamclient.PlayerInfo) error {
	var err error

	dbm := modl.NewDbMap(ds.db.DB, modl.SqliteDialect{})
	dbm.AddTableWithName(steamclient.PlayerSummary{}, "player_summary").SetKeys(false, "steamid")
	dbm.AddTableWithName(steamclient.RecentlyPlayedGames{}, "recently_played").SetKeys(false, "steamid")
	dbm.AddTableWithName(steamclient.GameStats{}, "player_stats").SetKeys(false, "steamid")
	dbm.AddTableWithName(steamclient.GameExtras{}, "player_extra").SetKeys(false, "steamid")

	if err := dbm.CreateTablesIfNotExists(); err != nil {
		log.Warn("Database not creatable: ", err)
		return err
	}

	// PlayerSummary
	dbm.Insert(&pi.PlayerSummary)

	//RecentlyPlayedGames
	dbm.Insert(&pi.RecentlyPlayedGames)

	//UserStatsForGame
	dbm.Delete(&pi.UserStatsForGame.Stats)
	dbm.Delete(&pi.UserStatsForGame.Extra)
	if err != nil {
		panic(err)
	}
	err = dbm.Insert(&pi.UserStatsForGame.Stats)

	if err != nil {
		panic(err)
	}

	err = dbm.Insert(&pi.UserStatsForGame.Extra)

	if err != nil {
		panic(err)
	}

	return nil
}
