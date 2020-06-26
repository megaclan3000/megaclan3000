package steamclient

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"io/ioutil"
	"os"
	"testing"
)

func TestSteamClient_ParseUserStatsForGame(t *testing.T) {

	steamConfig := SteamConfig{
		SteamAPIKey:     "000000000",
		SteamIDs:        []string{},
		HistoryInterval: 0,
		UpdateInterval:  0,
	}

	tests := []struct {
		name     string
		want     UserStatsForGame
		dataFile string
		wantErr  bool
	}{
		{
			name: "Parse data for ID: 76561197962156894",
			want: UserStatsForGame{
				SteamID:  "76561197962156894",
				GameName: "ValveTestApp260",
				Extra: GameExtras{
					SteamID:     "76561197962156894",
					TotalKD:     "1.054",
					LastMatchKD: "1.316",
					HitRatio:    "0.202",
					PlayedHours: "66",
					TotalADR:    "188.409",
				},
				Stats: GameStats{

					SteamID:                                   "76561197962156894",
					GILessonCsgoInstrExplainInspect:           "32",
					GILessonBombSitesA:                        "0",
					GILessonBombSitesB:                        "0",
					GILessonCsgoCycleWeaponsKb:                "0",
					GILessonCsgoHostageLeadToHrz:              "1",
					GILessonCsgoInstrExplainBombCarrier:       "1",
					GILessonCsgoInstrExplainBuyarmor:          "16",
					GILessonCsgoInstrExplainBuymenu:           "16",
					GILessonCsgoInstrExplainFollowBomber:      "1",
					GILessonCsgoInstrExplainPickupBomb:        "1",
					GILessonCsgoInstrExplainPlantBomb:         "16",
					GILessonCsgoInstrExplainPreventBombPickup: "1",
					GILessonCsgoInstrExplainReload:            "16",
					GILessonCsgoInstrExplainZoom:              "16",
					GILessonCsgoInstrRescueZone:               "1",
					GILessonDefusePlantedBomb:                 "1",
					GILessonFindPlantedBomb:                   "1",
					GILessonTrExplainPlantBomb:                "16",
					GILessonVersionNumber:                     "16",
					LastMatchContributionScore:                "40",
					LastMatchCtWins:                           "9",
					LastMatchDamage:                           "4145",
					LastMatchDeaths:                           "19",
					LastMatchDominations:                      "0",
					LastMatchFavweaponHits:                    "11",
					LastMatchFavweaponID:                      "9",
					LastMatchFavweaponKills:                   "11",
					LastMatchFavweaponShots:                   "17",
					LastMatchGgContributionScore:              "0",
					LastMatchKills:                            "25",
					LastMatchMaxPlayers:                       "10",
					LastMatchMoneySpent:                       "69700",
					LastMatchMvps:                             "3",
					LastMatchRevenges:                         "0",
					LastMatchRounds:                           "22",
					LastMatchTWins:                            "13",
					LastMatchWins:                             "6",
					SteamStatMatchwinscomp:                    "0",
					SteamStatSurvivedz:                        "0",
					SteamStatXpearnedgames:                    "2",
					TotalBrokenWindows:                        "0",
					TotalContributionScore:                    "13379",
					TotalDamageDone:                           "634865",
					TotalDeaths:                               "3940",
					TotalDefusedBombs:                         "73",
					TotalDominationOverkills:                  "89",
					TotalDominations:                          "36",
					TotalGgMatchesPlayed:                      "37",
					TotalGgMatchesWon:                         "1",
					TotalGunGameContributionScore:             "219",
					TotalGunGameRoundsPlayed:                  "49",
					TotalGunGameRoundsWon:                     "25",
					TotalHitsAk47:                             "2236",
					TotalHitsAug:                              "125",
					TotalHitsAwp:                              "1175",
					TotalHitsBizon:                            "352",
					TotalHitsDeagle:                           "130",
					TotalHitsElite:                            "157",
					TotalHitsFamas:                            "237",
					TotalHitsFiveseven:                        "66",
					TotalHitsG3sg1:                            "0",
					TotalHitsGalilar:                          "323",
					TotalHitsGlock:                            "1103",
					TotalHitsHkp2000:                          "1391",
					TotalHitsM249:                             "75",
					TotalHitsM4a1:                             "0",
					TotalHitsMac10:                            "0",
					TotalHitsMag7:                             "0",
					TotalHitsMp7:                              "0",
					TotalHitsMp9:                              "0",
					TotalHitsNegev:                            "792",
					TotalHitsNova:                             "240",
					TotalHitsP250:                             "133",
					TotalHitsP90:                              "178",
					TotalHitsSawedoff:                         "16",
					TotalHitsScar20:                           "61",
					TotalHitsSg556:                            "229",
					TotalHitsSsg08:                            "84",
					TotalHitsTec9:                             "0",
					TotalHitsUmp45:                            "0",
					TotalHitsXm1014:                           "0",
					TotalKills:                                "4153",
					TotalKillsAgainstZoomedSniper:             "399",
					TotalKillsAk47:                            "667",
					TotalKillsAug:                             "32",
					TotalKillsAwp:                             "1103",
					TotalKillsBizon:                           "57",
					TotalKillsDeagle:                          "54",
					TotalKillsElite:                           "42",
					TotalKillsEnemyBlinded:                    "63",
					TotalKillsEnemyWeapon:                     "366",
					TotalKillsFamas:                           "58",
					TotalKillsFiveseven:                       "17",
					TotalKillsG3sg1:                           "25",
					TotalKillsGalilar:                         "79",
					TotalKillsGlock:                           "244",
					TotalKillsHeadshot:                        "1790",
					TotalKillsHegrenade:                       "6",
					TotalKillsHkp2000:                         "357",
					TotalKillsKnife:                           "43",
					TotalKillsKnifeFight:                      "9",
					TotalKillsM249:                            "18",
					TotalKillsM4a1:                            "699",
					TotalKillsMac10:                           "19",
					TotalKillsMag7:                            "6",
					TotalKillsMolotov:                         "9",
					TotalKillsMp7:                             "76",
					TotalKillsMp9:                             "12",
					TotalKillsNegev:                           "190",
					TotalKillsNova:                            "21",
					TotalKillsP250:                            "26",
					TotalKillsP90:                             "24",
					TotalKillsSawedoff:                        "3",
					TotalKillsScar20:                          "32",
					TotalKillsSg556:                           "68",
					TotalKillsSsg08:                           "39",
					TotalKillsTaser:                           "1",
					TotalKillsTec9:                            "28",
					TotalKillsUmp45:                           "28",
					TotalKillsXm1014:                          "70",
					TotalMatchesPlayed:                        "273",
					TotalMatchesWon:                           "92",
					TotalMatchesWonBaggage:                    "0",
					TotalMatchesWonBank:                       "1",
					TotalMatchesWonLake:                       "1",
					TotalMatchesWonSafehouse:                  "0",
					TotalMatchesWonShoots:                     "0",
					TotalMatchesWonStmarc:                     "0",
					TotalMatchesWonSugarcane:                  "0",
					TotalMatchesWonTrain:                      "22",
					TotalMoneyEarned:                          "11502300",
					TotalMvps:                                 "823",
					TotalPlantedBombs:                         "259",
					TotalProgressiveMatchesWon:                "2",
					TotalRescuedHostages:                      "0",
					TotalRevenges:                             "14",
					TotalRoundsMapArBaggage:                   "0",
					TotalRoundsMapArMonastery:                 "0",
					TotalRoundsMapArShoots:                    "0",
					TotalRoundsMapCsAssault:                   "29",
					TotalRoundsMapCsItaly:                     "0",
					TotalRoundsMapCsMilitia:                   "0",
					TotalRoundsMapCsOffice:                    "0",
					TotalRoundsMapDeAztec:                     "0",
					TotalRoundsMapDeBank:                      "18",
					TotalRoundsMapDeCbble:                     "12",
					TotalRoundsMapDeDust:                      "0",
					TotalRoundsMapDeDust2:                     "895",
					TotalRoundsMapDeInferno:                   "1009",
					TotalRoundsMapDeLake:                      "1",
					TotalRoundsMapDeNuke:                      "0",
					TotalRoundsMapDeSafehouse:                 "1",
					TotalRoundsMapDeShorttrain:                "0",
					TotalRoundsMapDeStmarc:                    "0",
					TotalRoundsMapDeSugarcane:                 "0",
					TotalRoundsMapDeTrain:                     "956",
					TotalRoundsMapDeVertigo:                   "1",
					TotalRoundsPlayed:                         "4008",
					TotalShotsAk47:                            "11025",
					TotalShotsAug:                             "421",
					TotalShotsAwp:                             "2367",
					TotalShotsBizon:                           "1915",
					TotalShotsDeagle:                          "493",
					TotalShotsElite:                           "1222",
					TotalShotsFamas:                           "991",
					TotalShotsFired:                           "66879",
					TotalShotsFiveseven:                       "297",
					TotalShotsG3sg1:                           "156",
					TotalShotsGalilar:                         "1791",
					TotalShotsGlock:                           "5172",
					TotalShotsHit:                             "13477",
					TotalShotsHkp2000:                         "5770",
					TotalShotsM249:                            "1023",
					TotalShotsM4a1:                            "11880",
					TotalShotsMac10:                           "683",
					TotalShotsMag7:                            "329",
					TotalShotsMp7:                             "2357",
					TotalShotsMp9:                             "506",
					TotalShotsNegev:                           "10860",
					TotalShotsNova:                            "954",
					TotalShotsP250:                            "648",
					TotalShotsP90:                             "954",
					TotalShotsSawedoff:                        "83",
					TotalShotsScar20:                          "266",
					TotalShotsSg556:                           "1245",
					TotalShotsSsg08:                           "205",
					TotalShotsTaser:                           "16",
					TotalShotsTec9:                            "572",
					TotalShotsUmp45:                           "642",
					TotalShotsXm1014:                          "2036",
					TotalTRDefusedBombs:                       "0",
					TotalTRPlantedBombs:                       "0",
					TotalTimePlayed:                           "241144",
					TotalTrbombMatchesWon:                     "1",
					TotalWeaponsDonated:                       "302",
					TotalWins:                                 "1917",
					TotalWinsMapArBaggage:                     "0",
					TotalWinsMapArMonastery:                   "0",
					TotalWinsMapArShoots:                      "0",
					TotalWinsMapCsAssault:                     "23",
					TotalWinsMapCsItaly:                       "0",
					TotalWinsMapCsMilitia:                     "0",
					TotalWinsMapCsOffice:                      "0",
					TotalWinsMapDeAztec:                       "0",
					TotalWinsMapDeBank:                        "11",
					TotalWinsMapDeCbble:                       "3",
					TotalWinsMapDeDust:                        "0",
					TotalWinsMapDeDust2:                       "417",
					TotalWinsMapDeHouse:                       "0",
					TotalWinsMapDeInferno:                     "500",
					TotalWinsMapDeLake:                        "1",
					TotalWinsMapDeNuke:                        "0",
					TotalWinsMapDeSafehouse:                   "0",
					TotalWinsMapDeShorttrain:                  "0",
					TotalWinsMapDeStmarc:                      "0",
					TotalWinsMapDeSugarcane:                   "0",
					TotalWinsMapDeTrain:                       "434",
					TotalWinsMapDeVertigo:                     "0",
					TotalWinsPistolround:                      "204",
				},
			},
			dataFile: "../../test/steamclient/GetUserStatsForGame76561197962156894.json",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Get data from file path
			jsonFile, err := os.Open(tt.dataFile)
			if err != nil {
				panic(err)
			}
			defer jsonFile.Close()
			byteValue, _ := ioutil.ReadAll(jsonFile)
			var data userStatsForGameData
			json.Unmarshal(byteValue, &data)

			// Create a SteamClient
			sc := &SteamClient{Config: steamConfig}

			// Try to parse
			got, err := sc.ParseUserStatsForGame(data)

			if (err != nil) != tt.wantErr {
				t.Errorf("SteamClient.ParseUserStatsForGame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("SteamClient.ParseUserStatsForGame() mismatch (-want +got):\n%s", diff)
			}

		})
	}
}

func Test_nilToZeroString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "String with number 123",
			input: "123",
			want:  "123",
		},
		{
			name:  "String with number 0",
			input: "0",
			want:  "0",
		},
		{
			name:  "Empty String",
			input: "",
			want:  "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nilToZeroString(tt.input); got != tt.want {
				t.Errorf("nilToZeroString() = %v, want %v", got, tt.want)
			}
		})
	}
}
