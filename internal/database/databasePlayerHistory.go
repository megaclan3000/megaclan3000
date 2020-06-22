package database

import (
	"database/sql"
	"time"

	"github.com/pinpox/megaclan3000/internal/steamclient"
	log "github.com/sirupsen/logrus"
)

// GetPlayerHistoryLatestTime returns the time of the last entry in the
// player_history table for a specified ID. This is used to check whether a new
// entry should be added for the current values
func (ds *DataStorage) GetPlayerHistoryLatestTime(steamID string) (time.Time, error) {

	var updateTime time.Time
	var err error

	if rows, err := ds.statements["select_player_history_latest_time"].Query(steamID); err == nil {
		for rows.Next() {
			rows.Scan(
				updateTime,
			)
		}
	}
	return updateTime, err
}

// UpdatePlayerHistory takes a PlayerInfo object and saves an entry to the
// player_history table with the current time and the values from the
// PlayerInfo object
func (ds *DataStorage) UpdatePlayerHistory(pi steamclient.PlayerInfo) error {

	var result sql.Result
	var err error

	if result, err = ds.statements["insert_player_history"].Exec(
		pi.PlayerSummary.SteamID,
		pi.UserStatsForGame.Stats.TotalKills,
		pi.UserStatsForGame.Extra.TotalADR,
		pi.UserStatsForGame.Stats.TotalShotsHit,
		pi.UserStatsForGame.Stats.TotalShotsFired,
		pi.UserStatsForGame.Stats.TotalKillsHeadshot,
		pi.UserStatsForGame.Extra.TotalKD,
		pi.UserStatsForGame.Stats.LastMatchContributionScore,
		pi.UserStatsForGame.Stats.LastMatchDamage,
		pi.UserStatsForGame.Stats.LastMatchDeaths,
		pi.UserStatsForGame.Stats.LastMatchKills,
		pi.UserStatsForGame.Stats.LastMatchRounds,
		pi.UserStatsForGame.Extra.LastMatchKD,
		pi.UserStatsForGame.Extra.LastMatchADR,
		pi.UserStatsForGame.Extra.HitRatio,
		pi.RecentlyPlayedGames.Playtime2Weeks,
	); err != nil {
		return err
	}

	rows, err := result.RowsAffected()

	log.Debugf("Added entry for %v (%v) to player_history table. %v rows affected",
		pi.PlayerSummary.SteamID,
		pi.PlayerSummary.Personaname,
		rows)

	return err
}
