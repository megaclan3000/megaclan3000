package database

import (
	"database/sql"

	log "github.com/sirupsen/logrus"

	"github.com/pinpox/megaclan3000/internal/steamclient"
)

// UpdatePlayerSummary receives a PlayerSummary and updates the database entry
// for that steamID
func (ds *DataStorage) UpdatePlayerSummary(ps steamclient.PlayerSummary) error {

	var result sql.Result
	var err error

	if result, err = ds.statements["update_player_summary"].Exec(
		ps.SteamID,
		ps.Avatar,
		ps.Avatarfull,
		ps.Avatarmedium,
		ps.Cityid,
		ps.Commentpermission,
		ps.Communityvisibilitystate,
		ps.Gameextrainfo,
		ps.Gameid,
		ps.Gameserverip,
		ps.Lastlogoff,
		ps.Loccityid,
		ps.Loccountrycode,
		ps.Locstatecode,
		ps.Personaname,
		ps.Personastate,
		ps.Primaryclanid,
		ps.Profilestate,
		ps.Profileurl,
		ps.Realname,
		ps.Timecreated,
	); err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	log.Debugf("Added %v (%v) to player_summary table. %v rows affected", ps.SteamID, ps.Personaname, rows)
	return err
}
