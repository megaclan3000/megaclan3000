package steamclient

import common "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"

// GameStats holds the players stats data from the steam API
// endpoint UserStatsForGame
type GameStats struct {
	SteamID                                   uint64
	GILessonCsgoInstrExplainInspect           int
	GILessonBombSitesA                        int
	GILessonBombSitesB                        int
	GILessonCsgoCycleWeaponsKb                int
	GILessonCsgoHostageLeadToHrz              int
	GILessonCsgoInstrExplainBombCarrier       int
	GILessonCsgoInstrExplainBuyarmor          int
	GILessonCsgoInstrExplainBuymenu           int
	GILessonCsgoInstrExplainFollowBomber      int
	GILessonCsgoInstrExplainPickupBomb        int
	GILessonCsgoInstrExplainPlantBomb         int
	GILessonCsgoInstrExplainPreventBombPickup int
	GILessonCsgoInstrExplainReload            int
	GILessonCsgoInstrExplainZoom              int
	GILessonCsgoInstrRescueZone               int
	GILessonDefusePlantedBomb                 int
	GILessonFindPlantedBomb                   int
	GILessonTrExplainPlantBomb                int
	GILessonVersionNumber                     int
	LastMatchContributionScore                int
	LastMatchCtWins                           int
	LastMatchDamage                           int
	LastMatchDeaths                           int
	LastMatchDominations                      int
	LastMatchFavweaponHits                    int
	LastMatchFavweaponID                      int
	LastMatchFavweaponKills                   int
	LastMatchFavweaponShots                   int
	LastMatchGgContributionScore              int
	LastMatchKills                            int
	LastMatchMaxPlayers                       int
	LastMatchMoneySpent                       int
	LastMatchMvps                             int
	LastMatchRevenges                         int
	LastMatchRounds                           int
	LastMatchTWins                            int
	LastMatchWins                             int
	SteamStatMatchwinscomp                    int
	SteamStatSurvivedz                        int
	SteamStatXpearnedgames                    int
	TotalBrokenWindows                        int
	TotalContributionScore                    int
	TotalDamageDone                           int
	TotalDeaths                               int
	TotalDefusedBombs                         int
	TotalDominationOverkills                  int
	TotalDominations                          int
	TotalGgMatchesPlayed                      int
	TotalGgMatchesWon                         int
	TotalGunGameContributionScore             int
	TotalGunGameRoundsPlayed                  int
	TotalGunGameRoundsWon                     int
	TotalHitsAk47                             int
	TotalHitsAug                              int
	TotalHitsAwp                              int
	TotalHitsBizon                            int
	TotalHitsDeagle                           int
	TotalHitsElite                            int
	TotalHitsFamas                            int
	TotalHitsFiveseven                        int
	TotalHitsG3sg1                            int
	TotalHitsGalilar                          int
	TotalHitsGlock                            int
	TotalHitsHkp2000                          int
	TotalHitsM249                             int
	TotalHitsM4a1                             int
	TotalHitsMac10                            int
	TotalHitsMag7                             int
	TotalHitsMp7                              int
	TotalHitsMp9                              int
	TotalHitsNegev                            int
	TotalHitsNova                             int
	TotalHitsP250                             int
	TotalHitsP90                              int
	TotalHitsSawedoff                         int
	TotalHitsScar20                           int
	TotalHitsSg556                            int
	TotalHitsSsg08                            int
	TotalHitsTec9                             int
	TotalHitsUmp45                            int
	TotalHitsXm1014                           int
	TotalKills                                int
	TotalKillsAgainstZoomedSniper             int
	TotalKillsAk47                            int
	TotalKillsAug                             int
	TotalKillsAwp                             int
	TotalKillsBizon                           int
	TotalKillsDeagle                          int
	TotalKillsElite                           int
	TotalKillsEnemyBlinded                    int
	TotalKillsEnemyWeapon                     int
	TotalKillsFamas                           int
	TotalKillsFiveseven                       int
	TotalKillsG3sg1                           int
	TotalKillsGalilar                         int
	TotalKillsGlock                           int
	TotalKillsHeadshot                        int
	TotalKillsHegrenade                       int
	TotalKillsHkp2000                         int
	TotalKillsKnife                           int
	TotalKillsKnifeFight                      int
	TotalKillsM249                            int
	TotalKillsM4a1                            int
	TotalKillsMac10                           int
	TotalKillsMag7                            int
	TotalKillsMolotov                         int
	TotalKillsMp7                             int
	TotalKillsMp9                             int
	TotalKillsNegev                           int
	TotalKillsNova                            int
	TotalKillsP250                            int
	TotalKillsP90                             int
	TotalKillsSawedoff                        int
	TotalKillsScar20                          int
	TotalKillsSg556                           int
	TotalKillsSsg08                           int
	TotalKillsTaser                           int
	TotalKillsTec9                            int
	TotalKillsUmp45                           int
	TotalKillsXm1014                          int
	TotalMatchesPlayed                        int
	TotalMatchesWon                           int
	TotalMatchesWonBaggage                    int
	TotalMatchesWonBank                       int
	TotalMatchesWonLake                       int
	TotalMatchesWonSafehouse                  int
	TotalMatchesWonShoots                     int
	TotalMatchesWonStmarc                     int
	TotalMatchesWonSugarcane                  int
	TotalMatchesWonTrain                      int
	TotalMoneyEarned                          int
	TotalMvps                                 int
	TotalPlantedBombs                         int
	TotalProgressiveMatchesWon                int
	TotalRescuedHostages                      int
	TotalRevenges                             int
	TotalRoundsMapArBaggage                   int
	TotalRoundsMapArMonastery                 int
	TotalRoundsMapArShoots                    int
	TotalRoundsMapCsAssault                   int
	TotalRoundsMapCsItaly                     int
	TotalRoundsMapCsMilitia                   int
	TotalRoundsMapCsOffice                    int
	TotalRoundsMapDeAztec                     int
	TotalRoundsMapDeBank                      int
	TotalRoundsMapDeCbble                     int
	TotalRoundsMapDeDust                      int
	TotalRoundsMapDeDust2                     int
	TotalRoundsMapDeInferno                   int
	TotalRoundsMapDeLake                      int
	TotalRoundsMapDeNuke                      int
	TotalRoundsMapDeSafehouse                 int
	TotalRoundsMapDeShorttrain                int
	TotalRoundsMapDeStmarc                    int
	TotalRoundsMapDeSugarcane                 int
	TotalRoundsMapDeTrain                     int
	TotalRoundsMapDeVertigo                   int
	TotalRoundsPlayed                         int
	TotalShotsAk47                            int
	TotalShotsAug                             int
	TotalShotsAwp                             int
	TotalShotsBizon                           int
	TotalShotsDeagle                          int
	TotalShotsElite                           int
	TotalShotsFamas                           int
	TotalShotsFired                           int
	TotalShotsFiveseven                       int
	TotalShotsG3sg1                           int
	TotalShotsGalilar                         int
	TotalShotsGlock                           int
	TotalShotsHit                             int
	TotalShotsHkp2000                         int
	TotalShotsM249                            int
	TotalShotsM4a1                            int
	TotalShotsMac10                           int
	TotalShotsMag7                            int
	TotalShotsMp7                             int
	TotalShotsMp9                             int
	TotalShotsNegev                           int
	TotalShotsNova                            int
	TotalShotsP250                            int
	TotalShotsP90                             int
	TotalShotsSawedoff                        int
	TotalShotsScar20                          int
	TotalShotsSg556                           int
	TotalShotsSsg08                           int
	TotalShotsTaser                           int
	TotalShotsTec9                            int
	TotalShotsUmp45                           int
	TotalShotsXm1014                          int
	TotalTRDefusedBombs                       int
	TotalTRPlantedBombs                       int
	TotalTimePlayed                           int
	TotalTrbombMatchesWon                     int
	TotalWeaponsDonated                       int
	TotalWins                                 int
	TotalWinsMapArBaggage                     int
	TotalWinsMapArMonastery                   int
	TotalWinsMapArShoots                      int
	TotalWinsMapCsAssault                     int
	TotalWinsMapCsItaly                       int
	TotalWinsMapCsMilitia                     int
	TotalWinsMapCsOffice                      int
	TotalWinsMapDeAztec                       int
	TotalWinsMapDeBank                        int
	TotalWinsMapDeCbble                       int
	TotalWinsMapDeDust                        int
	TotalWinsMapDeDust2                       int
	TotalWinsMapDeHouse                       int
	TotalWinsMapDeInferno                     int
	TotalWinsMapDeLake                        int
	TotalWinsMapDeNuke                        int
	TotalWinsMapDeSafehouse                   int
	TotalWinsMapDeShorttrain                  int
	TotalWinsMapDeStmarc                      int
	TotalWinsMapDeSugarcane                   int
	TotalWinsMapDeTrain                       int
	TotalWinsMapDeVertigo                     int
	TotalWinsPistolround                      int
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
