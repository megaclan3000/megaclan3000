package database

import ()

func (ds *DataStorage) getInsertPreparedstatements() error {
	var err error
	ds.statements["insert_player_history"], err = ds.db.Prepare(
		`INSERT INTO player_history (
			steamid,
			time,
			total_kills,
			total_adr,
			total_shots_hit,
			total_shots_fired,
			total_kills_headshot,
			total_kd,
			last_match_contribution_score,
			last_match_damage,
			last_match_deaths,
			last_match_kills,
			last_match_rounds,
			last_match_kd,
			last_match_adr,
			hit_ratio,
			playtime_2_weeks)
		VALUES (?, datetime('now'), ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	return err
}

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
