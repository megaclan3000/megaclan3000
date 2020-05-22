package database

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
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
			name:    "Retrieval of UserStatsForGame from fixtures (ID: all_columns)",
			steamID: "all_columns",
			want: steamclient.UserStatsForGame{
				SteamID: "all_columns",
				Stats: steamclient.GameStats{
					GILessonBombSitesA:                        "1",
					GILessonBombSitesB:                        "2",
					GILessonCsgoCycleWeaponsKb:                "3",
					GILessonCsgoHostageLeadToHrz:              "4",
					GILessonCsgoInstrExplainBombCarrier:       "5",
					GILessonCsgoInstrExplainBuyarmor:          "6",
					GILessonCsgoInstrExplainBuymenu:           "7",
					GILessonCsgoInstrExplainFollowBomber:      "8",
					GILessonCsgoInstrExplainInspect:           "0",
					GILessonCsgoInstrExplainPickupBomb:        "9",
					GILessonCsgoInstrExplainPlantBomb:         "10",
					GILessonCsgoInstrExplainPreventBombPickup: "11",
					GILessonCsgoInstrExplainReload:            "12",
					GILessonCsgoInstrExplainZoom:              "13",
					GILessonCsgoInstrRescueZone:               "14",
					GILessonDefusePlantedBomb:                 "15",
					GILessonFindPlantedBomb:                   "16",
					GILessonTrExplainPlantBomb:                "17",
					GILessonVersionNumber:                     "18",
					LastMatchContributionScore:                "19",
					LastMatchCtWins:                           "20",
					LastMatchDamage:                           "21",
					LastMatchDeaths:                           "22",
					LastMatchDominations:                      "23",
					LastMatchFavweaponID:                      "25",
					LastMatchFavweaponHits:                    "24",
					LastMatchFavweaponKills:                   "26",
					LastMatchFavweaponShots:                   "27",
					LastMatchGgContributionScore:              "28",
					LastMatchKills:                            "29",
					LastMatchMaxPlayers:                       "30",
					LastMatchMoneySpent:                       "31",
					LastMatchMvps:                             "32",
					LastMatchRevenges:                         "33",
					LastMatchRounds:                           "34",
					LastMatchTWins:                            "35",
					LastMatchWins:                             "36",
					SteamStatMatchwinscomp:                    "37",
					SteamStatSurvivedz:                        "38",
					SteamStatXpearnedgames:                    "39",
					TotalBrokenWindows:                        "40",
					TotalContributionScore:                    "41",
					TotalDamageDone:                           "42",
					TotalDeaths:                               "43",
					TotalDefusedBombs:                         "44",
					TotalDominationOverkills:                  "45",
					TotalDominations:                          "46",
					TotalGgMatchesPlayed:                      "47",
					TotalGgMatchesWon:                         "48",
					TotalGunGameContributionScore:             "49",
					TotalGunGameRoundsPlayed:                  "50",
					TotalGunGameRoundsWon:                     "51",
					TotalHitsAk47:                             "52",
					TotalHitsAug:                              "53",
					TotalHitsAwp:                              "54",
					TotalHitsBizon:                            "55",
					TotalHitsDeagle:                           "56",
					TotalHitsElite:                            "57",
					TotalHitsFamas:                            "58",
					TotalHitsFiveseven:                        "59",
					TotalHitsG3sg1:                            "60",
					TotalHitsGalilar:                          "61",
					TotalHitsGlock:                            "62",
					TotalHitsHkp2000:                          "63",
					TotalHitsM249:                             "64",
					TotalHitsM4a1:                             "65",
					TotalHitsMac10:                            "66",
					TotalHitsMag7:                             "67",
					TotalHitsMp7:                              "68",
					TotalHitsMp9:                              "69",
					TotalHitsNegev:                            "70",
					TotalHitsNova:                             "71",
					TotalHitsP250:                             "72",
					TotalHitsP90:                              "73",
					TotalHitsSg556:                            "76",
					TotalHitsSawedoff:                         "74",
					TotalHitsScar20:                           "75",
					TotalHitsSsg08:                            "77",
					TotalHitsTec9:                             "78",
					TotalHitsUmp45:                            "79",
					TotalHitsXm1014:                           "80",
					TotalKills:                                "81",
					TotalKillsAgainstZoomedSniper:             "82",
					TotalKillsAk47:                            "83",
					TotalKillsAug:                             "84",
					TotalKillsAwp:                             "85",
					TotalKillsBizon:                           "86",
					TotalKillsDeagle:                          "87",
					TotalKillsElite:                           "88",
					TotalKillsEnemyBlinded:                    "89",
					TotalKillsEnemyWeapon:                     "90",
					TotalKillsFamas:                           "91",
					TotalKillsFiveseven:                       "92",
					TotalKillsG3sg1:                           "93",
					TotalKillsGalilar:                         "94",
					TotalKillsGlock:                           "95",
					TotalKillsHeadshot:                        "96",
					TotalKillsHegrenade:                       "97",
					TotalKillsHkp2000:                         "98",
					TotalKillsKnife:                           "99",
					TotalKillsKnifeFight:                      "100",
					TotalKillsM249:                            "101",
					TotalKillsM4a1:                            "102",
					TotalKillsMac10:                           "103",
					TotalKillsMag7:                            "104",
					TotalKillsMolotov:                         "105",
					TotalKillsMp7:                             "106",
					TotalKillsMp9:                             "107",
					TotalKillsNegev:                           "108",
					TotalKillsNova:                            "109",
					TotalKillsP250:                            "110",
					TotalKillsP90:                             "111",
					TotalKillsSawedoff:                        "112",
					TotalKillsScar20:                          "113",
					TotalKillsSg556:                           "114",
					TotalKillsSsg08:                           "115",
					TotalKillsTaser:                           "116",
					TotalKillsTec9:                            "117",
					TotalKillsUmp45:                           "118",
					TotalKillsXm1014:                          "119",
					TotalMatchesPlayed:                        "120",
					TotalMatchesWon:                           "121",
					TotalMatchesWonBaggage:                    "122",
					TotalMatchesWonBank:                       "123",
					TotalMatchesWonLake:                       "124",
					TotalMatchesWonSafehouse:                  "125",
					TotalMatchesWonShoots:                     "126",
					TotalMatchesWonStmarc:                     "127",
					TotalMatchesWonSugarcane:                  "128",
					TotalMatchesWonTrain:                      "129",
					TotalMoneyEarned:                          "130",
					TotalMvps:                                 "131",
					TotalPlantedBombs:                         "132",
					TotalProgressiveMatchesWon:                "133",
					TotalRescuedHostages:                      "134",
					TotalRevenges:                             "135",
					TotalRoundsMapArBaggage:                   "136",
					TotalRoundsMapArMonastery:                 "137",
					TotalRoundsMapArShoots:                    "138",
					TotalRoundsMapCsAssault:                   "139",
					TotalRoundsMapCsItaly:                     "140",
					TotalRoundsMapCsMilitia:                   "141",
					TotalRoundsMapCsOffice:                    "142",
					TotalRoundsMapDeAztec:                     "143",
					TotalRoundsMapDeBank:                      "144",
					TotalRoundsMapDeCbble:                     "145",
					TotalRoundsMapDeDust:                      "146",
					TotalRoundsMapDeDust2:                     "147",
					TotalRoundsMapDeInferno:                   "148",
					TotalRoundsMapDeLake:                      "149",
					TotalRoundsMapDeNuke:                      "150",
					TotalRoundsMapDeSafehouse:                 "151",
					TotalRoundsMapDeShorttrain:                "152",
					TotalRoundsMapDeStmarc:                    "153",
					TotalRoundsMapDeSugarcane:                 "154",
					TotalRoundsMapDeTrain:                     "155",
					TotalRoundsMapDeVertigo:                   "156",
					TotalRoundsPlayed:                         "157",
					TotalShotsAk47:                            "158",
					TotalShotsAug:                             "159",
					TotalShotsAwp:                             "160",
					TotalShotsBizon:                           "161",
					TotalShotsDeagle:                          "162",
					TotalShotsElite:                           "163",
					TotalShotsFamas:                           "164",
					TotalShotsFired:                           "165",
					TotalShotsFiveseven:                       "166",
					TotalShotsG3sg1:                           "167",
					TotalShotsGalilar:                         "168",
					TotalShotsGlock:                           "169",
					TotalShotsHit:                             "170",
					TotalShotsHkp2000:                         "171",
					TotalShotsM249:                            "172",
					TotalShotsM4a1:                            "173",
					TotalShotsMac10:                           "174",
					TotalShotsMag7:                            "175",
					TotalShotsMp7:                             "176",
					TotalShotsMp9:                             "177",
					TotalShotsNegev:                           "178",
					TotalShotsNova:                            "179",
					TotalShotsP250:                            "180",
					TotalShotsP90:                             "181",
					TotalShotsSawedoff:                        "182",
					TotalShotsScar20:                          "183",
					TotalShotsSg556:                           "184",
					TotalShotsSsg08:                           "185",
					TotalShotsTaser:                           "186",
					TotalShotsTec9:                            "187",
					TotalShotsUmp45:                           "188",
					TotalShotsXm1014:                          "189",
					TotalTimePlayed:                           "192",
					TotalTrbombMatchesWon:                     "193",
					TotalTRDefusedBombs:                       "190",
					TotalTRPlantedBombs:                       "191",
					TotalWeaponsDonated:                       "194",
					TotalWins:                                 "195",
					TotalWinsMapArBaggage:                     "196",
					TotalWinsMapArMonastery:                   "197",
					TotalWinsMapArShoots:                      "198",
					TotalWinsMapCsAssault:                     "199",
					TotalWinsMapCsItaly:                       "200",
					TotalWinsMapCsMilitia:                     "201",
					TotalWinsMapCsOffice:                      "202",
					TotalWinsMapDeAztec:                       "203",
					TotalWinsMapDeBank:                        "204",
					TotalWinsMapDeCbble:                       "205",
					TotalWinsMapDeDust:                        "206",
					TotalWinsMapDeDust2:                       "207",
					TotalWinsMapDeHouse:                       "208",
					TotalWinsMapDeInferno:                     "209",
					TotalWinsMapDeLake:                        "210",
					TotalWinsMapDeNuke:                        "211",
					TotalWinsMapDeSafehouse:                   "212",
					TotalWinsMapDeShorttrain:                  "213",
					TotalWinsMapDeStmarc:                      "214",
					TotalWinsMapDeSugarcane:                   "215",
					TotalWinsMapDeTrain:                       "216",
					TotalWinsMapDeVertigo:                     "217",
					TotalWinsPistolround:                      "218",
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

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("DataStorage.GetUserStatsForGame() mismatch (-want +got):\n%s", diff)
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("DataStorage.GetUserStatsForGame() = %v, want %v", got, tt.want)
			// }
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
			name: "Update some UserStatsForGame (ID: all_columns)",
			stats: steamclient.UserStatsForGame{
				SteamID:  "all_columns",
				GameName: "1",
				Stats:    steamclient.GameStats{
					//TODO
				},
			},
			want: steamclient.UserStatsForGame{
				SteamID:  "all_columns",
				GameName: "1",
				Stats:    steamclient.GameStats{
					//TODO
				},
			},
		},
		{
			name: "Update all UserStatsForGame (ID: all_columns)",
			stats: steamclient.UserStatsForGame{
				SteamID:  "all_columns",
				GameName: "1",
				Stats:    steamclient.GameStats{

					//TODO
				},
			},
			want: steamclient.UserStatsForGame{
				SteamID:  "all_columns",
				GameName: "1",
				Stats:    steamclient.GameStats{
					//TODO
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
