package database

import (
	"github.com/pinpox/megaclan3000/internal/steamclient"
)

// GetPlayerHistory returns a PlayerHistory object by fetching the values from
// the database using a prepared statement.
func (ds *DataStorage) GetPlayerHistory(steamID string) (steamclient.PlayerHistory, error) {
	ph := steamclient.PlayerHistory{}
	var err error

	if rows, err := ds.statements["select_player_history"].Query(steamID); err == nil {
		for rows.Next() {
			rows.Scan(
				&ph.SteamID,
				&ph.Time,
				&ph.TotalKills,
			)
		}
	}
	return ph, err
}
