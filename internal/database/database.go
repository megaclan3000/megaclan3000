package database

import (
	"database/sql"

	log "github.com/sirupsen/logrus"

	// Use sqlite backend
	_ "github.com/mattn/go-sqlite3"
	"github.com/pinpox/megaclan3000/internal/steamclient"
)

// DataStorage is the main interface to the saved data. It provides methods for
// retrieval as well as methods to ingress new data from the API or update
// existing values
type DataStorage struct {
	db         *sql.DB
	statements map[string]*sql.Stmt
}

// GetPlayerInfoBySteamID returns a PlayerInfo from a steamID. It will try to
// get the needed values from the database and return an error if steamID
// cannot be found in it.
func (ds *DataStorage) GetPlayerInfoBySteamID(steamID string) (steamclient.PlayerInfo, error) {

	info := steamclient.PlayerInfo{}
	var err error

	if info.PlayerSummary, err = ds.GetPlayerSummary(steamID); err != nil {
		return info, err
	}

	if info.RecentlyPlayedGames, err = ds.GetRecentlyPlayedGames(steamID); err != nil {
		return info, err
	}

	if info.UserStatsForGame, err = ds.GetUserStatsForGame(steamID); err != nil {
		return info, err
	}

	if info.PlayerHistory, err = ds.GetPlayerHistory(steamID); err != nil {
		return info, err
	}

	return info, nil
}

// NewDataStorage creates a new DataStorage for a given sqlite database filepath
func NewDataStorage(path string) (*DataStorage, error) {
	var err error

	// Initialize database
	storage := new(DataStorage)
	storage.statements = make(map[string]*sql.Stmt)

	log.Debugf("Reading %v", path)
	if storage.db, err = sql.Open("sqlite3", path); err != nil {
		log.Fatal("Failed to open sqlite file", err)
	}

	// Prepare CREATE statements
	if err = storage.getCreatePreparedstatements(); err != nil {
		log.Fatal("Failed to prepare CREATE statements", err)
	}

	// Create tables, if necessary
	if _, err = storage.statements["create_player_summary"].Exec(); err != nil {
		log.Fatal("Failed to create table player_summary", err)
	}

	if _, err = storage.statements["create_player_stats"].Exec(); err != nil {
		log.Fatal("Failed to create table player_stats", err)
	}

	if _, err = storage.statements["create_recently_played"].Exec(); err != nil {
		log.Fatal("Failed to create table recently_played", err)
	}

	if _, err = storage.statements["create_player_history"].Exec(); err != nil {
		log.Fatal("Failed to create table player_history", err)
	}

	if _, err = storage.statements["create_player_extra"].Exec(); err != nil {
		log.Fatal("Failed to create table player_extra", err)
	}

	// Prepare remaining statements
	if err = storage.getUpdatePreparedstatements(); err != nil {
		log.Fatal("Failed to prepare UPDATE statements", err)
	}

	if err = storage.getInsertPreparedstatements(); err != nil {
		log.Fatal("Failed to prepare INSERT statements", err)
	}

	if err = storage.getSelectPreparedstatements(); err != nil {
		log.Fatal("Failed to prepare SELECT statements", err)
	}

	return storage, nil
}

// GetAllPlayers returns a PlayerInfo object for all players known to the
// database
func (ds *DataStorage) GetAllPlayers() ([]steamclient.PlayerInfo, error) {
	var players []steamclient.PlayerInfo
	var rows *sql.Rows
	var err error

	if rows, err = ds.statements["select_all_player_ids"].Query(); err != nil {
		return players, err
	}

	var steamID string

	for rows.Next() {
		if err = rows.Scan(&steamID); err == nil {
			log.Debugf("Got ID from database: %v", steamID)
			if pi, err := ds.GetPlayerInfoBySteamID(steamID); err == nil {
				players = append(players, pi)
			} else {
				log.Fatal(err)
			}
		}
	}

	rows.Close() //good habit to close
	return players, nil
}

// UpdatePlayerInfo receives a PlayerInfo object and updates the database entry
// for it's steamID
func (ds *DataStorage) UpdatePlayerInfo(pi steamclient.PlayerInfo) error {
	var err error

	if err = ds.UpdatePlayerSummary(pi.PlayerSummary); err != nil {
		log.Fatalf("Error saving PlayerSummary for %v (%v)", pi.PlayerSummary.SteamID, pi.PlayerSummary.Personaname)
		return err
	}
	if err = ds.UpdateRecentlyPlayedGames(pi.RecentlyPlayedGames); err != nil {
		log.Fatalf("Error saving RecentlyPlayedGames for %v (%v)", pi.PlayerSummary.SteamID, pi.PlayerSummary.Personaname)
		return err
	}
	if err = ds.UpdateUserStatsForGame(pi.UserStatsForGame); err != nil {
		log.Fatalf("Error saving UserStatsForGame for %v (%v)", pi.PlayerSummary.SteamID, pi.PlayerSummary.Personaname)
		return err
	}

	return nil
}
