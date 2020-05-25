package steamclient

import (
	"fmt"
	"strconv"
)

// https://developer.valvesoftware.com/wiki/Steam_Web_API#GetUserStatsForGame_.28v0002.29
// Returns a list of achievements for this user by app id

type userStatsForGameData struct {
	Playerstats struct {
		SteamID      string `json:"steamID"`
		GameName     string `json:"gameName"`
		Stats        []Stats
		Achievements []Achievements `json:"achievements"`
	} `json:"playerstats"`
}

// Stats is the nested struct to hold the "stats" array returned by the steam
// API endopint GetUserStatsForGame
type Stats struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

// Achievements is the nested struct to hold the "archivements" array returned
// by the steam API endopint GetUserStatsForGame
type Achievements struct {
	Name     string `json:"name"`
	Achieved int    `json:"achieved"`
}

// UserStatsForGame holds the players summary data from the steam API
// endpoint GetUserStatsForGame
type UserStatsForGame struct {

	// SteamID of the player
	SteamID string

	// Name of the game
	GameName string

	// GamesStats object containing stats for the game
	Stats GameStats

	// GameArchievements object containing archievements for the game
	Archivements GameArchievements

	// GameExtras object containig additional, calculated information
	Extra GameExtras
}

// GameExtras holds data in the same way as the other nested structs. This data
// is not fetched from an endpoint but calculated based on other values locally
type GameExtras struct {

	// Total KD ratio
	TotalKD string

	// KD ratio of the last match
	LastMatchKD string

	// Total hit ratio
	HitRatio string

	// Total played hours
	PlayedHours string

	// Total Average damage per round
	TotalADR string

	// Last match average damage per round
	LastMatchADR string
}

func (sc *SteamClient) ParseUserStatsForGame(data userStatsForGameData) (UserStatsForGame, error) {

	//Create to maps for stats and archivements, so the search will be quicker afterwards
	statsMap := make(map[string]string)
	archivementsMap := make(map[string]string)

	for _, v := range data.Playerstats.Stats {
		statsMap[v.Name] = strconv.Itoa(v.Value)
	}

	for _, v := range data.Playerstats.Achievements {
		archivementsMap[v.Name] = strconv.Itoa(v.Achieved)
	}

	extra := GameExtras{}

	if totalDeaths, err := strconv.ParseFloat(statsMap["total_deaths"], 64); err == nil {
		if totalKills, err := strconv.ParseFloat(statsMap["total_kills"], 64); err == nil {
			extra.TotalKD = fmt.Sprintf("%.3f", totalKills/totalDeaths)
		}
	}

	if lastDeaths, err := strconv.ParseFloat(statsMap["last_match_deaths"], 64); err == nil {
		if lastKills, err := strconv.ParseFloat(statsMap["last_match_kills"], 64); err == nil {
			extra.LastMatchKD = fmt.Sprintf("%.3f", lastKills/lastDeaths)
		}
	}

	if totalShotsFired, err := strconv.ParseFloat(statsMap["total_shots_fired"], 64); err == nil {
		if totalShotsHit, err := strconv.ParseFloat(statsMap["total_shots_hit"], 64); err == nil {
			extra.HitRatio = fmt.Sprintf("%.3f", totalShotsHit/totalShotsFired)
		}
	}

	if secI, err := strconv.Atoi(statsMap["total_time_played"]); err == nil {
		extra.PlayedHours = strconv.Itoa(secI / 3600)
	}

	return UserStatsForGame{
		SteamID:  data.Playerstats.SteamID,
		GameName: data.Playerstats.GameName,
		Extra:    extra,
		Stats: GameStats{
			GILessonCsgoInstrExplainInspect:           statsMap["GI.lesson.csgo_instr_explain_inspect"],
			GILessonBombSitesA:                        statsMap["GI.lesson.bomb_sites_a"],
			GILessonBombSitesB:                        statsMap["GI.lesson.bomb_sites_b"],
			GILessonCsgoCycleWeaponsKb:                statsMap["GI.lesson.csgo_cycle_weapons_kb"],
			GILessonCsgoHostageLeadToHrz:              statsMap["GI.lesson.csgo_hostage_lead_to_hrz"],
			GILessonCsgoInstrExplainBombCarrier:       statsMap["GI.lesson.csgo_instr_explain_bomb_carrier"],
			GILessonCsgoInstrExplainBuyarmor:          statsMap["GI.lesson.csgo_instr_explain_buyarmor"],
			GILessonCsgoInstrExplainBuymenu:           statsMap["GI.lesson.csgo_instr_explain_buymenu"],
			GILessonCsgoInstrExplainFollowBomber:      statsMap["GI.lesson.csgo_instr_explain_follow_bomber"],
			GILessonCsgoInstrExplainPickupBomb:        statsMap["GI.lesson.csgo_instr_explain_pickup_bomb"],
			GILessonCsgoInstrExplainPlantBomb:         statsMap["GI.lesson.csgo_instr_explain_plant_bomb"],
			GILessonCsgoInstrExplainPreventBombPickup: statsMap["GI.lesson.csgo_instr_explain_prevent_bomb_pickup"],
			GILessonCsgoInstrExplainReload:            statsMap["GI.lesson.csgo_instr_explain_reload"],
			GILessonCsgoInstrExplainZoom:              statsMap["GI.lesson.csgo_instr_explain_zoom"],
			GILessonCsgoInstrRescueZone:               statsMap["GI.lesson.csgo_instr_rescue_zone"],
			GILessonDefusePlantedBomb:                 statsMap["GI.lesson.defuse_planted_bomb"],
			GILessonFindPlantedBomb:                   statsMap["GI.lesson.find_planted_bomb"],
			GILessonTrExplainPlantBomb:                statsMap["GI.lesson.tr_explain_plant_bomb"],
			GILessonVersionNumber:                     statsMap["GI.lesson.version_number"],
			LastMatchContributionScore:                statsMap["last_match_contribution_score"],
			LastMatchCtWins:                           statsMap["last_match_ct_wins"],
			LastMatchDamage:                           statsMap["last_match_damage"],
			LastMatchDeaths:                           statsMap["last_match_deaths"],
			LastMatchDominations:                      statsMap["last_match_dominations"],
			LastMatchFavweaponHits:                    statsMap["last_match_favweapon_hits"],
			LastMatchFavweaponID:                      statsMap["last_match_favweapon_id"],
			LastMatchFavweaponKills:                   statsMap["last_match_favweapon_kills"],
			LastMatchFavweaponShots:                   statsMap["last_match_favweapon_shots"],
			LastMatchGgContributionScore:              statsMap["last_match_gg_contribution_score"],
			LastMatchKills:                            statsMap["last_match_kills"],
			LastMatchMaxPlayers:                       statsMap["last_match_max_players"],
			LastMatchMoneySpent:                       statsMap["last_match_money_spent"],
			LastMatchMvps:                             statsMap["last_match_mvps"],
			LastMatchRevenges:                         statsMap["last_match_revenges"],
			LastMatchRounds:                           statsMap["last_match_rounds"],
			LastMatchTWins:                            statsMap["last_match_t_wins"],
			LastMatchWins:                             statsMap["last_match_wins"],
			SteamStatMatchwinscomp:                    statsMap["steam_stat_matchwinscomp"],
			SteamStatSurvivedz:                        statsMap["steam_stat_survivedz"],
			SteamStatXpearnedgames:                    statsMap["steam_stat_xpearnedgames"],
			TotalBrokenWindows:                        statsMap["total_broken_windows"],
			TotalContributionScore:                    statsMap["total_contribution_score"],
			TotalDamageDone:                           statsMap["total_damage_done"],
			TotalDeaths:                               statsMap["total_deaths"],
			TotalDefusedBombs:                         statsMap["total_defused_bombs"],
			TotalDominationOverkills:                  statsMap["total_domination_overkills"],
			TotalDominations:                          statsMap["total_dominations"],
			TotalGgMatchesPlayed:                      statsMap["total_gg_matches_played"],
			TotalGgMatchesWon:                         statsMap["total_gg_matches_won"],
			TotalGunGameContributionScore:             statsMap["total_gun_game_contribution_score"],
			TotalGunGameRoundsPlayed:                  statsMap["total_gun_game_rounds_played"],
			TotalGunGameRoundsWon:                     statsMap["total_gun_game_rounds_won"],
			TotalHitsAk47:                             statsMap["total_hits_ak47"],
			TotalHitsAug:                              statsMap["total_hits_aug"],
			TotalHitsAwp:                              statsMap["total_hits_awp"],
			TotalHitsBizon:                            statsMap["total_hits_bizon"],
			TotalHitsDeagle:                           statsMap["total_hits_deagle"],
			TotalHitsElite:                            statsMap["total_hits_elite"],
			TotalHitsFamas:                            statsMap["total_hits_famas"],
			TotalHitsFiveseven:                        statsMap["total_hits_fiveseven"],
			TotalHitsG3sg1:                            statsMap["total_hits_g_3sg_1"],
			TotalHitsGalilar:                          statsMap["total_hits_galilar"],
			TotalHitsGlock:                            statsMap["total_hits_glock"],
			TotalHitsHkp2000:                          statsMap["total_hits_hkp2000"],
			TotalHitsM249:                             statsMap["total_hits_m249"],
			TotalHitsM4a1:                             statsMap["total_hits_m_4a_1"],
			TotalHitsMac10:                            statsMap["total_hits_mac_10"],
			TotalHitsMag7:                             statsMap["total_hits_mag_7"],
			TotalHitsMp7:                              statsMap["total_hits_mp_7"],
			TotalHitsMp9:                              statsMap["total_hits_mp_9"],
			TotalHitsNegev:                            statsMap["total_hits_negev"],
			TotalHitsNova:                             statsMap["total_hits_nova"],
			TotalHitsP250:                             statsMap["total_hits_p250"],
			TotalHitsP90:                              statsMap["total_hits_p90"],
			TotalHitsSawedoff:                         statsMap["total_hits_sawedoff"],
			TotalHitsScar20:                           statsMap["total_hits_scar20"],
			TotalHitsSg556:                            statsMap["total_hits_sg556"],
			TotalHitsSsg08:                            statsMap["total_hits_ssg08"],
			TotalHitsTec9:                             statsMap["total_hits_tec_9"],
			TotalHitsUmp45:                            statsMap["total_hits_ump_45"],
			TotalHitsXm1014:                           statsMap["total_hits_xm_1014"],
			TotalKills:                                statsMap["total_kills"],
			TotalKillsAgainstZoomedSniper:             statsMap["total_kills_against_zoomed_sniper"],
			TotalKillsAk47:                            statsMap["total_kills_ak47"],
			TotalKillsAug:                             statsMap["total_kills_aug"],
			TotalKillsAwp:                             statsMap["total_kills_awp"],
			TotalKillsBizon:                           statsMap["total_kills_bizon"],
			TotalKillsDeagle:                          statsMap["total_kills_deagle"],
			TotalKillsElite:                           statsMap["total_kills_elite"],
			TotalKillsEnemyBlinded:                    statsMap["total_kills_enemy_blinded"],
			TotalKillsEnemyWeapon:                     statsMap["total_kills_enemy_weapon"],
			TotalKillsFamas:                           statsMap["total_kills_famas"],
			TotalKillsFiveseven:                       statsMap["total_kills_fiveseven"],
			TotalKillsG3sg1:                           statsMap["total_kills_g3sg1"],
			TotalKillsGalilar:                         statsMap["total_kills_galilar"],
			TotalKillsGlock:                           statsMap["total_kills_glock"],
			TotalKillsHeadshot:                        statsMap["total_kills_headshot"],
			TotalKillsHegrenade:                       statsMap["total_kills_hegrenade"],
			TotalKillsHkp2000:                         statsMap["total_kills_hkp2000"],
			TotalKillsKnife:                           statsMap["total_kills_knife"],
			TotalKillsKnifeFight:                      statsMap["total_kills_knife_fight"],
			TotalKillsM249:                            statsMap["total_kills_m249"],
			TotalKillsM4a1:                            statsMap["total_kills_m4a1"],
			TotalKillsMac10:                           statsMap["total_kills_mac10"],
			TotalKillsMag7:                            statsMap["total_kills_mag7"],
			TotalKillsMolotov:                         statsMap["total_kills_molotov"],
			TotalKillsMp7:                             statsMap["total_kills_mp7"],
			TotalKillsMp9:                             statsMap["total_kills_mp9"],
			TotalKillsNegev:                           statsMap["total_kills_negev"],
			TotalKillsNova:                            statsMap["total_kills_nova"],
			TotalKillsP250:                            statsMap["total_kills_p250"],
			TotalKillsP90:                             statsMap["total_kills_p90"],
			TotalKillsSawedoff:                        statsMap["total_kills_sawedoff"],
			TotalKillsScar20:                          statsMap["total_kills_scar20"],
			TotalKillsSg556:                           statsMap["total_kills_sg556"],
			TotalKillsSsg08:                           statsMap["total_kills_ssg08"],
			TotalKillsTaser:                           statsMap["total_kills_taser"],
			TotalKillsTec9:                            statsMap["total_kills_tec9"],
			TotalKillsUmp45:                           statsMap["total_kills_ump45"],
			TotalKillsXm1014:                          statsMap["total_kills_xm1014"],
			TotalMatchesPlayed:                        statsMap["total_matches_played"],
			TotalMatchesWon:                           statsMap["total_matches_won"],
			TotalMatchesWonBaggage:                    statsMap["total_matches_won_baggage"],
			TotalMatchesWonBank:                       statsMap["total_matches_won_bank"],
			TotalMatchesWonLake:                       statsMap["total_matches_won_lake"],
			TotalMatchesWonSafehouse:                  statsMap["total_matches_won_safehouse"],
			TotalMatchesWonShoots:                     statsMap["total_matches_won_shoots"],
			TotalMatchesWonStmarc:                     statsMap["total_matches_won_stmarc"],
			TotalMatchesWonSugarcane:                  statsMap["total_matches_won_sugarcane"],
			TotalMatchesWonTrain:                      statsMap["total_matches_won_train"],
			TotalMoneyEarned:                          statsMap["total_money_earned"],
			TotalMvps:                                 statsMap["total_mvps"],
			TotalPlantedBombs:                         statsMap["total_planted_bombs"],
			TotalProgressiveMatchesWon:                statsMap["total_progressive_matches_won"],
			TotalRescuedHostages:                      statsMap["total_rescued_hostages"],
			TotalRevenges:                             statsMap["total_revenges"],
			TotalRoundsMapArBaggage:                   statsMap["total_rounds_map_ar_baggage"],
			TotalRoundsMapArMonastery:                 statsMap["total_rounds_map_ar_monastery"],
			TotalRoundsMapArShoots:                    statsMap["total_rounds_map_ar_shoots"],
			TotalRoundsMapCsAssault:                   statsMap["total_rounds_map_cs_assault"],
			TotalRoundsMapCsItaly:                     statsMap["total_rounds_map_cs_italy"],
			TotalRoundsMapCsMilitia:                   statsMap["total_rounds_map_cs_militia"],
			TotalRoundsMapCsOffice:                    statsMap["total_rounds_map_cs_office"],
			TotalRoundsMapDeAztec:                     statsMap["total_rounds_map_de_aztec"],
			TotalRoundsMapDeBank:                      statsMap["total_rounds_map_de_bank"],
			TotalRoundsMapDeCbble:                     statsMap["total_rounds_map_de_cbble"],
			TotalRoundsMapDeDust:                      statsMap["total_rounds_map_de_dust"],
			TotalRoundsMapDeDust2:                     statsMap["total_rounds_map_de_dust_2"],
			TotalRoundsMapDeInferno:                   statsMap["total_rounds_map_de_inferno"],
			TotalRoundsMapDeLake:                      statsMap["total_rounds_map_de_lake"],
			TotalRoundsMapDeNuke:                      statsMap["total_rounds_map_de_nuke"],
			TotalRoundsMapDeSafehouse:                 statsMap["total_rounds_map_de_safehouse"],
			TotalRoundsMapDeShorttrain:                statsMap["total_rounds_map_de_shorttrain"],
			TotalRoundsMapDeStmarc:                    statsMap["total_rounds_map_de_stmarc"],
			TotalRoundsMapDeSugarcane:                 statsMap["total_rounds_map_de_sugarcane"],
			TotalRoundsMapDeTrain:                     statsMap["total_rounds_map_de_train"],
			TotalRoundsMapDeVertigo:                   statsMap["total_rounds_map_de_vertigo"],
			TotalRoundsPlayed:                         statsMap["total_rounds_played"],
			TotalShotsAk47:                            statsMap["total_shots_ak47"],
			TotalShotsAug:                             statsMap["total_shots_aug"],
			TotalShotsAwp:                             statsMap["total_shots_awp"],
			TotalShotsBizon:                           statsMap["total_shots_bizon"],
			TotalShotsDeagle:                          statsMap["total_shots_deagle"],
			TotalShotsElite:                           statsMap["total_shots_elite"],
			TotalShotsFamas:                           statsMap["total_shots_famas"],
			TotalShotsFired:                           statsMap["total_shots_fired"],
			TotalShotsFiveseven:                       statsMap["total_shots_fiveseven"],
			TotalShotsG3sg1:                           statsMap["total_shots_g3sg1"],
			TotalShotsGalilar:                         statsMap["total_shots_galilar"],
			TotalShotsGlock:                           statsMap["total_shots_glock"],
			TotalShotsHit:                             statsMap["total_shots_hit"],
			TotalShotsHkp2000:                         statsMap["total_shots_hkp2000"],
			TotalShotsM249:                            statsMap["total_shots_m249"],
			TotalShotsM4a1:                            statsMap["total_shots_m4a1"],
			TotalShotsMac10:                           statsMap["total_shots_mac10"],
			TotalShotsMag7:                            statsMap["total_shots_mag7"],
			TotalShotsMp7:                             statsMap["total_shots_mp7"],
			TotalShotsMp9:                             statsMap["total_shots_mp9"],
			TotalShotsNegev:                           statsMap["total_shots_negev"],
			TotalShotsNova:                            statsMap["total_shots_nova"],
			TotalShotsP250:                            statsMap["total_shots_p250"],
			TotalShotsP90:                             statsMap["total_shots_p90"],
			TotalShotsSawedoff:                        statsMap["total_shots_sawedoff"],
			TotalShotsScar20:                          statsMap["total_shots_scar20"],
			TotalShotsSg556:                           statsMap["total_shots_sg556"],
			TotalShotsSsg08:                           statsMap["total_shots_ssg08"],
			TotalShotsTaser:                           statsMap["total_shots_taser"],
			TotalShotsTec9:                            statsMap["total_shots_tec9"],
			TotalShotsUmp45:                           statsMap["total_shots_ump45"],
			TotalShotsXm1014:                          statsMap["total_shots_xm1014"],
			TotalTRDefusedBombs:                       statsMap["total_tr_defused_bombs"],
			TotalTRPlantedBombs:                       statsMap["total_tr_planted_bombs"],
			TotalTimePlayed:                           statsMap["total_time_played"],
			TotalTrbombMatchesWon:                     statsMap["total_trbomb_matches_won"],
			TotalWeaponsDonated:                       statsMap["total_weapons_donated"],
			TotalWins:                                 statsMap["total_wins"],
			TotalWinsMapArBaggage:                     statsMap["total_wins_map_ar_baggage"],
			TotalWinsMapArMonastery:                   statsMap["total_wins_map_ar_monastery"],
			TotalWinsMapArShoots:                      statsMap["total_wins_map_ar_shoots"],
			TotalWinsMapCsAssault:                     statsMap["total_wins_map_cs_assault"],
			TotalWinsMapCsItaly:                       statsMap["total_wins_map_cs_italy"],
			TotalWinsMapCsMilitia:                     statsMap["total_wins_map_cs_militia"],
			TotalWinsMapCsOffice:                      statsMap["total_wins_map_cs_office"],
			TotalWinsMapDeAztec:                       statsMap["total_wins_map_de_aztec"],
			TotalWinsMapDeBank:                        statsMap["total_wins_map_de_bank"],
			TotalWinsMapDeCbble:                       statsMap["total_wins_map_de_cbble"],
			TotalWinsMapDeDust:                        statsMap["total_wins_map_de_dust"],
			TotalWinsMapDeDust2:                       statsMap["total_wins_map_de_dust2"],
			TotalWinsMapDeHouse:                       statsMap["total_wins_map_de_house"],
			TotalWinsMapDeInferno:                     statsMap["total_wins_map_de_inferno"],
			TotalWinsMapDeLake:                        statsMap["total_wins_map_de_lake"],
			TotalWinsMapDeNuke:                        statsMap["total_wins_map_de_nuke"],
			TotalWinsMapDeSafehouse:                   statsMap["total_wins_map_de_safehouse"],
			TotalWinsMapDeShorttrain:                  statsMap["total_wins_map_de_shorttrain"],
			TotalWinsMapDeStmarc:                      statsMap["total_wins_map_de_stmarc"],
			TotalWinsMapDeSugarcane:                   statsMap["total_wins_map_de_sugarcane"],
			TotalWinsMapDeTrain:                       statsMap["total_wins_map_de_train"],
			TotalWinsMapDeVertigo:                     statsMap["total_wins_map_de_vertigo"],
			TotalWinsPistolround:                      statsMap["total_wins_pistolround"],
		},
		Archivements: GameArchievements{
			//TODO implement achievements, if ever needed
		},
	}, nil

}
