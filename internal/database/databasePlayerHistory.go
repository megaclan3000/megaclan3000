package database

import (
	"github.com/jmoiron/modl"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/pinpox/megaclan3000/internal/steamclient"
	log "github.com/sirupsen/logrus"
)

// GetPlayerHistoryLatestTime returns the time of the last entry in the
// player_history table for a specified ID. This is used to check whether a new
// entry should be added for the current values
func (ds *DataStorage) GetPlayerHistoryLatestTime(steamID string) time.Time {

	var err error
	var updateTime time.Time

	dbm := modl.NewDbMap(ds.db.DB, modl.SqliteDialect{})

	if err = dbm.SelectOne(&updateTime,
		`SELECT time FROM player_history
			WHERE steamid = ?
			ORDER BY time DESC
			LIMIT 1`, steamID); err != nil {
		log.Warn(err)
	}

	return updateTime
}

// UpdatePlayerHistory takes a PlayerInfo object and saves an entry to the
// player_history table with the current time and the values from the
// PlayerInfo object
func (ds *DataStorage) UpdatePlayerHistory(phe steamclient.PlayerHistoryEntry) error {

	dbm := modl.NewDbMap(ds.db.DB, modl.SqliteDialect{})
	dbm.AddTableWithName(steamclient.PlayerHistoryEntry{}, "player_history")

	if err := dbm.CreateTablesIfNotExists(); err != nil {
		logrus.Fatal("Database not creatable: ", err)
		return err
	}

	return dbm.Insert(&phe)
}
