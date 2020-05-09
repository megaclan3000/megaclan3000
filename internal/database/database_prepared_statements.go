package database

import (
	log "github.com/sirupsen/logrus"
)

func (ds *DataStorage) getCreatePreparedstatements() error {
	var err error

	ds.statements["create_player_summary"], err = ds.db.Prepare(
		`CREATE TABLE IF NOT EXISTS player_summary (
			steamid TEXT PRIMARY KEY,
			communityvisibilitystate TEXT,
			profilestate TEXT,
			personaname TEXT,
			profileurl TEXT,
			avatar TEXT,
			avatarmedium TEXT,
			avatarfull TEXT,
			lastlogoff TEXT,
			personastate TEXT,
			primaryclanid TEXT,
			timecreated TEXT)`)

	if err != nil {
		return err
	}

	ds.statements["create_player_stats"], err = ds.db.Prepare(
		`CREATE TABLE IF NOT EXISTS player_stats (
			steamid TEXT PRIMARY KEY,
			total_kills TEXT,
			total_deaths TEXT,
			total_time_played TEXT,
			total_planted_bombs TEXT,
			total_defused_bombs TEXT,
			total_wins TEXT,
			total_damage_done TEXT,
			total_money_earned TEXT,
			total_kills_knife TEXT,
			total_kills_hegrenade TEXT,
			total_kills_glock TEXT,
			total_kills_deagle TEXT,
			total_kills_elite TEXT,
			total_kills_fiveseven TEXT,
			total_kills_xm1014 TEXT,
			total_kills_mac10 TEXT,
			total_kills_ump45 TEXT,
			total_kills_p90 TEXT,
			total_kills_awp TEXT,
			total_kills_ak47 TEXT,
			total_kills_aug TEXT,
			total_kills_famas TEXT,
			total_kills_g3sg1 TEXT,
			total_kills_m249 TEXT,
			total_kills_headshot TEXT,
			total_kills_enemy_weapon TEXT,
			total_wins_pistolround TEXT,
			total_wins_map_cs_assault TEXT,
			total_wins_map_de_dust2 TEXT,
			total_wins_map_de_inferno TEXT,
			total_wins_map_de_train TEXT,
			total_weapons_donated TEXT,
			total_kills_enemy_blinded TEXT,
			total_kills_knife_fight TEXT,
			total_kills_against_zoomed_sniper TEXT,
			total_dominations TEXT,
			total_domination_overkills TEXT,
			total_revenges TEXT,
			total_shots_hit TEXT,
			total_shots_fired TEXT,
			total_rounds_played TEXT,
			total_shots_deagle TEXT,
			total_shots_glock TEXT,
			total_shots_elite TEXT,
			total_shots_fiveseven TEXT,
			total_shots_awp TEXT,
			total_shots_ak47 TEXT,
			total_shots_aug TEXT,
			total_shots_famas TEXT,
			total_shots_g3sg1 TEXT,
			total_shots_p90 TEXT,
			total_shots_mac10 TEXT,
			total_shots_ump45 TEXT,
			total_shots_xm1014 TEXT,
			total_shots_m249 TEXT,
			total_hits_deagle TEXT,
			total_hits_glock TEXT,
			total_hits_elite TEXT,
			total_hits_fiveseven TEXT,
			total_hits_awp TEXT,
			total_hits_ak47 TEXT,
			total_hits_aug TEXT,
			total_hits_famas TEXT,
			total_hits_g3sg1 TEXT,
			total_hits_p90 TEXT,
			total_hits_mac10 TEXT,
			total_hits_ump45 TEXT,
			total_hits_xm1014 TEXT,
			total_hits_m249 TEXT,
			total_rounds_map_cs_assault TEXT,
			total_rounds_map_de_dust2 TEXT,
			total_rounds_map_de_inferno TEXT,
			total_rounds_map_de_train TEXT,
			last_match_t_wins TEXT,
			last_match_ct_wins TEXT,
			last_match_wins TEXT,
			last_match_max_players TEXT,
			last_match_kills TEXT,
			last_match_deaths TEXT,
			last_match_mvps TEXT,
			last_match_favweapon_id TEXT,
			last_match_favweapon_shots TEXT,
			last_match_favweapon_hits TEXT,
			last_match_favweapon_kills TEXT,
			last_match_damage TEXT,
			last_match_money_spent TEXT,
			last_match_dominations TEXT,
			last_match_revenges TEXT,
			total_mvps TEXT,
			total_rounds_map_de_lake TEXT,
			total_rounds_map_de_safehouse TEXT,
			total_rounds_map_de_bank TEXT,
			total_TR_planted_bombs TEXT,
			total_gun_game_rounds_won TEXT,
			total_gun_game_rounds_played TEXT,
			total_wins_map_de_bank TEXT,
			total_wins_map_de_lake TEXT,
			total_matches_won_bank TEXT,
			total_matches_won TEXT,
			total_matches_played TEXT,
			total_gg_matches_won TEXT,
			total_gg_matches_played TEXT,
			total_progressive_matches_won TEXT,
			total_trbomb_matches_won TEXT,
			total_contribution_score TEXT,
			last_match_contribution_score TEXT,
			last_match_rounds TEXT,
			total_kills_hkp2000 TEXT,
			total_shots_hkp2000 TEXT,
			total_hits_hkp2000 TEXT,
			total_hits_p250 TEXT,
			total_kills_p250 TEXT,
			total_shots_p250 TEXT,
			total_kills_sg556 TEXT,
			total_shots_sg556 TEXT,
			total_hits_sg556 TEXT,
			total_hits_scar20 TEXT,
			total_kills_scar20 TEXT,
			total_shots_scar20 TEXT,
			total_shots_ssg08 TEXT,
			total_hits_ssg08 TEXT,
			total_kills_ssg08 TEXT,
			total_shots_mp7 TEXT,
			total_hits_mp7 TEXT,
			total_kills_mp7 TEXT,
			total_kills_mp9 TEXT,
			total_shots_mp9 TEXT,
			total_hits_mp9 TEXT,
			total_hits_nova TEXT,
			total_kills_nova TEXT,
			total_shots_nova TEXT,
			total_hits_negev TEXT,
			total_kills_negev TEXT,
			total_shots_negev TEXT,
			total_shots_sawedoff TEXT,
			total_hits_sawedoff TEXT,
			total_kills_sawedoff TEXT,
			total_shots_bizon TEXT,
			total_hits_bizon TEXT,
			total_kills_bizon TEXT,
			total_kills_tec9 TEXT,
			total_shots_tec9 TEXT,
			total_hits_tec9 TEXT,
			total_shots_mag7 TEXT,
			total_hits_mag7 TEXT,
			total_kills_mag7 TEXT,
			total_gun_game_contribution_score TEXT,
			last_match_gg_contribution_score TEXT,
			total_kills_m4a1 TEXT,
			total_kills_galilar TEXT,
			total_kills_molotov TEXT,
			total_kills_taser TEXT,
			total_shots_m4a1 TEXT,
			total_shots_galilar TEXT,
			total_shots_taser TEXT,
			total_hits_m4a1 TEXT,
			total_hits_galilar TEXT,
			total_matches_won_train TEXT,
			total_matches_won_lake TEXT,
			GI_lesson_csgo_instr_explain_buymenu TEXT,
			GI_lesson_csgo_instr_explain_buyarmor TEXT,
			GI_lesson_csgo_instr_explain_plant_bomb TEXT,
			GI_lesson_csgo_instr_explain_bomb_carrier TEXT,
			GI_lesson_bomb_sites_A TEXT,
			GI_lesson_defuse_planted_bomb TEXT,
			GI_lesson_csgo_instr_explain_follow_bomber TEXT,
			GI_lesson_csgo_instr_explain_pickup_bomb TEXT,
			GI_lesson_csgo_instr_explain_prevent_bomb_pickup TEXT,
			GI_lesson_Csgo_cycle_weapons_kb TEXT,
			GI_lesson_csgo_instr_explain_zoom TEXT,
			GI_lesson_csgo_instr_explain_reload TEXT,
			GI_lesson_tr_explain_plant_bomb TEXT,
			GI_lesson_bomb_sites_B TEXT,
			GI_lesson_version_number TEXT,
			GI_lesson_find_planted_bomb TEXT,
			GI_lesson_csgo_hostage_lead_to_hrz TEXT,
			GI_lesson_csgo_instr_rescue_zone TEXT,
			GI_lesson_csgo_instr_explain_inspect TEXT,
			steam_stat_xpearnedgames TEXT)`)
	if err != nil {
		return err
	}

	ds.statements["create_player_extra"], err = ds.db.Prepare(
		`CREATE TABLE IF NOT EXISTS player_extra (
			steamid TEXT PRIMARY KEY,
			total_kd TEXT,
			last_match_kd TEXT,
			hit_ratio TEXT,
			played_hours TEXT)`)
	if err != nil {
		return err
	}

	ds.statements["create_recently_played"], err = ds.db.Prepare(
		`CREATE TABLE IF NOT EXISTS recently_played (
			steamid TEXT PRIMARY KEY,
			playtime_2weeks TEXT,
			playtime_forever TEXT,
			playtime_windows_forever TEXT,
			playtime_mac_forever TEXT,
			playtime_linux_forever TEXT)`)
	if err != nil {
		return err
	}

	// TODO add all fields for which we want historical info
	ds.statements["create_player_history"], err = ds.db.Prepare(
		`CREATE TABLE IF NOT EXISTS player_history (
			steamid TEXT,
			time TEXT,
			total_kills TEXT)`)

	return err
}
func (ds *DataStorage) getUpdatePreparedstatements() error {
	var err error

	ds.statements["update_recently_played"], err = ds.db.Prepare(
		`INSERT OR REPLACE INTO recently_played (
			steamid,
			playtime_2weeks,
			playtime_forever,
			playtime_windows_forever,
			playtime_mac_forever,
			playtime_linux_forever)
		VALUES (?,?,?,?,?,?)`)

	if err != nil {
		log.Fatal("Failed to prepare statement: update_recently_played")
		return err
	}

	// - update player_extra
	ds.statements["update_player_extra"], err = ds.db.Prepare(
		`INSERT OR REPLACE INTO player_extra(
				steamid,
				total_kd,
				last_match_kd,
				hit_ratio,
				played_hours)
			VALUES ( ?,?,?,?,?) `)
	if err != nil {
		log.Fatal("Failed to prepare statement: update_player_extra")
		return err
	}

	// - update player_summary
	ds.statements["update_player_summary"], err = ds.db.Prepare(
		`INSERT OR REPLACE INTO player_summary (
				steamid,
				communityvisibilitystate,
				profilestate,
				personaname,
				profileurl,
				avatar,
				avatarmedium,
				avatarfull,
				lastlogoff,
				personastate,
				primaryclanid,
				timecreated)
			VALUES ( ?,?,?,?,?,?,?,?,?,?,?,?) `)
	if err != nil {
		log.Fatal("Failed to prepare statement: update_player_summary")
		return err
	}

	ds.statements["update_player_stats"], err = ds.db.Prepare(
		`INSERT OR REPLACE INTO player_stats (
		    steamid,
			total_kills,
			total_deaths,
			total_time_played,
			total_planted_bombs,
			total_defused_bombs,
			total_wins,
			total_damage_done,
			total_money_earned,
			total_kills_knife,
			total_kills_hegrenade,
			total_kills_glock,
			total_kills_deagle,
			total_kills_elite,
			total_kills_fiveseven,
			total_kills_xm1014,
			total_kills_mac10,
			total_kills_ump45,
			total_kills_p90,
			total_kills_awp,
			total_kills_ak47,
			total_kills_aug,
			total_kills_famas,
			total_kills_g3sg1,
			total_kills_m249,
			total_kills_headshot,
			total_kills_enemy_weapon,
			total_wins_pistolround,
			total_wins_map_cs_assault,
			total_wins_map_de_dust2,
			total_wins_map_de_inferno,
			total_wins_map_de_train,
			total_weapons_donated,
			total_kills_enemy_blinded,
			total_kills_knife_fight,
			total_kills_against_zoomed_sniper,
			total_dominations,
			total_domination_overkills,
			total_revenges,
			total_shots_hit,
			total_shots_fired,
			total_rounds_played,
			total_shots_deagle,
			total_shots_glock,
			total_shots_elite,
			total_shots_fiveseven,
			total_shots_awp,
			total_shots_ak47,
			total_shots_aug,
			total_shots_famas,
			total_shots_g3sg1,
			total_shots_p90,
			total_shots_mac10,
			total_shots_ump45,
			total_shots_xm1014,
			total_shots_m249,
			total_hits_deagle,
			total_hits_glock,
			total_hits_elite,
			total_hits_fiveseven,
			total_hits_awp,
			total_hits_ak47,
			total_hits_aug,
			total_hits_famas,
			total_hits_g3sg1,
			total_hits_p90,
			total_hits_mac10,
			total_hits_ump45,
			total_hits_xm1014,
			total_hits_m249,
			total_rounds_map_cs_assault,
			total_rounds_map_de_dust2,
			total_rounds_map_de_inferno,
			total_rounds_map_de_train,
			last_match_t_wins,
			last_match_ct_wins,
			last_match_wins,
			last_match_max_players,
			last_match_kills,
			last_match_deaths,
			last_match_mvps,
			last_match_favweapon_id,
			last_match_favweapon_shots,
			last_match_favweapon_hits,
			last_match_favweapon_kills,
			last_match_damage,
			last_match_money_spent,
			last_match_dominations,
			last_match_revenges,
			total_mvps,
			total_rounds_map_de_lake,
			total_rounds_map_de_safehouse,
			total_rounds_map_de_bank,
			total_TR_planted_bombs,
			total_gun_game_rounds_won,
			total_gun_game_rounds_played,
			total_wins_map_de_bank,
			total_wins_map_de_lake,
			total_matches_won_bank,
			total_matches_won,
			total_matches_played,
			total_gg_matches_won,
			total_gg_matches_played,
			total_progressive_matches_won,
			total_trbomb_matches_won,
			total_contribution_score,
			last_match_contribution_score,
			last_match_rounds,
			total_kills_hkp2000,
			total_shots_hkp2000,
			total_hits_hkp2000,
			total_hits_p250,
			total_kills_p250,
			total_shots_p250,
			total_kills_sg556,
			total_shots_sg556,
			total_hits_sg556,
			total_hits_scar20,
			total_kills_scar20,
			total_shots_scar20,
			total_shots_ssg08,
			total_hits_ssg08,
			total_kills_ssg08,
			total_shots_mp7,
			total_hits_mp7,
			total_kills_mp7,
			total_kills_mp9,
			total_shots_mp9,
			total_hits_mp9,
			total_hits_nova,
			total_kills_nova,
			total_shots_nova,
			total_hits_negev,
			total_kills_negev,
			total_shots_negev,
			total_shots_sawedoff,
			total_hits_sawedoff,
			total_kills_sawedoff,
			total_shots_bizon,
			total_hits_bizon,
			total_kills_bizon,
			total_kills_tec9,
			total_shots_tec9,
			total_hits_tec9,
			total_shots_mag7,
			total_hits_mag7,
			total_kills_mag7,
			total_gun_game_contribution_score,
			last_match_gg_contribution_score,
			total_kills_m4a1,
			total_kills_galilar,
			total_kills_molotov,
			total_kills_taser,
			total_shots_m4a1,
			total_shots_galilar,
			total_shots_taser,
			total_hits_m4a1,
			total_hits_galilar,
			total_matches_won_train,
			total_matches_won_lake,
			GI_lesson_csgo_instr_explain_buymenu,
			GI_lesson_csgo_instr_explain_buyarmor,
			GI_lesson_csgo_instr_explain_plant_bomb,
			GI_lesson_csgo_instr_explain_bomb_carrier,
			GI_lesson_bomb_sites_A,
			GI_lesson_defuse_planted_bomb,
			GI_lesson_csgo_instr_explain_follow_bomber,
			GI_lesson_csgo_instr_explain_pickup_bomb,
			GI_lesson_csgo_instr_explain_prevent_bomb_pickup,
			GI_lesson_Csgo_cycle_weapons_kb,
			GI_lesson_csgo_instr_explain_zoom,
			GI_lesson_csgo_instr_explain_reload,
			GI_lesson_tr_explain_plant_bomb,
			GI_lesson_bomb_sites_B,
			GI_lesson_version_number,
			GI_lesson_find_planted_bomb,
			GI_lesson_csgo_hostage_lead_to_hrz,
			GI_lesson_csgo_instr_rescue_zone,
			GI_lesson_csgo_instr_explain_inspect,
			steam_stat_xpearnedgames)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)

	if err != nil {
		log.Fatal("Failed to prepare statement: update_player_stats")
	}
	return err
}
func (ds *DataStorage) getInsertPreparedstatements() error {
	var err error
	ds.statements["insert_history"], err = ds.db.Prepare(
		`INSERT INTO player_history (
			steamid,
			time,
			total_kills
		) VALUES (?, ?, ?)`)
	return err
}

func (ds *DataStorage) getSelectPreparedstatements() error {
	// Prepare all statements
	var err error

	// - query player_summary for player
	if ds.statements["select_player_summary"], err = ds.db.Prepare(`
			SELECT * FROM player_summary
			WHERE steamid = ?
			LIMIT 1`); err != nil {
		return err
	}

	// - query player_extra for player
	if ds.statements["select_player_extra"], err = ds.db.Prepare(`
			SELECT * FROM player_extra
			WHERE steamid = ?
			LIMIT 1`); err != nil {
		return err
	}

	// - query player_stats
	if ds.statements["select_player_stats"], err = ds.db.Prepare(`
			SELECT * FROM player_stats
			WHERE steamid = ?
			LIMIT 1`); err != nil {
		return err
	}
	// - query recently_played
	if ds.statements["select_recently_played"], err = ds.db.Prepare(`
			SELECT * FROM recently_played
			WHERE steamid = ?
			LIMIT 1`); err != nil {
		return err
	}

	// Statements for player_history

	// Query last 10 entries for steamID
	if ds.statements["select_player_history"], err = ds.db.Prepare(`
			SELECT * FROM player_history
			WHERE steamid = ?
			ORDER BY time DESC
			LIMIT 10`); err != nil {
		return err
	}

	// Get latest timestamp for steamID
	if ds.statements["select_player_history_latest_time"], err = ds.db.Prepare(`
			SELECT time FROM player_history
			WHERE steamid = ?
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
