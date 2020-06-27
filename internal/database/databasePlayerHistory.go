package database

import (
	"time"

	"github.com/megaclan3000/megaclan3000/internal/steamclient"
	log "github.com/sirupsen/logrus"
)

// GetPlayerHistoryLatestTime returns the time of the last entry in the
// player_history table for a specified ID. This is used to check whether a new
// entry should be added for the current values
func (ds *DataStorage) GetPlayerHistoryLatestTime(steamID string) time.Time {

	var err error
	var updateTime time.Time

	if err = ds.dbm.SelectOne(&updateTime,
		`SELECT time FROM player_history
			WHERE steamid = ?
			ORDER BY time DESC
			LIMIT 1`, steamID); err != nil {
		log.Warn("Could not get latest history time for id:", steamID)
		log.Warn(err)
	}

	return updateTime
}

// UpdatePlayerHistory takes a PlayerInfo object and saves an entry to the
// player_history table with the current time and the values from the
// PlayerInfo object
func (ds *DataStorage) UpdatePlayerHistory(phe steamclient.PlayerHistoryEntry) error {
	return ds.dbm.Insert(&phe)
}
