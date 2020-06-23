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

	var err error
	dbm.Delete(&stats.Stats)
	dbm.Delete(&stats.Extra)
	if err != nil {
		panic(err)
	}
	err = dbm.Insert(&stats.Stats)

	if err != nil {
		panic(err)
	}

	err = dbm.Insert(&stats.Extra)

	if err != nil {
		panic(err)
	}
	return nil

}
