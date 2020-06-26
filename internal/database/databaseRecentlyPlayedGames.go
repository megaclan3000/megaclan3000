package database

import (
	"github.com/pinpox/megaclan3000/internal/steamclient"
	"github.com/sirupsen/logrus"
)

// UpdateRecentlyPlayedGames receives a RecentlyPlayedGames from a PlayerInfo
// object and updates the databaes for this steamID
func (ds *DataStorage) UpdateRecentlyPlayedGames(rpg steamclient.RecentlyPlayedGames) error {

	logrus.Debugf("Inserting id %v into recently_played", rpg.SteamID)
	if _, err := ds.dbm.Delete(&rpg); err != nil {
		return err
	}

	return ds.dbm.Insert(&rpg)
}
