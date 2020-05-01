package main

import "log"

func (ds *DataStorage) getCreatePreparedstatements() error {
	var err error

	ds.statements["create_player_summary"], err = ds.db.Prepare(
		`CREATE TABLE IF NOT EXISTS player_summary (
			steamid INTEGER PRIMARY KEY,
			communityvisibilitystate INTEGER,
			profilestate INTEGER,
			personaname TEXT,
			profileurl TEXT,
			avatar TEXT,
			avatarmedium TEXT,
			avatarfull TEXT,
			lastlogoff INTEGER,
			personastate INTEGER,
			primaryclanid INTEGER,
			timecreated INTEGER)`)

	if err != nil {
		return err
	}

	ds.statements["create_player_stats"], err = ds.db.Prepare(
		`CREATE TABLE IF NOT EXISTS player_stats (
			steamid INTEGER PRIMARY KEY,
			total_kills INTEGER,
			total_deaths INTEGER,
			total_time_played INTEGER,
			total_planted_bombs INTEGER,
			total_defused_bombs INTEGER,
			total_wins INTEGER,
			total_damage_done INTEGER,
			total_money_earned INTEGER,
			total_kills_knife INTEGER,
			total_kills_hegrenade INTEGER,
			total_kills_glock INTEGER,
			total_kills_deagle INTEGER,
			total_kills_elite INTEGER,
			total_kills_fiveseven INTEGER,
			total_kills_xm1014 INTEGER,
			total_kills_mac10 INTEGER,
			total_kills_ump45 INTEGER,
			total_kills_p90 INTEGER,
			total_kills_awp INTEGER,
			total_kills_ak47 INTEGER,
			total_kills_aug INTEGER,
			total_kills_famas INTEGER,
			total_kills_g3sg1 INTEGER,
			total_kills_m249 INTEGER,
			total_kills_headshot INTEGER,
			total_kills_enemy_weapon INTEGER,
			total_wins_pistolround INTEGER,
			total_wins_map_cs_assault INTEGER,
			total_wins_map_de_dust2 INTEGER,
			total_wins_map_de_inferno INTEGER,
			total_wins_map_de_train INTEGER,
			total_weapons_donated INTEGER,
			total_kills_enemy_blinded INTEGER,
			total_kills_knife_fight INTEGER,
			total_kills_against_zoomed_sniper INTEGER,
			total_dominations INTEGER,
			total_domination_overkills INTEGER,
			total_revenges INTEGER,
			total_shots_hit INTEGER,
			total_shots_fired INTEGER,
			total_rounds_played INTEGER,
			total_shots_deagle INTEGER,
			total_shots_glock INTEGER,
			total_shots_elite INTEGER,
			total_shots_fiveseven INTEGER,
			total_shots_awp INTEGER,
			total_shots_ak47 INTEGER,
			total_shots_aug INTEGER,
			total_shots_famas INTEGER,
			total_shots_g3sg1 INTEGER,
			total_shots_p90 INTEGER,
			total_shots_mac10 INTEGER,
			total_shots_ump45 INTEGER,
			total_shots_xm1014 INTEGER,
			total_shots_m249 INTEGER,
			total_hits_deagle INTEGER,
			total_hits_glock INTEGER,
			total_hits_elite INTEGER,
			total_hits_fiveseven INTEGER,
			total_hits_awp INTEGER,
			total_hits_ak47 INTEGER,
			total_hits_aug INTEGER,
			total_hits_famas INTEGER,
			total_hits_g3sg1 INTEGER,
			total_hits_p90 INTEGER,
			total_hits_mac10 INTEGER,
			total_hits_ump45 INTEGER,
			total_hits_xm1014 INTEGER,
			total_hits_m249 INTEGER,
			total_rounds_map_cs_assault INTEGER,
			total_rounds_map_de_dust2 INTEGER,
			total_rounds_map_de_inferno INTEGER,
			total_rounds_map_de_train INTEGER,
			last_match_t_wins INTEGER,
			last_match_ct_wins INTEGER,
			last_match_wins INTEGER,
			last_match_max_players INTEGER,
			last_match_kills INTEGER,
			last_match_deaths INTEGER,
			last_match_mvps INTEGER,
			last_match_favweapon_id INTEGER,
			last_match_favweapon_shots INTEGER,
			last_match_favweapon_hits INTEGER,
			last_match_favweapon_kills INTEGER,
			last_match_damage INTEGER,
			last_match_money_spent INTEGER,
			last_match_dominations INTEGER,
			last_match_revenges INTEGER,
			total_mvps INTEGER,
			total_rounds_map_de_lake INTEGER,
			total_rounds_map_de_safehouse INTEGER,
			total_rounds_map_de_bank INTEGER,
			total_TR_planted_bombs INTEGER,
			total_gun_game_rounds_won INTEGER,
			total_gun_game_rounds_played INTEGER,
			total_wins_map_de_bank INTEGER,
			total_wins_map_de_lake INTEGER,
			total_matches_won_bank INTEGER,
			total_matches_won INTEGER,
			total_matches_played INTEGER,
			total_gg_matches_won INTEGER,
			total_gg_matches_played INTEGER,
			total_progressive_matches_won INTEGER,
			total_trbomb_matches_won INTEGER,
			total_contribution_score INTEGER,
			last_match_contribution_score INTEGER,
			last_match_rounds INTEGER,
			total_kills_hkp2000 INTEGER,
			total_shots_hkp2000 INTEGER,
			total_hits_hkp2000 INTEGER,
			total_hits_p250 INTEGER,
			total_kills_p250 INTEGER,
			total_shots_p250 INTEGER,
			total_kills_sg556 INTEGER,
			total_shots_sg556 INTEGER,
			total_hits_sg556 INTEGER,
			total_hits_scar20 INTEGER,
			total_kills_scar20 INTEGER,
			total_shots_scar20 INTEGER,
			total_shots_ssg08 INTEGER,
			total_hits_ssg08 INTEGER,
			total_kills_ssg08 INTEGER,
			total_shots_mp7 INTEGER,
			total_hits_mp7 INTEGER,
			total_kills_mp7 INTEGER,
			total_kills_mp9 INTEGER,
			total_shots_mp9 INTEGER,
			total_hits_mp9 INTEGER,
			total_hits_nova INTEGER,
			total_kills_nova INTEGER,
			total_shots_nova INTEGER,
			total_hits_negev INTEGER,
			total_kills_negev INTEGER,
			total_shots_negev INTEGER,
			total_shots_sawedoff INTEGER,
			total_hits_sawedoff INTEGER,
			total_kills_sawedoff INTEGER,
			total_shots_bizon INTEGER,
			total_hits_bizon INTEGER,
			total_kills_bizon INTEGER,
			total_kills_tec9 INTEGER,
			total_shots_tec9 INTEGER,
			total_hits_tec9 INTEGER,
			total_shots_mag7 INTEGER,
			total_hits_mag7 INTEGER,
			total_kills_mag7 INTEGER,
			total_gun_game_contribution_score INTEGER,
			last_match_gg_contribution_score INTEGER,
			total_kills_m4a1 INTEGER,
			total_kills_galilar INTEGER,
			total_kills_molotov INTEGER,
			total_kills_taser INTEGER,
			total_shots_m4a1 INTEGER,
			total_shots_galilar INTEGER,
			total_shots_taser INTEGER,
			total_hits_m4a1 INTEGER,
			total_hits_galilar INTEGER,
			total_matches_won_train INTEGER,
			total_matches_won_lake INTEGER,
			GI_lesson_csgo_instr_explain_buymenu INTEGER,
			GI_lesson_csgo_instr_explain_buyarmor INTEGER,
			GI_lesson_csgo_instr_explain_plant_bomb INTEGER,
			GI_lesson_csgo_instr_explain_bomb_carrier INTEGER,
			GI_lesson_bomb_sites_A INTEGER,
			GI_lesson_defuse_planted_bomb INTEGER,
			GI_lesson_csgo_instr_explain_follow_bomber INTEGER,
			GI_lesson_csgo_instr_explain_pickup_bomb INTEGER,
			GI_lesson_csgo_instr_explain_prevent_bomb_pickup INTEGER,
			GI_lesson_Csgo_cycle_weapons_kb INTEGER,
			GI_lesson_csgo_instr_explain_zoom INTEGER,
			GI_lesson_csgo_instr_explain_reload INTEGER,
			GI_lesson_tr_explain_plant_bomb INTEGER,
			GI_lesson_bomb_sites_B INTEGER,
			GI_lesson_version_number INTEGER,
			GI_lesson_find_planted_bomb INTEGER,
			GI_lesson_csgo_hostage_lead_to_hrz INTEGER,
			GI_lesson_csgo_instr_rescue_zone INTEGER,
			GI_lesson_csgo_instr_explain_inspect INTEGER,
			steam_stat_xpearnedgames INTEGER)`)
	if err != nil {
		return err
	}

	ds.statements["create_recently_played"], err = ds.db.Prepare(
		`CREATE TABLE IF NOT EXISTS recently_played (
			steamid INTEGER PRIMARY KEY,
			playtime_2weeks INTEGER,
			playtime_forever INTEGER,
			playtime_windows_forever INTEGER,
			playtime_mac_forever INTEGER,
			playtime_linux_forever INTEGER)`)
	if err != nil {
		return err
	}

	// TODO add all fields for which we want historical info
	ds.statements["create_player_history"], err = ds.db.Prepare(
		`CREATE TABLE IF NOT EXISTS player_history (
			steamid INTEGER,
			time INTEGER,
			total_kills INTEGER)`)

	return err
}
func (ds *DataStorage) getUpdatePreparedstatements() error {
	var err error

	ds.statements["update_recently_played"], err = ds.db.Prepare(
		`UPDATE recently_played SET
			playtime_2weeks = ?,
			playtime_forever= ?,
			playtime_windows_forever= ?,
			playtime_mac_forever = ?,
			playtime_linux_forever = ?
			WHERE steamid = ?`)

	if err != nil {
		log.Println("Failed to prepare statement: update_recently_played")
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
		log.Println("Failed to prepare statement: update_player_summary")
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
		log.Println("Failed to prepare statement: update_player_stats")
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

	// - insert histor	// - query player_summary for player
	ds.statements["select_player_summary"], err = ds.db.Prepare(`
			SELECT * FROM player_summary
			WHERE steamid = ?
			LIMIT 1`)

	// - query player_stats
	ds.statements["select_player_stats"], err = ds.db.Prepare(`
			SELECT * FROM player_stats
			WHERE steamid = ?
			LIMIT 1`)
	// - query recently_played
	ds.statements["select_recently_played"], err = ds.db.Prepare(`
			SELECT * FROM recently_played
			WHERE steamid = ?
			LIMIT 1`)
	// - query history for last n entries of player
	ds.statements["select_player_history"], err = ds.db.Prepare(`
			SELECT * FROM player_history
			WHERE steamid = ?
			ORDER BY time DESC
			LIMIT ?`)

	return err
}
