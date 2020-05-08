package database

import (
	"database/sql"

	log "github.com/Sirupsen/logrus"

	"github.com/pinpox/megaclan3000/internal/steamclient"
)

// GetPlayerSummary returns a PlayerSummary object by fetching the values from
// the database using a prepared statement.
func (ds *DataStorage) GetPlayerSummary(steamID string) (steamclient.PlayerSummary, error) {

	ps := steamclient.PlayerSummary{}
	var err error

	if rows, err := ds.statements["select_player_summary"].Query(steamID); err == nil {
		for rows.Next() {
			rows.Scan(
				&ps.SteamID,
				&ps.Communityvisibilitystate,
				&ps.Profilestate,
				&ps.Personaname,
				&ps.Profileurl,
				&ps.Avatar,
				&ps.Avatarmedium,
				&ps.Avatarfull,
				&ps.Lastlogoff,
				&ps.Personastate,
				&ps.Primaryclanid,
				&ps.Timecreated,
			)
		}
	}
	return ps, err
}

// UpdatePlayerSummary receives a PlayerSummary and updates the database entry
// for that steamID
func (ds *DataStorage) UpdatePlayerSummary(ps steamclient.PlayerSummary) error {

	var result sql.Result
	var err error

	if result, err = ds.statements["update_player_summary"].Exec(
		ps.SteamID,
		ps.Communityvisibilitystate,
		ps.Profilestate,
		ps.Personaname,
		ps.Profileurl,
		ps.Avatar,
		ps.Avatarmedium,
		ps.Avatarfull,
		ps.Lastlogoff,
		ps.Personastate,
		ps.Primaryclanid,
		ps.Timecreated,
	); err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	log.Println("Rows affected:", rows)
	log.Println("Added", ps.SteamID, ps.Personaname, "to player_summary table")
	return err
}
