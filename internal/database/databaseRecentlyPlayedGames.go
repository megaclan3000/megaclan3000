package database

import (
	"github.com/pinpox/megaclan3000/internal/steamclient"
)

// UpdateRecentlyPlayedGames receives a RecentlyPlayedGames from a PlayerInfo
// object and updates the databaes for this steamID
func (ds *DataStorage) UpdateRecentlyPlayedGames(rpg steamclient.RecentlyPlayedGames) error {

	_, err := ds.db.NamedExec(
		`INSERT OR REPLACE INTO recently_played (
			steamid,
            appid,
            img_icon_url,
            img_logo_url,
            name,
            playtime_2_weeks,
            playtime_forever,
            playtime_linux_forever,
            playtime_mac_forever,
            playtime_windows_forever)
		VALUES (
			:steamid,
            :appid,
            :img_icon_url,
            :img_logo_url,
            :name,
            :playtime_2_weeks,
            :playtime_forever,
            :playtime_linux_forever,
            :playtime_mac_forever,
            :playtime_windows_forever)
	`, &rpg)

	return err
}
