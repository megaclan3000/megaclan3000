package database

import (
	"github.com/pinpox/megaclan3000/internal/steamclient"
)

// UpdatePlayerSummary receives a PlayerSummary and updates the database entry
// for that steamID
func (ds *DataStorage) UpdatePlayerSummary(ps steamclient.PlayerSummary) error {

	if _, err := ds.dbm.Delete(&ps); err != nil {
		return err
	}

	return ds.dbm.Insert(&ps)
}
