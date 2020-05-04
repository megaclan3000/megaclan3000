package database

import (
	"reflect"
	"testing"

	"github.com/pinpox/megaclan3000/internal/steamclient"
)

func TestDataStorage_GetUserStatsForGame(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name    string
		steamID string
		want    steamclient.UserStatsForGame
		wantErr bool
	}{
		{
			name:    "Retrieval of UserStatsForGame from fixtures",
			steamID: "123456789",
			want: steamclient.UserStatsForGame{
				SteamID: "123456789",
				Stats: steamclient.GameStats{
					TotalKills:                                "10",
					TotalDeaths:                               "0",
					TotalTimePlayed:                           "1",
					TotalPlantedBombs:                         "2",
					TotalDefusedBombs:                         "3",
					TotalWins:                                 "4",
					TotalDamageDone:                           "5",
					TotalMoneyEarned:                          "6",
					TotalKillsKnife:                           "7",
					TotalKillsHegrenade:                       "8",
					TotalKillsGlock:                           "9",
					TotalKillsDeagle:                          "10",
					TotalKillsElite:                           "11",
					TotalKillsFiveseven:                       "12",
					TotalKillsXm1014:                          "13",
					TotalKillsMac10:                           "14",
					TotalKillsUmp45:                           "15",
					TotalKillsP90:                             "16",
					TotalKillsAwp:                             "17",
					TotalKillsAk47:                            "18",
					TotalKillsAug:                             "19",
					TotalKillsFamas:                           "20",
					TotalKillsG3sg1:                           "21",
					TotalKillsM249:                            "22",
					TotalKillsHeadshot:                        "23",
					TotalKillsEnemyWeapon:                     "24",
					TotalWinsPistolround:                      "25",
					TotalWinsMapCsAssault:                     "26",
					TotalWinsMapDeDust2:                       "27",
					TotalWinsMapDeInferno:                     "28",
					TotalWinsMapDeTrain:                       "29",
					TotalWeaponsDonated:                       "30",
					TotalKillsEnemyBlinded:                    "31",
					TotalKillsKnifeFight:                      "32",
					TotalKillsAgainstZoomedSniper:             "33",
					TotalDominations:                          "34",
					TotalDominationOverkills:                  "35",
					TotalRevenges:                             "36",
					TotalShotsHit:                             "37",
					TotalShotsFired:                           "38",
					TotalRoundsPlayed:                         "39",
					TotalShotsDeagle:                          "40",
					TotalShotsGlock:                           "41",
					TotalShotsElite:                           "42",
					TotalShotsFiveseven:                       "43",
					TotalShotsAwp:                             "44",
					TotalShotsAk47:                            "45",
					TotalShotsAug:                             "46",
					TotalShotsFamas:                           "47",
					TotalShotsG3sg1:                           "48",
					TotalShotsP90:                             "49",
					TotalShotsMac10:                           "50",
					TotalShotsUmp45:                           "51",
					TotalShotsXm1014:                          "52",
					TotalShotsM249:                            "53",
					TotalHitsDeagle:                           "54",
					TotalHitsGlock:                            "55",
					TotalHitsElite:                            "56",
					TotalHitsFiveseven:                        "57",
					TotalHitsAwp:                              "58",
					TotalHitsAk47:                             "59",
					TotalHitsAug:                              "60",
					TotalHitsFamas:                            "61",
					TotalHitsG3sg1:                            "62",
					TotalHitsP90:                              "63",
					TotalHitsMac10:                            "64",
					TotalHitsUmp45:                            "65",
					TotalHitsXm1014:                           "66",
					TotalHitsM249:                             "67",
					TotalRoundsMapCsAssault:                   "68",
					TotalRoundsMapDeDust2:                     "69",
					TotalRoundsMapDeInferno:                   "70",
					TotalRoundsMapDeTrain:                     "71",
					LastMatchTWins:                            "72",
					LastMatchCtWins:                           "73",
					LastMatchWins:                             "74",
					LastMatchMaxPlayers:                       "75",
					LastMatchKills:                            "76",
					LastMatchDeaths:                           "77",
					LastMatchMvps:                             "78",
					LastMatchFavweaponID:                      "79",
					LastMatchFavweaponShots:                   "80",
					LastMatchFavweaponHits:                    "81",
					LastMatchFavweaponKills:                   "82",
					LastMatchDamage:                           "83",
					LastMatchMoneySpent:                       "84",
					LastMatchDominations:                      "85",
					LastMatchRevenges:                         "86",
					TotalMvps:                                 "87",
					TotalRoundsMapDeLake:                      "88",
					TotalRoundsMapDeSafehouse:                 "89",
					TotalRoundsMapDeBank:                      "90",
					TotalTRPlantedBombs:                       "91",
					TotalGunGameRoundsWon:                     "92",
					TotalGunGameRoundsPlayed:                  "93",
					TotalWinsMapDeBank:                        "94",
					TotalWinsMapDeLake:                        "95",
					TotalMatchesWonBank:                       "96",
					TotalMatchesWon:                           "97",
					TotalMatchesPlayed:                        "98",
					TotalGgMatchesWon:                         "99",
					TotalGgMatchesPlayed:                      "100",
					TotalProgressiveMatchesWon:                "101",
					TotalTrbombMatchesWon:                     "102",
					TotalContributionScore:                    "103",
					LastMatchContributionScore:                "104",
					LastMatchRounds:                           "105",
					TotalKillsHkp2000:                         "106",
					TotalShotsHkp2000:                         "107",
					TotalHitsHkp2000:                          "108",
					TotalHitsP250:                             "109",
					TotalKillsP250:                            "110",
					TotalShotsP250:                            "111",
					TotalKillsSg556:                           "112",
					TotalShotsSg556:                           "113",
					TotalHitsSg556:                            "114",
					TotalHitsScar20:                           "115",
					TotalKillsScar20:                          "116",
					TotalShotsScar20:                          "117",
					TotalShotsSsg08:                           "118",
					TotalHitsSsg08:                            "119",
					TotalKillsSsg08:                           "120",
					TotalShotsMp7:                             "121",
					TotalHitsMp7:                              "122",
					TotalKillsMp7:                             "123",
					TotalKillsMp9:                             "124",
					TotalShotsMp9:                             "125",
					TotalHitsMp9:                              "126",
					TotalHitsNova:                             "127",
					TotalKillsNova:                            "128",
					TotalShotsNova:                            "129",
					TotalHitsNegev:                            "130",
					TotalKillsNegev:                           "131",
					TotalShotsNegev:                           "132",
					TotalShotsSawedoff:                        "133",
					TotalHitsSawedoff:                         "134",
					TotalKillsSawedoff:                        "135",
					TotalShotsBizon:                           "136",
					TotalHitsBizon:                            "137",
					TotalKillsBizon:                           "138",
					TotalKillsTec9:                            "139",
					TotalShotsTec9:                            "140",
					TotalHitsTec9:                             "141",
					TotalShotsMag7:                            "142",
					TotalHitsMag7:                             "143",
					TotalKillsMag7:                            "144",
					TotalGunGameContributionScore:             "145",
					LastMatchGgContributionScore:              "146",
					TotalKillsM4a1:                            "147",
					TotalKillsGalilar:                         "148",
					TotalKillsMolotov:                         "149",
					TotalKillsTaser:                           "150",
					TotalShotsM4a1:                            "151",
					TotalShotsGalilar:                         "152",
					TotalShotsTaser:                           "153",
					TotalHitsM4a1:                             "154",
					TotalHitsGalilar:                          "155",
					TotalMatchesWonTrain:                      "156",
					TotalMatchesWonLake:                       "157",
					GILessonCsgoInstrExplainBuymenu:           "158",
					GILessonCsgoInstrExplainBuyarmor:          "159",
					GILessonCsgoInstrExplainPlantBomb:         "160",
					GILessonCsgoInstrExplainBombCarrier:       "161",
					GILessonBombSitesA:                        "162",
					GILessonDefusePlantedBomb:                 "163",
					GILessonCsgoInstrExplainFollowBomber:      "164",
					GILessonCsgoInstrExplainPickupBomb:        "165",
					GILessonCsgoInstrExplainPreventBombPickup: "166",
					GILessonCsgoCycleWeaponsKb:                "167",
					GILessonCsgoInstrExplainZoom:              "168",
					GILessonCsgoInstrExplainReload:            "169",
					GILessonTrExplainPlantBomb:                "170",
					GILessonBombSitesB:                        "171",
					GILessonVersionNumber:                     "172",
					GILessonFindPlantedBomb:                   "173",
					GILessonCsgoHostageLeadToHrz:              "174",
					GILessonCsgoInstrRescueZone:               "175",
					GILessonCsgoInstrExplainInspect:           "176",
					SteamStatXpearnedgames:                    "177",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			prepareDB()

			got, err := db.GetUserStatsForGame(tt.steamID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.GetUserStatsForGame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataStorage.GetUserStatsForGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataStorage_UpdateUserStatsForGame(t *testing.T) {
	tests := []struct {
		name    string
		stats   steamclient.UserStatsForGame
		want    steamclient.UserStatsForGame
		wantErr bool
	}{
		{
			name: "Update UserStatsForGame for one ID",
			stats: steamclient.UserStatsForGame{
				SteamID:  "123456789",
				GameName: "1",
				Stats: steamclient.GameStats{
					TotalKills: "9999",
				},
			},
			want: steamclient.UserStatsForGame{
				SteamID:  "123456789",
				GameName: "1",
				Stats: steamclient.GameStats{
					TotalKills: "9999",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prepareDB()

			if err := db.UpdateUserStatsForGame(tt.stats); (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.UpdateUserStatsForGame() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(tt.stats, tt.want) {
				t.Errorf("DataStorage.UpdateStatsForgame() = %v, want %v", tt.stats, tt.want)
			}
		})
	}
}
