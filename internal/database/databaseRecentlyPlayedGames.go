package database

import (
	"database/sql"

	log "github.com/sirupsen/logrus"

	"github.com/pinpox/megaclan3000/internal/steamclient"
)

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
