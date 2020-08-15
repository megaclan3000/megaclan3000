package steamclient

import common "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"

// GameStats holds the players stats data from the steam API
// endpoint UserStatsForGame
type GameStats struct {
	SteamID                                   uint64 `db:"steamid"`
	GILessonCsgoInstrExplainInspect           int    `db:"gi_lesson_csgo_instr_explain_inspect"`
	GILessonBombSitesA                        int    `db:"gi_lesson_bomb_sites_a"`
	GILessonBombSitesB                        int    `db:"gi_lesson_bomb_sites_b"`
	GILessonCsgoCycleWeaponsKb                int    `db:"gi_lesson_csgo_cycle_weapons_kb"`
	GILessonCsgoHostageLeadToHrz              int    `db:"gi_lesson_csgo_hostage_lead_to_hrz"`
	GILessonCsgoInstrExplainBombCarrier       int    `db:"gi_lesson_csgo_instr_explain_bomb_carrier"`
	GILessonCsgoInstrExplainBuyarmor          int    `db:"gi_lesson_csgo_instr_explain_buyarmor"`
	GILessonCsgoInstrExplainBuymenu           int    `db:"gi_lesson_csgo_instr_explain_buymenu"`
	GILessonCsgoInstrExplainFollowBomber      int    `db:"gi_lesson_csgo_instr_explain_follow_bomber"`
	GILessonCsgoInstrExplainPickupBomb        int    `db:"gi_lesson_csgo_instr_explain_pickup_bomb"`
	GILessonCsgoInstrExplainPlantBomb         int    `db:"gi_lesson_csgo_instr_explain_plant_bomb"`
	GILessonCsgoInstrExplainPreventBombPickup int    `db:"gi_lesson_csgo_instr_explain_prevent_bomb_pickup"`
	GILessonCsgoInstrExplainReload            int    `db:"gi_lesson_csgo_instr_explain_reload"`
	GILessonCsgoInstrExplainZoom              int    `db:"gi_lesson_csgo_instr_explain_zoom"`
	GILessonCsgoInstrRescueZone               int    `db:"gi_lesson_csgo_instr_rescue_zone"`
	GILessonDefusePlantedBomb                 int    `db:"gi_lesson_defuse_planted_bomb"`
	GILessonFindPlantedBomb                   int    `db:"gi_lesson_find_planted_bomb"`
	GILessonTrExplainPlantBomb                int    `db:"gi_lesson_tr_explain_plant_bomb"`
	GILessonVersionNumber                     int    `db:"gi_lesson_version_number"`
	LastMatchContributionScore                int    `db:"last_match_contribution_score"`
	LastMatchCtWins                           int    `db:"last_match_ct_wins"`
	LastMatchDamage                           int    `db:"last_match_damage"`
	LastMatchDeaths                           int    `db:"last_match_deaths"`
	LastMatchDominations                      int    `db:"last_match_dominations"`
	LastMatchFavweaponHits                    int    `db:"last_match_favweapon_hits"`
	LastMatchFavweaponID                      int    `db:"last_match_favweapon_id"`
	LastMatchFavweaponKills                   int    `db:"last_match_favweapon_kills"`
	LastMatchFavweaponShots                   int    `db:"last_match_favweapon_shots"`
	LastMatchGgContributionScore              int    `db:"last_match_gg_contribution_score"`
	LastMatchKills                            int    `db:"last_match_kills"`
	LastMatchMaxPlayers                       int    `db:"last_match_max_players"`
	LastMatchMoneySpent                       int    `db:"last_match_money_spent"`
	LastMatchMvps                             int    `db:"last_match_mvps"`
	LastMatchRevenges                         int    `db:"last_match_revenges"`
	LastMatchRounds                           int    `db:"last_match_rounds"`
	LastMatchTWins                            int    `db:"last_match_t_wins"`
	LastMatchWins                             int    `db:"last_match_wins"`
	SteamStatMatchwinscomp                    int    `db:"steam_stat_matchwinscomp"`
	SteamStatSurvivedz                        int    `db:"steam_stat_survivedz"`
	SteamStatXpearnedgames                    int    `db:"steam_stat_xpearnedgames"`
	TotalBrokenWindows                        int    `db:"total_broken_windows"`
	TotalContributionScore                    int    `db:"total_contribution_score"`
	TotalDamageDone                           int    `db:"total_damage_done"`
	TotalDeaths                               int    `db:"total_deaths"`
	TotalDefusedBombs                         int    `db:"total_defused_bombs"`
	TotalDominationOverkills                  int    `db:"total_domination_overkills"`
	TotalDominations                          int    `db:"total_dominations"`
	TotalGgMatchesPlayed                      int    `db:"total_gg_matches_played"`
	TotalGgMatchesWon                         int    `db:"total_gg_matches_won"`
	TotalGunGameContributionScore             int    `db:"total_gun_game_contribution_score"`
	TotalGunGameRoundsPlayed                  int    `db:"total_gun_game_rounds_played"`
	TotalGunGameRoundsWon                     int    `db:"total_gun_game_rounds_won"`
	TotalHitsAk47                             int    `db:"total_hits_ak47"`
	TotalHitsAug                              int    `db:"total_hits_aug"`
	TotalHitsAwp                              int    `db:"total_hits_awp"`
	TotalHitsBizon                            int    `db:"total_hits_bizon"`
	TotalHitsDeagle                           int    `db:"total_hits_deagle"`
	TotalHitsElite                            int    `db:"total_hits_elite"`
	TotalHitsFamas                            int    `db:"total_hits_famas"`
	TotalHitsFiveseven                        int    `db:"total_hits_fiveseven"`
	TotalHitsG3sg1                            int    `db:"total_hits_g3sg1"`
	TotalHitsGalilar                          int    `db:"total_hits_galilar"`
	TotalHitsGlock                            int    `db:"total_hits_glock"`
	TotalHitsHkp2000                          int    `db:"total_hits_hkp2000"`
	TotalHitsM249                             int    `db:"total_hits_m249"`
	TotalHitsM4a1                             int    `db:"total_hits_m4a1"`
	TotalHitsMac10                            int    `db:"total_hits_mac10"`
	TotalHitsMag7                             int    `db:"total_hits_mag7"`
	TotalHitsMp7                              int    `db:"total_hits_mp7"`
	TotalHitsMp9                              int    `db:"total_hits_mp9"`
	TotalHitsNegev                            int    `db:"total_hits_negev"`
	TotalHitsNova                             int    `db:"total_hits_nova"`
	TotalHitsP250                             int    `db:"total_hits_p250"`
	TotalHitsP90                              int    `db:"total_hits_p90"`
	TotalHitsSawedoff                         int    `db:"total_hits_sawedoff"`
	TotalHitsScar20                           int    `db:"total_hits_scar20"`
	TotalHitsSg556                            int    `db:"total_hits_sg556"`
	TotalHitsSsg08                            int    `db:"total_hits_ssg08"`
	TotalHitsTec9                             int    `db:"total_hits_tec9"`
	TotalHitsUmp45                            int    `db:"total_hits_ump45"`
	TotalHitsXm1014                           int    `db:"total_hits_xm1014"`
	TotalKills                                int    `db:"total_kills"`
	TotalKillsAgainstZoomedSniper             int    `db:"total_kills_against_zoomed_sniper"`
	TotalKillsAk47                            int    `db:"total_kills_ak47"`
	TotalKillsAug                             int    `db:"total_kills_aug"`
	TotalKillsAwp                             int    `db:"total_kills_awp"`
	TotalKillsBizon                           int    `db:"total_kills_bizon"`
	TotalKillsDeagle                          int    `db:"total_kills_deagle"`
	TotalKillsElite                           int    `db:"total_kills_elite"`
	TotalKillsEnemyBlinded                    int    `db:"total_kills_enemy_blinded"`
	TotalKillsEnemyWeapon                     int    `db:"total_kills_enemy_weapon"`
	TotalKillsFamas                           int    `db:"total_kills_famas"`
	TotalKillsFiveseven                       int    `db:"total_kills_fiveseven"`
	TotalKillsG3sg1                           int    `db:"total_kills_g3sg1"`
	TotalKillsGalilar                         int    `db:"total_kills_galilar"`
	TotalKillsGlock                           int    `db:"total_kills_glock"`
	TotalKillsHeadshot                        int    `db:"total_kills_headshot"`
	TotalKillsHegrenade                       int    `db:"total_kills_hegrenade"`
	TotalKillsHkp2000                         int    `db:"total_kills_hkp2000"`
	TotalKillsKnife                           int    `db:"total_kills_knife"`
	TotalKillsKnifeFight                      int    `db:"total_kills_knife_fight"`
	TotalKillsM249                            int    `db:"total_kills_m249"`
	TotalKillsM4a1                            int    `db:"total_kills_m4a1"`
	TotalKillsMac10                           int    `db:"total_kills_mac10"`
	TotalKillsMag7                            int    `db:"total_kills_mag7"`
	TotalKillsMolotov                         int    `db:"total_kills_molotov"`
	TotalKillsMp7                             int    `db:"total_kills_mp7"`
	TotalKillsMp9                             int    `db:"total_kills_mp9"`
	TotalKillsNegev                           int    `db:"total_kills_negev"`
	TotalKillsNova                            int    `db:"total_kills_nova"`
	TotalKillsP250                            int    `db:"total_kills_p250"`
	TotalKillsP90                             int    `db:"total_kills_p90"`
	TotalKillsSawedoff                        int    `db:"total_kills_sawedoff"`
	TotalKillsScar20                          int    `db:"total_kills_scar20"`
	TotalKillsSg556                           int    `db:"total_kills_sg556"`
	TotalKillsSsg08                           int    `db:"total_kills_ssg08"`
	TotalKillsTaser                           int    `db:"total_kills_taser"`
	TotalKillsTec9                            int    `db:"total_kills_tec9"`
	TotalKillsUmp45                           int    `db:"total_kills_ump45"`
	TotalKillsXm1014                          int    `db:"total_kills_xm1014"`
	TotalMatchesPlayed                        int    `db:"total_matches_played"`
	TotalMatchesWon                           int    `db:"total_matches_won"`
	TotalMatchesWonBaggage                    int    `db:"total_matches_won_baggage"`
	TotalMatchesWonBank                       int    `db:"total_matches_won_bank"`
	TotalMatchesWonLake                       int    `db:"total_matches_won_lake"`
	TotalMatchesWonSafehouse                  int    `db:"total_matches_won_safehouse"`
	TotalMatchesWonShoots                     int    `db:"total_matches_won_shoots"`
	TotalMatchesWonStmarc                     int    `db:"total_matches_won_stmarc"`
	TotalMatchesWonSugarcane                  int    `db:"total_matches_won_sugarcane"`
	TotalMatchesWonTrain                      int    `db:"total_matches_won_train"`
	TotalMoneyEarned                          int    `db:"total_money_earned"`
	TotalMvps                                 int    `db:"total_mvps"`
	TotalPlantedBombs                         int    `db:"total_planted_bombs"`
	TotalProgressiveMatchesWon                int    `db:"total_progressive_matches_won"`
	TotalRescuedHostages                      int    `db:"total_rescued_hostages"`
	TotalRevenges                             int    `db:"total_revenges"`
	TotalRoundsMapArBaggage                   int    `db:"total_rounds_map_ar_baggage"`
	TotalRoundsMapArMonastery                 int    `db:"total_rounds_map_ar_monastery"`
	TotalRoundsMapArShoots                    int    `db:"total_rounds_map_ar_shoots"`
	TotalRoundsMapCsAssault                   int    `db:"total_rounds_map_cs_assault"`
	TotalRoundsMapCsItaly                     int    `db:"total_rounds_map_cs_italy"`
	TotalRoundsMapCsMilitia                   int    `db:"total_rounds_map_cs_militia"`
	TotalRoundsMapCsOffice                    int    `db:"total_rounds_map_cs_office"`
	TotalRoundsMapDeAztec                     int    `db:"total_rounds_map_de_aztec"`
	TotalRoundsMapDeBank                      int    `db:"total_rounds_map_de_bank"`
	TotalRoundsMapDeCbble                     int    `db:"total_rounds_map_de_cbble"`
	TotalRoundsMapDeDust                      int    `db:"total_rounds_map_de_dust"`
	TotalRoundsMapDeDust2                     int    `db:"total_rounds_map_de_dust_2"`
	TotalRoundsMapDeInferno                   int    `db:"total_rounds_map_de_inferno"`
	TotalRoundsMapDeLake                      int    `db:"total_rounds_map_de_lake"`
	TotalRoundsMapDeNuke                      int    `db:"total_rounds_map_de_nuke"`
	TotalRoundsMapDeSafehouse                 int    `db:"total_rounds_map_de_safehouse"`
	TotalRoundsMapDeShorttrain                int    `db:"total_rounds_map_de_shorttrain"`
	TotalRoundsMapDeStmarc                    int    `db:"total_rounds_map_de_stmarc"`
	TotalRoundsMapDeSugarcane                 int    `db:"total_rounds_map_de_sugarcane"`
	TotalRoundsMapDeTrain                     int    `db:"total_rounds_map_de_train"`
	TotalRoundsMapDeVertigo                   int    `db:"total_rounds_map_de_vertigo"`
	TotalRoundsPlayed                         int    `db:"total_rounds_played"`
	TotalShotsAk47                            int    `db:"total_shots_ak47"`
	TotalShotsAug                             int    `db:"total_shots_aug"`
	TotalShotsAwp                             int    `db:"total_shots_awp"`
	TotalShotsBizon                           int    `db:"total_shots_bizon"`
	TotalShotsDeagle                          int    `db:"total_shots_deagle"`
	TotalShotsElite                           int    `db:"total_shots_elite"`
	TotalShotsFamas                           int    `db:"total_shots_famas"`
	TotalShotsFired                           int    `db:"total_shots_fired"`
	TotalShotsFiveseven                       int    `db:"total_shots_fiveseven"`
	TotalShotsG3sg1                           int    `db:"total_shots_g3sg1"`
	TotalShotsGalilar                         int    `db:"total_shots_galilar"`
	TotalShotsGlock                           int    `db:"total_shots_glock"`
	TotalShotsHit                             int    `db:"total_shots_hit"`
	TotalShotsHkp2000                         int    `db:"total_shots_hkp2000"`
	TotalShotsM249                            int    `db:"total_shots_m249"`
	TotalShotsM4a1                            int    `db:"total_shots_m4a1"`
	TotalShotsMac10                           int    `db:"total_shots_mac10"`
	TotalShotsMag7                            int    `db:"total_shots_mag7"`
	TotalShotsMp7                             int    `db:"total_shots_mp7"`
	TotalShotsMp9                             int    `db:"total_shots_mp9"`
	TotalShotsNegev                           int    `db:"total_shots_negev"`
	TotalShotsNova                            int    `db:"total_shots_nova"`
	TotalShotsP250                            int    `db:"total_shots_p250"`
	TotalShotsP90                             int    `db:"total_shots_p90"`
	TotalShotsSawedoff                        int    `db:"total_shots_sawedoff"`
	TotalShotsScar20                          int    `db:"total_shots_scar20"`
	TotalShotsSg556                           int    `db:"total_shots_sg556"`
	TotalShotsSsg08                           int    `db:"total_shots_ssg08"`
	TotalShotsTaser                           int    `db:"total_shots_taser"`
	TotalShotsTec9                            int    `db:"total_shots_tec9"`
	TotalShotsUmp45                           int    `db:"total_shots_ump45"`
	TotalShotsXm1014                          int    `db:"total_shots_xm1014"`
	TotalTRDefusedBombs                       int    `db:"total_tr_defused_bombs"`
	TotalTRPlantedBombs                       int    `db:"total_tr_planted_bombs"`
	TotalTimePlayed                           int    `db:"total_time_played"`
	TotalTrbombMatchesWon                     int    `db:"total_trbomb_matches_won"`
	TotalWeaponsDonated                       int    `db:"total_weapons_donated"`
	TotalWins                                 int    `db:"total_wins"`
	TotalWinsMapArBaggage                     int    `db:"total_wins_map_ar_baggage"`
	TotalWinsMapArMonastery                   int    `db:"total_wins_map_ar_monastery"`
	TotalWinsMapArShoots                      int    `db:"total_wins_map_ar_shoots"`
	TotalWinsMapCsAssault                     int    `db:"total_wins_map_cs_assault"`
	TotalWinsMapCsItaly                       int    `db:"total_wins_map_cs_italy"`
	TotalWinsMapCsMilitia                     int    `db:"total_wins_map_cs_militia"`
	TotalWinsMapCsOffice                      int    `db:"total_wins_map_cs_office"`
	TotalWinsMapDeAztec                       int    `db:"total_wins_map_de_aztec"`
	TotalWinsMapDeBank                        int    `db:"total_wins_map_de_bank"`
	TotalWinsMapDeCbble                       int    `db:"total_wins_map_de_cbble"`
	TotalWinsMapDeDust                        int    `db:"total_wins_map_de_dust"`
	TotalWinsMapDeDust2                       int    `db:"total_wins_map_de_dust_2"`
	TotalWinsMapDeHouse                       int    `db:"total_wins_map_de_house"`
	TotalWinsMapDeInferno                     int    `db:"total_wins_map_de_inferno"`
	TotalWinsMapDeLake                        int    `db:"total_wins_map_de_lake"`
	TotalWinsMapDeNuke                        int    `db:"total_wins_map_de_nuke"`
	TotalWinsMapDeSafehouse                   int    `db:"total_wins_map_de_safehouse"`
	TotalWinsMapDeShorttrain                  int    `db:"total_wins_map_de_shorttrain"`
	TotalWinsMapDeStmarc                      int    `db:"total_wins_map_de_stmarc"`
	TotalWinsMapDeSugarcane                   int    `db:"total_wins_map_de_sugarcane"`
	TotalWinsMapDeTrain                       int    `db:"total_wins_map_de_train"`
	TotalWinsMapDeVertigo                     int    `db:"total_wins_map_de_vertigo"`
	TotalWinsPistolround                      int    `db:"total_wins_pistolround"`
}

type weaponstat struct {
	Weapon common.EquipmentType
	Hits   int
	Shots  int
	Kills  int
}

func (gs GameStats) WeaponStats() []weaponstat {

	ret := []weaponstat{

		{
			Weapon: common.EqAK47,
			Hits:   gs.TotalHitsAk47,
			Shots:  gs.TotalShotsAk47,
			Kills:  gs.TotalKillsAk47,
		},
		{
			Weapon: common.EqAUG,
			Hits:   gs.TotalHitsAug,
			Shots:  gs.TotalShotsAug,
			Kills:  gs.TotalShotsAug,
		},
		{
			Weapon: common.EqAWP,
			Hits:   gs.TotalHitsAwp,
			Shots:  gs.TotalShotsAwp,
			Kills:  gs.TotalKillsAwp,
		},
		{
			Weapon: common.EqBizon,
			Hits:   gs.TotalHitsBizon,
			Shots:  gs.TotalShotsBizon,
			Kills:  gs.TotalKillsBizon,
		},
		{
			Weapon: common.EqDeagle,
			Hits:   gs.TotalHitsDeagle,
			Shots:  gs.TotalShotsDeagle,
			Kills:  gs.TotalKillsDeagle,
		},
		{Weapon: common.EqHE,
			Kills: gs.TotalKillsHegrenade,
		},
		{
			Weapon: common.EqDualBerettas,
			Hits:   gs.TotalHitsElite,
			Shots:  gs.TotalShotsElite,
			Kills:  gs.TotalKillsElite,
		},
		{
			Weapon: common.EqFamas,
			Hits:   gs.TotalHitsFamas,
			Shots:  gs.TotalShotsFamas,
			Kills:  gs.TotalKillsFamas,
		},
		{
			Weapon: common.EqFiveSeven,
			Hits:   gs.TotalHitsFiveseven,
			Shots:  gs.TotalShotsFiveseven,
			Kills:  gs.TotalKillsFiveseven,
		},
		{
			Weapon: common.EqG3SG1,
			Hits:   gs.TotalHitsG3sg1,
			Shots:  gs.TotalShotsG3sg1,
			Kills:  gs.TotalKillsG3sg1,
		},
		{
			Weapon: common.EqGalil,
			Hits:   gs.TotalHitsGalilar,
			Shots:  gs.TotalShotsGalilar,
			Kills:  gs.TotalKillsGalilar,
		},
		{
			Weapon: common.EqGlock,
			Hits:   gs.TotalHitsGlock,
			Shots:  gs.TotalShotsGlock,
			Kills:  gs.TotalKillsGlock,
		},
		{
			Weapon: common.EqKnife,
			Kills:  gs.TotalKillsKnife,
		},
		{
			Weapon: common.EqM249,
			Hits:   gs.TotalHitsM249,
			Shots:  gs.TotalShotsM249,
			Kills:  gs.TotalKillsM249,
		},
		{
			Weapon: common.EqM4A1,
			Hits:   gs.TotalHitsM4a1,
			Shots:  gs.TotalShotsM4a1,
			Kills:  gs.TotalKillsM4a1,
		},
		{
			Weapon: common.EqMP7,
			Hits:   gs.TotalHitsMp7,
			Shots:  gs.TotalShotsMp7,
			Kills:  gs.TotalKillsMp7,
		},
		{
			Weapon: common.EqMP9,
			Hits:   gs.TotalHitsMp9,
			Shots:  gs.TotalShotsMp9,
			Kills:  gs.TotalKillsMp9,
		},
		{
			Weapon: common.EqMac10,
			Hits:   gs.TotalHitsMac10,
			Shots:  gs.TotalShotsMac10,
			Kills:  gs.TotalKillsMac10,
		},
		{
			Weapon: common.EqMag7,
			Hits:   gs.TotalHitsMag7,
			Shots:  gs.TotalShotsMag7,
			Kills:  gs.TotalKillsMag7,
		},
		{
			Weapon: common.EqMolotov,
			Kills:  gs.TotalKillsMolotov,
		},
		{
			Weapon: common.EqNegev,
			Hits:   gs.TotalHitsNegev,
			Shots:  gs.TotalShotsNegev,
			Kills:  gs.TotalKillsNegev,
		},
		{
			Weapon: common.EqNova,
			Hits:   gs.TotalHitsNova,
			Shots:  gs.TotalShotsNova,
			Kills:  gs.TotalKillsNova,
		},
		{
			Weapon: common.EqP2000,
			Hits:   gs.TotalHitsHkp2000,
			Shots:  gs.TotalShotsHkp2000,
			Kills:  gs.TotalKillsHkp2000,
		},
		{
			Weapon: common.EqP250,
			Hits:   gs.TotalHitsP250,
			Shots:  gs.TotalShotsP250,
			Kills:  gs.TotalKillsP250,
		},
		{
			Weapon: common.EqP90,
			Hits:   gs.TotalHitsP90,
			Shots:  gs.TotalShotsP90,
			Kills:  gs.TotalKillsP90,
		},
		{
			Weapon: common.EqSG556,
			Hits:   gs.TotalHitsSg556,
			Shots:  gs.TotalShotsSg556,
			Kills:  gs.TotalKillsSg556,
		},
		{
			Weapon: common.EqSSG08,
			Hits:   gs.TotalHitsSsg08,
			Shots:  gs.TotalShotsSsg08,
			Kills:  gs.TotalKillsSsg08,
		},
		{
			Weapon: common.EqSawedOff,
			Hits:   gs.TotalHitsSawedoff,
			Shots:  gs.TotalShotsSawedoff,
			Kills:  gs.TotalKillsSawedoff,
		},
		{
			Weapon: common.EqScar20,
			Hits:   gs.TotalHitsScar20,
			Shots:  gs.TotalShotsScar20,
			Kills:  gs.TotalKillsScar20,
		},
		{
			Weapon: common.EqTec9,
			Hits:   gs.TotalHitsTec9,
			Shots:  gs.TotalShotsTec9,
			Kills:  gs.TotalKillsTec9,
		},
		{
			Weapon: common.EqUMP,
			Hits:   gs.TotalHitsUmp45,
			Shots:  gs.TotalShotsUmp45,
			Kills:  gs.TotalKillsUmp45,
		},
		{
			Weapon: common.EqXM1014,
			Hits:   gs.TotalHitsXm1014,
			Shots:  gs.TotalShotsXm1014,
			Kills:  gs.TotalKillsXm1014,
		},
		{
			Weapon: common.EqZeus,
			Kills:  gs.TotalKillsTaser,
		},

		// {Weapon: common.EqDecoy,
		// 	Hits:  gs.TotalHitsDecoy,
		// 	Shots: gs.TotalShotsDecoy,
		// 	Kills: gs.TotalKillsDecoy,
		// },
		// {Weapon: common.EqFlash,
		// 	Hits:  gs.TotalHitsFlash,
		// 	Shots: gs.TotalShotsFlash,
		// 	Kills: gs.TotalKillsFlash,
		// },
		// {
		// 	Weapon: common.EqIncendiary,
		// 	Hits:   gs.TotalHitsIncendiary,
		// 	Shots:  gs.TotalShotsIncendiary,
		// 	Kills:  gs.TotalKillsIncendiary,
		// },
		// {Weapon: common.EqM4A4,
		// 	Hits:  gs.TotalHitsM4a4,
		// 	Shots: gs.TotalShotsM4A4,
		// 	Kills: gs.TotalKillsM4a4,
		// },
		// {
		// 	Weapon: common.EqMP5,
		// 	Hits:   gs.TotalHitsMp5
		// 	Shots:  gs.TotalShotsMp5,
		// 	Kills:  gs.TotalKillsMP5,
		// },
		// {Weapon: common.EqRevolver,
		// 	Hits:  gs.TotalHitsRevolver,
		// 	Shots: gs.TotalShotsRevolver,
		// 	Kills: gs.TotalKillsRevolver,
		// },
		// {
		// 	Weapon: common.EqScout,
		// 	Hits:   gs.TotalHitsScout,
		// 	Shots:  gs.TotalShotsScout,
		// 	Kills:  gs.TotalKillsScout,
		// },
		// {Weapon: common.EqSmoke,
		// 	Hits:  gs.TotalHitsSmoke,
		// 	Shots: gs.TotalShotsSmoke,
		// 	Kills: gs.TotalKillsSmoke,
		// },
		// {Weapon: common.EqSwag7,
		// 	Hits:  gs.TotalHits
		// 	Shots: gs.TotalShotsSwag7,
		// 	Kills: gs.TotalKillsSwag7,
		// },
		// {Weapon: common.EqUSP,
		// 	Hits:  gs.TotalHitsUsp,
		// 	Shots: gs.TotalShotsUsp,
		// 	Kills: gs.TotalKillsUsp,
		// },
		// {
		// 	Weapon: common.EqSG553,
		// 	Hits:   gs.TotalHitsSg556,
		// 	Shots:  gs.TotalShotsSg556,
		// 	Kills:  gs.TotalKillsSg556,
		// },

	}

	return ret

}
