package database

import (
	"github.com/jmoiron/modl"

	"github.com/sirupsen/logrus"

	"github.com/pinpox/megaclan3000/internal/steamclient"
)

// UpdateUserStatsForGame receives a UserStatsForGame object and updates the
// corresponding entry in the database for the steamID
func (ds *DataStorage) UpdateUserStatsForGame(stats steamclient.UserStatsForGame) error {

	dbm := modl.NewDbMap(ds.db.DB, modl.SqliteDialect{})
	dbm.AddTableWithName(steamclient.GameStats{}, "player_stats").SetKeys(false, "steamid")
	dbm.AddTableWithName(steamclient.GameExtras{}, "player_extra").SetKeys(false, "steamid")

	if err := dbm.CreateTablesIfNotExists(); err != nil {
		logrus.Fatal("Database not creatable: ", err)
		return err
	}

	if _, err := dbm.Delete(&stats.Stats); err != nil {
		return err
	}

	if _, err := dbm.Delete(&stats.Extra); err != nil {
		return err
	}
	if err := dbm.Insert(&stats.Stats); err != nil {
		return err
	}

	return dbm.Insert(&stats.Extra)

}
