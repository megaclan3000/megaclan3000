package database

import (
	"database/sql"
	"time"

	"github.com/pinpox/megaclan3000/internal/steamclient"
	log "github.com/sirupsen/logrus"
)

// GetPlayerHistory returns a PlayerHistory object by fetching the values from
// the database using a prepared statement.
func (ds *DataStorage) GetPlayerHistory(steamID string) (steamclient.PlayerHistory, error) {
	ph := steamclient.PlayerHistory{}
	var err error

	var entry steamclient.PlayerHistoryEntry

	if rows, err := ds.statements["select_player_history"].Query(steamID); err == nil {
		for rows.Next() {
			rows.Scan(
				&ph.SteamID,
				&entry.Time,
				&entry.TotalKD,
			)
			ph.Data = append(ph.Data, entry)
		}
	}
	return ph, err
}

// GetPlayerHistoryLatestTime returns the time of the last entry in the
// player_history table for a specified ID. This is used to check whether a new
// entry should be added for the current values
func (ds *DataStorage) GetPlayerHistoryLatestTime(steamID string) (int, error) {

	var time int
	var err error

	if rows, err := ds.statements["select_player_history_latest_time"].Query(steamID); err == nil {
		for rows.Next() {
			rows.Scan(
				time,
			)
		}
	}
	return time, err
}

// UpdatePlayerHistory takes a PlayerInfo object and saves an entry to the
// player_history table with the current time and the values from the
// PlayerInfo object
func (ds *DataStorage) UpdatePlayerHistory(pi steamclient.PlayerInfo) error {

	var result sql.Result
	var err error

	if result, err = ds.statements["insert_player_history"].Exec(
		pi.PlayerSummary.SteamID,
		time.Now(),
		//TODO add other stats here, e.g. ADR
		pi.UserStatsForGame.Extra.TotalKD,
	); err != nil {
		return err
	}

	rows, err := result.RowsAffected()

	log.Debugf("Added entry for %v (%v) to player_history table. %v rows affected",
		pi.PlayerSummary.SteamID,
		pi.PlayerSummary.Personaname,
		rows)

	return err
}
