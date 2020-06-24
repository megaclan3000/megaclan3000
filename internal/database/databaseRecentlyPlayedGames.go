package database

import (
	"github.com/jmoiron/modl"
	"github.com/pinpox/megaclan3000/internal/steamclient"
	log "github.com/sirupsen/logrus"
)

// UpdateRecentlyPlayedGames receives a RecentlyPlayedGames from a PlayerInfo
// object and updates the databaes for this steamID
func (ds *DataStorage) UpdateRecentlyPlayedGames(rpg steamclient.RecentlyPlayedGames) error {

	dbm := modl.NewDbMap(ds.db.DB, modl.SqliteDialect{})
	dbm.AddTableWithName(steamclient.RecentlyPlayedGames{}, "recently_played").SetKeys(false, "steamid")

	if err := dbm.CreateTablesIfNotExists(); err != nil {
		log.Warn("Database not creatable: ", err)
		return err
	}

	if _, err := dbm.Delete(&rpg); err != nil {
		return err
	}

	return dbm.Insert(&rpg)

}
