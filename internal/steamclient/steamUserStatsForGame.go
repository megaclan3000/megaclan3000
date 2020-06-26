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

	// SteamID of the player
	SteamID string `db:"steamid"`

	// Total KD ratio
	TotalKD string `db:"total_kd"`

	// KD ratio of the last match
	LastMatchKD string `db:"last_match_kd"`

	// Total hit ratio
	HitRatio string `db:"hit_ratio"`

	// Total played hours
	PlayedHours string `db:"played_hours"`

	// Total Average damage per round
	TotalADR string `db:"total_adr"`

	// Last match average damage per round
	LastMatchADR string `db:"last_match_adr"`
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

	extra := GameExtras{
		SteamID: data.Playerstats.SteamID,
	}

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

	if totalDamage, err := strconv.ParseFloat(statsMap["total_damage_done"], 64); err == nil {
		if totalRounds, err := strconv.ParseFloat(statsMap["total_rounds_played"], 64); err == nil {
			extra.TotalADR = fmt.Sprintf("%.3f", totalDamage/totalRounds)
		}
	}

	if lastMatchDamage, err := strconv.ParseFloat(statsMap["last_match_damage"], 64); err == nil {
		if lastMatchRounds, err := strconv.ParseFloat(statsMap["last_match_rounds"], 64); err == nil {
			extra.TotalADR = fmt.Sprintf("%.3f", lastMatchDamage/lastMatchRounds)
		}
	}

	return UserStatsForGame{
		SteamID:  data.Playerstats.SteamID,
		GameName: data.Playerstats.GameName,
		Extra:    extra,
		Stats: GameStats{
			SteamID:                                   data.Playerstats.SteamID,
			GILessonCsgoInstrExplainInspect:           nilToZeroString(statsMap["GI.lesson.csgo_instr_explain_inspect"]),
			GILessonBombSitesA:                        nilToZeroString(statsMap["GI.lesson.bomb_sites_a"]),
			GILessonBombSitesB:                        nilToZeroString(statsMap["GI.lesson.bomb_sites_b"]),
			GILessonCsgoCycleWeaponsKb:                nilToZeroString(statsMap["GI.lesson.csgo_cycle_weapons_kb"]),
			GILessonCsgoHostageLeadToHrz:              nilToZeroString(statsMap["GI.lesson.csgo_hostage_lead_to_hrz"]),
			GILessonCsgoInstrExplainBombCarrier:       nilToZeroString(statsMap["GI.lesson.csgo_instr_explain_bomb_carrier"]),
			GILessonCsgoInstrExplainBuyarmor:          nilToZeroString(statsMap["GI.lesson.csgo_instr_explain_buyarmor"]),
			GILessonCsgoInstrExplainBuymenu:           nilToZeroString(statsMap["GI.lesson.csgo_instr_explain_buymenu"]),
			GILessonCsgoInstrExplainFollowBomber:      nilToZeroString(statsMap["GI.lesson.csgo_instr_explain_follow_bomber"]),
			GILessonCsgoInstrExplainPickupBomb:        nilToZeroString(statsMap["GI.lesson.csgo_instr_explain_pickup_bomb"]),
			GILessonCsgoInstrExplainPlantBomb:         nilToZeroString(statsMap["GI.lesson.csgo_instr_explain_plant_bomb"]),
			GILessonCsgoInstrExplainPreventBombPickup: nilToZeroString(statsMap["GI.lesson.csgo_instr_explain_prevent_bomb_pickup"]),
			GILessonCsgoInstrExplainReload:            nilToZeroString(statsMap["GI.lesson.csgo_instr_explain_reload"]),
			GILessonCsgoInstrExplainZoom:              nilToZeroString(statsMap["GI.lesson.csgo_instr_explain_zoom"]),
			GILessonCsgoInstrRescueZone:               nilToZeroString(statsMap["GI.lesson.csgo_instr_rescue_zone"]),
			GILessonDefusePlantedBomb:                 nilToZeroString(statsMap["GI.lesson.defuse_planted_bomb"]),
			GILessonFindPlantedBomb:                   nilToZeroString(statsMap["GI.lesson.find_planted_bomb"]),
			GILessonTrExplainPlantBomb:                nilToZeroString(statsMap["GI.lesson.tr_explain_plant_bomb"]),
			GILessonVersionNumber:                     nilToZeroString(statsMap["GI.lesson.version_number"]),
			LastMatchContributionScore:                nilToZeroString(statsMap["last_match_contribution_score"]),
			LastMatchCtWins:                           nilToZeroString(statsMap["last_match_ct_wins"]),
			LastMatchDamage:                           nilToZeroString(statsMap["last_match_damage"]),
			LastMatchDeaths:                           nilToZeroString(statsMap["last_match_deaths"]),
			LastMatchDominations:                      nilToZeroString(statsMap["last_match_dominations"]),
			LastMatchFavweaponHits:                    nilToZeroString(statsMap["last_match_favweapon_hits"]),
			LastMatchFavweaponID:                      nilToZeroString(statsMap["last_match_favweapon_id"]),
			LastMatchFavweaponKills:                   nilToZeroString(statsMap["last_match_favweapon_kills"]),
			LastMatchFavweaponShots:                   nilToZeroString(statsMap["last_match_favweapon_shots"]),
			LastMatchGgContributionScore:              nilToZeroString(statsMap["last_match_gg_contribution_score"]),
			LastMatchKills:                            nilToZeroString(statsMap["last_match_kills"]),
			LastMatchMaxPlayers:                       nilToZeroString(statsMap["last_match_max_players"]),
			LastMatchMoneySpent:                       nilToZeroString(statsMap["last_match_money_spent"]),
			LastMatchMvps:                             nilToZeroString(statsMap["last_match_mvps"]),
			LastMatchRevenges:                         nilToZeroString(statsMap["last_match_revenges"]),
			LastMatchRounds:                           nilToZeroString(statsMap["last_match_rounds"]),
			LastMatchTWins:                            nilToZeroString(statsMap["last_match_t_wins"]),
			LastMatchWins:                             nilToZeroString(statsMap["last_match_wins"]),
			SteamStatMatchwinscomp:                    nilToZeroString(statsMap["steam_stat_matchwinscomp"]),
			SteamStatSurvivedz:                        nilToZeroString(statsMap["steam_stat_survivedz"]),
			SteamStatXpearnedgames:                    nilToZeroString(statsMap["steam_stat_xpearnedgames"]),
			TotalBrokenWindows:                        nilToZeroString(statsMap["total_broken_windows"]),
			TotalContributionScore:                    nilToZeroString(statsMap["total_contribution_score"]),
			TotalDamageDone:                           nilToZeroString(statsMap["total_damage_done"]),
			TotalDeaths:                               nilToZeroString(statsMap["total_deaths"]),
			TotalDefusedBombs:                         nilToZeroString(statsMap["total_defused_bombs"]),
			TotalDominationOverkills:                  nilToZeroString(statsMap["total_domination_overkills"]),
			TotalDominations:                          nilToZeroString(statsMap["total_dominations"]),
			TotalGgMatchesPlayed:                      nilToZeroString(statsMap["total_gg_matches_played"]),
			TotalGgMatchesWon:                         nilToZeroString(statsMap["total_gg_matches_won"]),
			TotalGunGameContributionScore:             nilToZeroString(statsMap["total_gun_game_contribution_score"]),
			TotalGunGameRoundsPlayed:                  nilToZeroString(statsMap["total_gun_game_rounds_played"]),
			TotalGunGameRoundsWon:                     nilToZeroString(statsMap["total_gun_game_rounds_won"]),
			TotalHitsAk47:                             nilToZeroString(statsMap["total_hits_ak47"]),
			TotalHitsAug:                              nilToZeroString(statsMap["total_hits_aug"]),
			TotalHitsAwp:                              nilToZeroString(statsMap["total_hits_awp"]),
			TotalHitsBizon:                            nilToZeroString(statsMap["total_hits_bizon"]),
			TotalHitsDeagle:                           nilToZeroString(statsMap["total_hits_deagle"]),
			TotalHitsElite:                            nilToZeroString(statsMap["total_hits_elite"]),
			TotalHitsFamas:                            nilToZeroString(statsMap["total_hits_famas"]),
			TotalHitsFiveseven:                        nilToZeroString(statsMap["total_hits_fiveseven"]),
			TotalHitsG3sg1:                            nilToZeroString(statsMap["total_hits_g_3sg_1"]),
			TotalHitsGalilar:                          nilToZeroString(statsMap["total_hits_galilar"]),
			TotalHitsGlock:                            nilToZeroString(statsMap["total_hits_glock"]),
			TotalHitsHkp2000:                          nilToZeroString(statsMap["total_hits_hkp2000"]),
			TotalHitsM249:                             nilToZeroString(statsMap["total_hits_m249"]),
			TotalHitsM4a1:                             nilToZeroString(statsMap["total_hits_m_4a_1"]),
			TotalHitsMac10:                            nilToZeroString(statsMap["total_hits_mac_10"]),
			TotalHitsMag7:                             nilToZeroString(statsMap["total_hits_mag_7"]),
			TotalHitsMp7:                              nilToZeroString(statsMap["total_hits_mp_7"]),
			TotalHitsMp9:                              nilToZeroString(statsMap["total_hits_mp_9"]),
			TotalHitsNegev:                            nilToZeroString(statsMap["total_hits_negev"]),
			TotalHitsNova:                             nilToZeroString(statsMap["total_hits_nova"]),
			TotalHitsP250:                             nilToZeroString(statsMap["total_hits_p250"]),
			TotalHitsP90:                              nilToZeroString(statsMap["total_hits_p90"]),
			TotalHitsSawedoff:                         nilToZeroString(statsMap["total_hits_sawedoff"]),
			TotalHitsScar20:                           nilToZeroString(statsMap["total_hits_scar20"]),
			TotalHitsSg556:                            nilToZeroString(statsMap["total_hits_sg556"]),
			TotalHitsSsg08:                            nilToZeroString(statsMap["total_hits_ssg08"]),
			TotalHitsTec9:                             nilToZeroString(statsMap["total_hits_tec_9"]),
			TotalHitsUmp45:                            nilToZeroString(statsMap["total_hits_ump_45"]),
			TotalHitsXm1014:                           nilToZeroString(statsMap["total_hits_xm_1014"]),
			TotalKills:                                nilToZeroString(statsMap["total_kills"]),
			TotalKillsAgainstZoomedSniper:             nilToZeroString(statsMap["total_kills_against_zoomed_sniper"]),
			TotalKillsAk47:                            nilToZeroString(statsMap["total_kills_ak47"]),
			TotalKillsAug:                             nilToZeroString(statsMap["total_kills_aug"]),
			TotalKillsAwp:                             nilToZeroString(statsMap["total_kills_awp"]),
			TotalKillsBizon:                           nilToZeroString(statsMap["total_kills_bizon"]),
			TotalKillsDeagle:                          nilToZeroString(statsMap["total_kills_deagle"]),
			TotalKillsElite:                           nilToZeroString(statsMap["total_kills_elite"]),
			TotalKillsEnemyBlinded:                    nilToZeroString(statsMap["total_kills_enemy_blinded"]),
			TotalKillsEnemyWeapon:                     nilToZeroString(statsMap["total_kills_enemy_weapon"]),
			TotalKillsFamas:                           nilToZeroString(statsMap["total_kills_famas"]),
			TotalKillsFiveseven:                       nilToZeroString(statsMap["total_kills_fiveseven"]),
			TotalKillsG3sg1:                           nilToZeroString(statsMap["total_kills_g3sg1"]),
			TotalKillsGalilar:                         nilToZeroString(statsMap["total_kills_galilar"]),
			TotalKillsGlock:                           nilToZeroString(statsMap["total_kills_glock"]),
			TotalKillsHeadshot:                        nilToZeroString(statsMap["total_kills_headshot"]),
			TotalKillsHegrenade:                       nilToZeroString(statsMap["total_kills_hegrenade"]),
			TotalKillsHkp2000:                         nilToZeroString(statsMap["total_kills_hkp2000"]),
			TotalKillsKnife:                           nilToZeroString(statsMap["total_kills_knife"]),
			TotalKillsKnifeFight:                      nilToZeroString(statsMap["total_kills_knife_fight"]),
			TotalKillsM249:                            nilToZeroString(statsMap["total_kills_m249"]),
			TotalKillsM4a1:                            nilToZeroString(statsMap["total_kills_m4a1"]),
			TotalKillsMac10:                           nilToZeroString(statsMap["total_kills_mac10"]),
			TotalKillsMag7:                            nilToZeroString(statsMap["total_kills_mag7"]),
			TotalKillsMolotov:                         nilToZeroString(statsMap["total_kills_molotov"]),
			TotalKillsMp7:                             nilToZeroString(statsMap["total_kills_mp7"]),
			TotalKillsMp9:                             nilToZeroString(statsMap["total_kills_mp9"]),
			TotalKillsNegev:                           nilToZeroString(statsMap["total_kills_negev"]),
			TotalKillsNova:                            nilToZeroString(statsMap["total_kills_nova"]),
			TotalKillsP250:                            nilToZeroString(statsMap["total_kills_p250"]),
			TotalKillsP90:                             nilToZeroString(statsMap["total_kills_p90"]),
			TotalKillsSawedoff:                        nilToZeroString(statsMap["total_kills_sawedoff"]),
			TotalKillsScar20:                          nilToZeroString(statsMap["total_kills_scar20"]),
			TotalKillsSg556:                           nilToZeroString(statsMap["total_kills_sg556"]),
			TotalKillsSsg08:                           nilToZeroString(statsMap["total_kills_ssg08"]),
			TotalKillsTaser:                           nilToZeroString(statsMap["total_kills_taser"]),
			TotalKillsTec9:                            nilToZeroString(statsMap["total_kills_tec9"]),
			TotalKillsUmp45:                           nilToZeroString(statsMap["total_kills_ump45"]),
			TotalKillsXm1014:                          nilToZeroString(statsMap["total_kills_xm1014"]),
			TotalMatchesPlayed:                        nilToZeroString(statsMap["total_matches_played"]),
			TotalMatchesWon:                           nilToZeroString(statsMap["total_matches_won"]),
			TotalMatchesWonBaggage:                    nilToZeroString(statsMap["total_matches_won_baggage"]),
			TotalMatchesWonBank:                       nilToZeroString(statsMap["total_matches_won_bank"]),
			TotalMatchesWonLake:                       nilToZeroString(statsMap["total_matches_won_lake"]),
			TotalMatchesWonSafehouse:                  nilToZeroString(statsMap["total_matches_won_safehouse"]),
			TotalMatchesWonShoots:                     nilToZeroString(statsMap["total_matches_won_shoots"]),
			TotalMatchesWonStmarc:                     nilToZeroString(statsMap["total_matches_won_stmarc"]),
			TotalMatchesWonSugarcane:                  nilToZeroString(statsMap["total_matches_won_sugarcane"]),
			TotalMatchesWonTrain:                      nilToZeroString(statsMap["total_matches_won_train"]),
			TotalMoneyEarned:                          nilToZeroString(statsMap["total_money_earned"]),
			TotalMvps:                                 nilToZeroString(statsMap["total_mvps"]),
			TotalPlantedBombs:                         nilToZeroString(statsMap["total_planted_bombs"]),
			TotalProgressiveMatchesWon:                nilToZeroString(statsMap["total_progressive_matches_won"]),
			TotalRescuedHostages:                      nilToZeroString(statsMap["total_rescued_hostages"]),
			TotalRevenges:                             nilToZeroString(statsMap["total_revenges"]),
			TotalRoundsMapArBaggage:                   nilToZeroString(statsMap["total_rounds_map_ar_baggage"]),
			TotalRoundsMapArMonastery:                 nilToZeroString(statsMap["total_rounds_map_ar_monastery"]),
			TotalRoundsMapArShoots:                    nilToZeroString(statsMap["total_rounds_map_ar_shoots"]),
			TotalRoundsMapCsAssault:                   nilToZeroString(statsMap["total_rounds_map_cs_assault"]),
			TotalRoundsMapCsItaly:                     nilToZeroString(statsMap["total_rounds_map_cs_italy"]),
			TotalRoundsMapCsMilitia:                   nilToZeroString(statsMap["total_rounds_map_cs_militia"]),
			TotalRoundsMapCsOffice:                    nilToZeroString(statsMap["total_rounds_map_cs_office"]),
			TotalRoundsMapDeAztec:                     nilToZeroString(statsMap["total_rounds_map_de_aztec"]),
			TotalRoundsMapDeBank:                      nilToZeroString(statsMap["total_rounds_map_de_bank"]),
			TotalRoundsMapDeCbble:                     nilToZeroString(statsMap["total_rounds_map_de_cbble"]),
			TotalRoundsMapDeDust:                      nilToZeroString(statsMap["total_rounds_map_de_dust"]),
			TotalRoundsMapDeDust2:                     nilToZeroString(statsMap["total_rounds_map_de_dust2"]),
			TotalRoundsMapDeInferno:                   nilToZeroString(statsMap["total_rounds_map_de_inferno"]),
			TotalRoundsMapDeLake:                      nilToZeroString(statsMap["total_rounds_map_de_lake"]),
			TotalRoundsMapDeNuke:                      nilToZeroString(statsMap["total_rounds_map_de_nuke"]),
			TotalRoundsMapDeSafehouse:                 nilToZeroString(statsMap["total_rounds_map_de_safehouse"]),
			TotalRoundsMapDeShorttrain:                nilToZeroString(statsMap["total_rounds_map_de_shorttrain"]),
			TotalRoundsMapDeStmarc:                    nilToZeroString(statsMap["total_rounds_map_de_stmarc"]),
			TotalRoundsMapDeSugarcane:                 nilToZeroString(statsMap["total_rounds_map_de_sugarcane"]),
			TotalRoundsMapDeTrain:                     nilToZeroString(statsMap["total_rounds_map_de_train"]),
			TotalRoundsMapDeVertigo:                   nilToZeroString(statsMap["total_rounds_map_de_vertigo"]),
			TotalRoundsPlayed:                         nilToZeroString(statsMap["total_rounds_played"]),
			TotalShotsAk47:                            nilToZeroString(statsMap["total_shots_ak47"]),
			TotalShotsAug:                             nilToZeroString(statsMap["total_shots_aug"]),
			TotalShotsAwp:                             nilToZeroString(statsMap["total_shots_awp"]),
			TotalShotsBizon:                           nilToZeroString(statsMap["total_shots_bizon"]),
			TotalShotsDeagle:                          nilToZeroString(statsMap["total_shots_deagle"]),
			TotalShotsElite:                           nilToZeroString(statsMap["total_shots_elite"]),
			TotalShotsFamas:                           nilToZeroString(statsMap["total_shots_famas"]),
			TotalShotsFired:                           nilToZeroString(statsMap["total_shots_fired"]),
			TotalShotsFiveseven:                       nilToZeroString(statsMap["total_shots_fiveseven"]),
			TotalShotsG3sg1:                           nilToZeroString(statsMap["total_shots_g3sg1"]),
			TotalShotsGalilar:                         nilToZeroString(statsMap["total_shots_galilar"]),
			TotalShotsGlock:                           nilToZeroString(statsMap["total_shots_glock"]),
			TotalShotsHit:                             nilToZeroString(statsMap["total_shots_hit"]),
			TotalShotsHkp2000:                         nilToZeroString(statsMap["total_shots_hkp2000"]),
			TotalShotsM249:                            nilToZeroString(statsMap["total_shots_m249"]),
			TotalShotsM4a1:                            nilToZeroString(statsMap["total_shots_m4a1"]),
			TotalShotsMac10:                           nilToZeroString(statsMap["total_shots_mac10"]),
			TotalShotsMag7:                            nilToZeroString(statsMap["total_shots_mag7"]),
			TotalShotsMp7:                             nilToZeroString(statsMap["total_shots_mp7"]),
			TotalShotsMp9:                             nilToZeroString(statsMap["total_shots_mp9"]),
			TotalShotsNegev:                           nilToZeroString(statsMap["total_shots_negev"]),
			TotalShotsNova:                            nilToZeroString(statsMap["total_shots_nova"]),
			TotalShotsP250:                            nilToZeroString(statsMap["total_shots_p250"]),
			TotalShotsP90:                             nilToZeroString(statsMap["total_shots_p90"]),
			TotalShotsSawedoff:                        nilToZeroString(statsMap["total_shots_sawedoff"]),
			TotalShotsScar20:                          nilToZeroString(statsMap["total_shots_scar20"]),
			TotalShotsSg556:                           nilToZeroString(statsMap["total_shots_sg556"]),
			TotalShotsSsg08:                           nilToZeroString(statsMap["total_shots_ssg08"]),
			TotalShotsTaser:                           nilToZeroString(statsMap["total_shots_taser"]),
			TotalShotsTec9:                            nilToZeroString(statsMap["total_shots_tec9"]),
			TotalShotsUmp45:                           nilToZeroString(statsMap["total_shots_ump45"]),
			TotalShotsXm1014:                          nilToZeroString(statsMap["total_shots_xm1014"]),
			TotalTRDefusedBombs:                       nilToZeroString(statsMap["total_tr_defused_bombs"]),
			TotalTRPlantedBombs:                       nilToZeroString(statsMap["total_tr_planted_bombs"]),
			TotalTimePlayed:                           nilToZeroString(statsMap["total_time_played"]),
			TotalTrbombMatchesWon:                     nilToZeroString(statsMap["total_trbomb_matches_won"]),
			TotalWeaponsDonated:                       nilToZeroString(statsMap["total_weapons_donated"]),
			TotalWins:                                 nilToZeroString(statsMap["total_wins"]),
			TotalWinsMapArBaggage:                     nilToZeroString(statsMap["total_wins_map_ar_baggage"]),
			TotalWinsMapArMonastery:                   nilToZeroString(statsMap["total_wins_map_ar_monastery"]),
			TotalWinsMapArShoots:                      nilToZeroString(statsMap["total_wins_map_ar_shoots"]),
			TotalWinsMapCsAssault:                     nilToZeroString(statsMap["total_wins_map_cs_assault"]),
			TotalWinsMapCsItaly:                       nilToZeroString(statsMap["total_wins_map_cs_italy"]),
			TotalWinsMapCsMilitia:                     nilToZeroString(statsMap["total_wins_map_cs_militia"]),
			TotalWinsMapCsOffice:                      nilToZeroString(statsMap["total_wins_map_cs_office"]),
			TotalWinsMapDeAztec:                       nilToZeroString(statsMap["total_wins_map_de_aztec"]),
			TotalWinsMapDeBank:                        nilToZeroString(statsMap["total_wins_map_de_bank"]),
			TotalWinsMapDeCbble:                       nilToZeroString(statsMap["total_wins_map_de_cbble"]),
			TotalWinsMapDeDust:                        nilToZeroString(statsMap["total_wins_map_de_dust"]),
			TotalWinsMapDeDust2:                       nilToZeroString(statsMap["total_wins_map_de_dust2"]),
			TotalWinsMapDeHouse:                       nilToZeroString(statsMap["total_wins_map_de_house"]),
			TotalWinsMapDeInferno:                     nilToZeroString(statsMap["total_wins_map_de_inferno"]),
			TotalWinsMapDeLake:                        nilToZeroString(statsMap["total_wins_map_de_lake"]),
			TotalWinsMapDeNuke:                        nilToZeroString(statsMap["total_wins_map_de_nuke"]),
			TotalWinsMapDeSafehouse:                   nilToZeroString(statsMap["total_wins_map_de_safehouse"]),
			TotalWinsMapDeShorttrain:                  nilToZeroString(statsMap["total_wins_map_de_shorttrain"]),
			TotalWinsMapDeStmarc:                      nilToZeroString(statsMap["total_wins_map_de_stmarc"]),
			TotalWinsMapDeSugarcane:                   nilToZeroString(statsMap["total_wins_map_de_sugarcane"]),
			TotalWinsMapDeTrain:                       nilToZeroString(statsMap["total_wins_map_de_train"]),
			TotalWinsMapDeVertigo:                     nilToZeroString(statsMap["total_wins_map_de_vertigo"]),
			TotalWinsPistolround:                      nilToZeroString(statsMap["total_wins_pistolround"]),
		},
		Archivements: GameArchievements{
			//TODO implement achievements, if ever needed
		},
	}, nil

}

func nilToZeroString(input string) string {
	if input == "" {
		return "0"
	}
	return input
}
