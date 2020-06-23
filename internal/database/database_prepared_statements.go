package database

import ()

func (ds *DataStorage) getSelectPreparedstatements() error {
	// Prepare all statements
	var err error

	// - query player_summary for player

	// - query player_extra for player

	// Get latest timestamp for steamID
	if ds.statements["select_player_history_latest_time"], err = ds.db.Prepare(`
			SELECT time FROM player_history
			WHERE steamid = '?'
			ORDER BY time DESC
			LIMIT 1`); err != nil {
		return err
	}

	// Other statements

	// Get all steamIDs known to player_stats table
	if ds.statements["select_all_player_ids"], err = ds.db.Prepare(`
			SELECT steamid FROM player_stats`); err != nil {
		return err
	}
	return nil
}
