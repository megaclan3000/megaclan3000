package database

import (
	// "database/sql"
	"github.com/jmoiron/modl"

	log "github.com/sirupsen/logrus"

	"github.com/pinpox/megaclan3000/internal/steamclient"
)

// UpdatePlayerSummary receives a PlayerSummary and updates the database entry
// for that steamID
func (ds *DataStorage) UpdatePlayerSummary(ps steamclient.PlayerSummary) error {

	dbm := modl.NewDbMap(ds.db.DB, modl.SqliteDialect{})
	dbm.AddTableWithName(steamclient.PlayerSummary{}, "player_summary").SetKeys(false, "steamid")

	if err := dbm.CreateTablesIfNotExists(); err != nil {
		log.Warn("Database not creatable: ", err)
		return err
	}

	dbm.Insert(&ps)
	return nil
}
