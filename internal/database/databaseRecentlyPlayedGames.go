package database

import (
	"database/sql"

	log "github.com/sirupsen/logrus"

	"github.com/pinpox/megaclan3000/internal/steamclient"
)

// GetRecentlyPlayedGames returns a RecentlyPlayedGames object by fetching the values from
// the database using a prepared statement.
func (ds *DataStorage) GetRecentlyPlayedGames(steamID string) (steamclient.RecentlyPlayedGames, error) {
	rpg := steamclient.RecentlyPlayedGames{}
	var err error

	if rows, err := ds.statements["select_recently_played"].Query(steamID); err == nil {

		for rows.Next() {
			rows.Scan(
				&rpg.SteamID,
				&rpg.Appid,
				&rpg.ImgIconURL,
				&rpg.ImgLogoURL,
				&rpg.Name,
				&rpg.Playtime2Weeks,
				&rpg.PlaytimeForever,
				&rpg.PlaytimeLinuxForever,
				&rpg.PlaytimeMacForever,
				&rpg.PlaytimeWindowsForever,
			)
		}
	}
	return rpg, err
}

// UpdateRecentlyPlayedGames receives a RecentlyPlayedGames from a PlayerInfo
// object and updates the databaes for this steamID
func (ds *DataStorage) UpdateRecentlyPlayedGames(rpg steamclient.RecentlyPlayedGames) error {
	var result sql.Result
	var err error

	if result, err = ds.statements["update_recently_played"].Exec(
		rpg.SteamID,
		rpg.Appid,
		rpg.ImgIconURL,
		rpg.ImgLogoURL,
		rpg.Name,
		rpg.Playtime2Weeks,
		rpg.PlaytimeForever,
		rpg.PlaytimeLinuxForever,
		rpg.PlaytimeMacForever,
		rpg.PlaytimeWindowsForever,
	); err != nil {
		return err
	}

	if rows, err := result.RowsAffected(); err == nil {
		log.Debugf("Added %v to recently_played table. %v rows affected", rpg.SteamID, rows)
		return nil
	}
	return err
}
