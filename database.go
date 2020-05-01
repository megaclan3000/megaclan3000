package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Public data query methods
func (ds *DataStorage) GetPlayerSummary(steamID string) (PlayerSummary, error) {

	ps := PlayerSummary{}
	var err error

	if rows, err := ds.statements["select_player_summary"].Query(); err == nil {
		rows.Scan(
			&ps.Steamid,
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
	return ps, err
}

func (ds *DataStorage) GetUserStatsForGame(steamID string) (UserStatsForGame, error) {
	//TODO impletment
	panic("not implemented yet!")
}

func (ds *DataStorage) GetRecentlyPlayedGames(steamID string) (RecentlyPlayedGames, error) {
	rpg := RecentlyPlayedGames{}
	var err error
	var id int

	if rows, err := ds.statements["select_player_stats"].Query(); err == nil {
		rows.Scan(
			&id,
			&rpg.Playtime2Weeks,
			&rpg.PlaytimeForever,
			&rpg.PlaytimeWindowsForever,
			&rpg.PlaytimeMacForever,
			&rpg.PlaytimeLinuxForever,
		)
	}
	return rpg, err
}

func (ds *DataStorage) GetPlayerHistory(steamID string, numPoints int) (PlayerHistory, error) {
	ph := PlayerHistory{}
	var err error

	if rows, err := ds.statements["select_player_history"].Query(); err == nil {
		rows.Scan(
			&ph.steamID,
			&ph.time,
			&ph.TotalKills,
		)
	}
	return ph, err
}

// Private data retrieval methods
func (ds *DataStorage) updatePlayerSummary(steamID string) {

	ps := getPlayerSummary(steamID)

	if _, err := ds.statements["update_player_summary"].Exec(
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
		steamID,
	); err != nil {
		log.Fatal(err)
	}
}

func (ds *DataStorage) updateUserStatsForGame(steamID string) {
	stats := getUserStatsForGame(steamID)
	if _, err := ds.statements["update_player_stats"].Exec(
		stats.Stats["total_kills"],
		stats.Stats["total_deaths"],
		stats.Stats["total_time_played"],
		stats.Stats["total_planted_bombs"],
		stats.Stats["total_defused_bombs"],
		stats.Stats["total_wins"],
		stats.Stats["total_damage_done"],
		stats.Stats["total_money_earned"],
		stats.Stats["total_kills_knife"],
		stats.Stats["total_kills_hegrenade"],
		stats.Stats["total_kills_glock"],
		stats.Stats["total_kills_deagle"],
		stats.Stats["total_kills_elite"],
		stats.Stats["total_kills_fiveseven"],
		stats.Stats["total_kills_xm1014"],
		stats.Stats["total_kills_mac10"],
		stats.Stats["total_kills_ump45"],
		stats.Stats["total_kills_p90"],
		stats.Stats["total_kills_awp"],
		stats.Stats["total_kills_ak47"],
		stats.Stats["total_kills_aug"],
		stats.Stats["total_kills_famas"],
		stats.Stats["total_kills_g3sg1"],
		stats.Stats["total_kills_m249"],
		stats.Stats["total_kills_headshot"],
		stats.Stats["total_kills_enemy_weapon"],
		stats.Stats["total_wins_pistolround"],
		stats.Stats["total_wins_map_cs_assault"],
		stats.Stats["total_wins_map_de_dust2"],
		stats.Stats["total_wins_map_de_inferno"],
		stats.Stats["total_wins_map_de_train"],
		stats.Stats["total_weapons_donated"],
		stats.Stats["total_kills_enemy_blinded"],
		stats.Stats["total_kills_knife_fight"],
		stats.Stats["total_kills_against_zoomed_sniper"],
		stats.Stats["total_dominations"],
		stats.Stats["total_domination_overkills"],
		stats.Stats["total_revenges"],
		stats.Stats["total_shots_hit"],
		stats.Stats["total_shots_fired"],
		stats.Stats["total_rounds_played"],
		stats.Stats["total_shots_deagle"],
		stats.Stats["total_shots_glock"],
		stats.Stats["total_shots_elite"],
		stats.Stats["total_shots_fiveseven"],
		stats.Stats["total_shots_awp"],
		stats.Stats["total_shots_ak47"],
		stats.Stats["total_shots_aug"],
		stats.Stats["total_shots_famas"],
		stats.Stats["total_shots_g3sg1"],
		stats.Stats["total_shots_p90"],
		stats.Stats["total_shots_mac10"],
		stats.Stats["total_shots_ump45"],
		stats.Stats["total_shots_xm1014"],
		stats.Stats["total_shots_m249"],
		stats.Stats["total_hits_deagle"],
		stats.Stats["total_hits_glock"],
		stats.Stats["total_hits_elite"],
		stats.Stats["total_hits_fiveseven"],
		stats.Stats["total_hits_awp"],
		stats.Stats["total_hits_ak47"],
		stats.Stats["total_hits_aug"],
		stats.Stats["total_hits_famas"],
		stats.Stats["total_hits_g3sg1"],
		stats.Stats["total_hits_p90"],
		stats.Stats["total_hits_mac10"],
		stats.Stats["total_hits_ump45"],
		stats.Stats["total_hits_xm1014"],
		stats.Stats["total_hits_m249"],
		stats.Stats["total_rounds_map_cs_assault"],
		stats.Stats["total_rounds_map_de_dust2"],
		stats.Stats["total_rounds_map_de_inferno"],
		stats.Stats["total_rounds_map_de_train"],
		stats.Stats["last_match_t_wins"],
		stats.Stats["last_match_ct_wins"],
		stats.Stats["last_match_wins"],
		stats.Stats["last_match_max_players"],
		stats.Stats["last_match_kills"],
		stats.Stats["last_match_deaths"],
		stats.Stats["last_match_mvps"],
		stats.Stats["last_match_favweapon_id"],
		stats.Stats["last_match_favweapon_shots"],
		stats.Stats["last_match_favweapon_hits"],
		stats.Stats["last_match_favweapon_kills"],
		stats.Stats["last_match_damage"],
		stats.Stats["last_match_money_spent"],
		stats.Stats["last_match_dominations"],
		stats.Stats["last_match_revenges"],
		stats.Stats["total_mvps"],
		stats.Stats["total_rounds_map_de_lake"],
		stats.Stats["total_rounds_map_de_safehouse"],
		stats.Stats["total_rounds_map_de_bank"],
		stats.Stats["total_TR_planted_bombs"],
		stats.Stats["total_gun_game_rounds_won"],
		stats.Stats["total_gun_game_rounds_played"],
		stats.Stats["total_wins_map_de_bank"],
		stats.Stats["total_wins_map_de_lake"],
		stats.Stats["total_matches_won_bank"],
		stats.Stats["total_matches_won"],
		stats.Stats["total_matches_played"],
		stats.Stats["total_gg_matches_won"],
		stats.Stats["total_gg_matches_played"],
		stats.Stats["total_progressive_matches_won"],
		stats.Stats["total_trbomb_matches_won"],
		stats.Stats["total_contribution_score"],
		stats.Stats["last_match_contribution_score"],
		stats.Stats["last_match_rounds"],
		stats.Stats["total_kills_hkp2000"],
		stats.Stats["total_shots_hkp2000"],
		stats.Stats["total_hits_hkp2000"],
		stats.Stats["total_hits_p250"],
		stats.Stats["total_kills_p250"],
		stats.Stats["total_shots_p250"],
		stats.Stats["total_kills_sg556"],
		stats.Stats["total_shots_sg556"],
		stats.Stats["total_hits_sg556"],
		stats.Stats["total_hits_scar20"],
		stats.Stats["total_kills_scar20"],
		stats.Stats["total_shots_scar20"],
		stats.Stats["total_shots_ssg08"],
		stats.Stats["total_hits_ssg08"],
		stats.Stats["total_kills_ssg08"],
		stats.Stats["total_shots_mp7"],
		stats.Stats["total_hits_mp7"],
		stats.Stats["total_kills_mp7"],
		stats.Stats["total_kills_mp9"],
		stats.Stats["total_shots_mp9"],
		stats.Stats["total_hits_mp9"],
		stats.Stats["total_hits_nova"],
		stats.Stats["total_kills_nova"],
		stats.Stats["total_shots_nova"],
		stats.Stats["total_hits_negev"],
		stats.Stats["total_kills_negev"],
		stats.Stats["total_shots_negev"],
		stats.Stats["total_shots_sawedoff"],
		stats.Stats["total_hits_sawedoff"],
		stats.Stats["total_kills_sawedoff"],
		stats.Stats["total_shots_bizon"],
		stats.Stats["total_hits_bizon"],
		stats.Stats["total_kills_bizon"],
		stats.Stats["total_kills_tec9"],
		stats.Stats["total_shots_tec9"],
		stats.Stats["total_hits_tec9"],
		stats.Stats["total_shots_mag7"],
		stats.Stats["total_hits_mag7"],
		stats.Stats["total_kills_mag7"],
		stats.Stats["total_gun_game_contribution_score"],
		stats.Stats["last_match_gg_contribution_score"],
		stats.Stats["total_kills_m4a1"],
		stats.Stats["total_kills_galilar"],
		stats.Stats["total_kills_molotov"],
		stats.Stats["total_kills_taser"],
		stats.Stats["total_shots_m4a1"],
		stats.Stats["total_shots_galilar"],
		stats.Stats["total_shots_taser"],
		stats.Stats["total_hits_m4a1"],
		stats.Stats["total_hits_galilar"],
		stats.Stats["total_matches_won_train"],
		stats.Stats["total_matches_won_lake"],
		stats.Stats["GI_lesson_csgo_instr_explain_buymenu"],
		stats.Stats["GI_lesson_csgo_instr_explain_buyarmor"],
		stats.Stats["GI_lesson_csgo_instr_explain_plant_bomb"],
		stats.Stats["GI_lesson_csgo_instr_explain_bomb_carrier"],
		stats.Stats["GI_lesson_bomb_sites_A"],
		stats.Stats["GI_lesson_defuse_planted_bomb"],
		stats.Stats["GI_lesson_csgo_instr_explain_follow_bomber"],
		stats.Stats["GI_lesson_csgo_instr_explain_pickup_bomb"],
		stats.Stats["GI_lesson_csgo_instr_explain_prevent_bomb_pickup"],
		stats.Stats["GI_lesson_Csgo_cycle_weapons_kb"],
		stats.Stats["GI_lesson_csgo_instr_explain_zoom"],
		stats.Stats["GI_lesson_csgo_instr_explain_reload"],
		stats.Stats["GI_lesson_tr_explain_plant_bomb"],
		stats.Stats["GI_lesson_bomb_sites_B"],
		stats.Stats["GI_lesson_version_number"],
		stats.Stats["GI_lesson_find_planted_bomb"],
		stats.Stats["GI_lesson_csgo_hostage_lead_to_hrz"],
		stats.Stats["GI_lesson_csgo_instr_rescue_zone"],
		stats.Stats["GI_lesson_csgo_instr_explain_inspect"],
		stats.Stats["steam_stat_xpearnedgames"],
		steamID,
	); err != nil {
		log.Fatal(err)
	}

}

func (ds *DataStorage) updateRecentlyPlayedGames(steamID string) {

	rp := getRecentlyPlayedGames(steamID)

	if _, err := ds.statements["update_recently_played"].Exec(
		rp.Playtime2Weeks,
		rp.PlaytimeForever,
		rp.PlaytimeWindowsForever,
		rp.PlaytimeMacForever,
		rp.PlaytimeLinuxForever,
		steamID,
	); err != nil {
		log.Fatal(err)
	}
}

type DataStorage struct {
	db         *sql.DB
	statements map[string]*sql.Stmt
}

func NewDataStorage(path string) (*DataStorage, error) {
	var err error

	// Initialize database
	storage := new(DataStorage)
	storage.statements = make(map[string]*sql.Stmt)

	log.Println("Reading", path)
	if storage.db, err = sql.Open("sqlite3", path); err != nil {
		log.Fatal("Failed to open sqlite file", err)
	}

	// Prepare all statements
	log.Println("Preparing CREATE statements")

	if err = storage.getCreatePreparedstatements(); err != nil {
		log.Fatal("Failed to prepare CREATE statements", err)
	}

	// Create tables, if necessary
	log.Println("Creating player_summary table")
	if _, err = storage.statements["create_player_summary"].Exec(); err != nil {
		log.Fatal("Failed to create table player_summary", err)
	}

	log.Println("Creating player_stats table")
	if _, err = storage.statements["create_player_stats"].Exec(); err != nil {
		log.Fatal("Failed to create table player_stats", err)
	}
	log.Println("Creating recetnly_played table")
	if _, err = storage.statements["create_recently_played"].Exec(); err != nil {
		log.Fatal("Failed to create table recently_played", err)
	}
	log.Println("Creating player_history table")
	if _, err = storage.statements["create_player_history"].Exec(); err != nil {
		log.Fatal("Failed to create table player_history", err)
	}

	log.Println("Preparing UPDATE statements")
	if err = storage.getUpdatePreparedstatements(); err != nil {
		log.Fatal("Failed to prepare UPDATE statements", err)
	}

	log.Println("Preparing INSERT statements")
	if err = storage.getInsertPreparedstatements(); err != nil {
		log.Fatal("Failed to prepare INSERT statements", err)
	}

	log.Println("Preparing SELECT statements")
	if err = storage.getSelectPreparedstatements(); err != nil {
		log.Fatal("Failed to prepare SELECT statements", err)
	}

	for _, v := range config.SteamIDs {
		log.Println("Updating PlayerSummary for ID:", v)
		storage.updatePlayerSummary(v)

		log.Println("Updating RecentlyPlayedGames for ID:", v)
		storage.updateRecentlyPlayedGames(v)

		log.Println("Updating UserStatsForGame for ID:", v)
		storage.updateUserStatsForGame(v)
	}

	return storage, nil
}
