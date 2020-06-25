package database

import (
	"github.com/sirupsen/logrus"

	"github.com/pinpox/megaclan3000/internal/steamclient"
)

// UpdateUserStatsForGame receives a UserStatsForGame object and updates the
// corresponding entry in the database for the steamID
func (ds *DataStorage) UpdateUserStatsForGame(stats steamclient.UserStatsForGame) error {

	if _, err := ds.dbm.Delete(&stats.Stats); err != nil {
		return err
	}

	if _, err := ds.dbm.Delete(&stats.Extra); err != nil {
		return err
	}

	logrus.Debugf("Inserting id %v into player_stats", stats.Stats.SteamID)
	if err := ds.dbm.Insert(&stats.Stats); err != nil {
		return err
	}

	logrus.Debugf("Inserting id %v into player_extra", stats.Extra.SteamID)
	return ds.dbm.Insert(&stats.Extra)

}
