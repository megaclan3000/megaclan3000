package database

import (
	"testing"
	"time"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/google/go-cmp/cmp"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pinpox/megaclan3000/internal/steamclient"
)

var (
	db       *DataStorage
	fixtures *testfixtures.Loader
)

func prepareDB() {

	var err error
	if db, err = NewDataStorage("../../test/database/test.db", "../../schema.sql"); err != nil {
		panic(err)
	}

	if fixtures, err = testfixtures.New(
		testfixtures.Database(db.db.DB),
		testfixtures.Dialect("sqlite"),
		testfixtures.Directory(
			"../../test/database/fixtures",
		),
	); err != nil {
		panic(err)
	}

	if err = fixtures.Load(); err != nil {
		panic(err)
	}
}

func TestDataStorage_GetAllPlayers(t *testing.T) {
	tests := []struct {
		name    string
		want    []steamclient.PlayerInfo
		wantErr bool
	}{
		// all_columns
		// other1
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prepareDB()
			got, err := db.GetAllPlayers()
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.GetAllPlayers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("DataStorage.GetAllPlayers() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDataStorage_UpdatePlayerInfo(t *testing.T) {
	tests := []struct {
		name    string
		pi      steamclient.PlayerInfo
		wantErr bool
	}{
		{
			name: "Update player info for existing ID (ID: all_columns)",
			pi: steamclient.PlayerInfo{
				//TODO add tests
				PlayerSummary: steamclient.PlayerSummary{
					SteamID: "all_columns",
				},
				UserStatsForGame: steamclient.UserStatsForGame{
					SteamID: "all_columns",
				},
				RecentlyPlayedGames: steamclient.RecentlyPlayedGames{},
				PlayerHistory:       steamclient.PlayerHistory{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prepareDB()
			if err := db.UpdatePlayerInfo(tt.pi); (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.UpdatePlayerInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDataStorage_getPlayerSummary(t *testing.T) {
	tests := []struct {
		name    string
		steamID string
		want    steamclient.PlayerSummary
		wantErr bool
	}{
		{
			name:    "Retrieve PlayerSummary from fixtures for ID: all_columns",
			steamID: "all_columns",
			want: steamclient.PlayerSummary{
				SteamID:                  "all_columns",
				Avatar:                   "0",
				Avatarfull:               "1",
				Avatarmedium:             "2",
				Cityid:                   "3",
				Commentpermission:        "4",
				Communityvisibilitystate: "5",
				Gameextrainfo:            "6",
				Gameid:                   "7",
				Gameserverip:             "8",
				Lastlogoff:               "9",
				Loccityid:                "10",
				Loccountrycode:           "11",
				Locstatecode:             "12",
				Personaname:              "13",
				Personastate:             "14",
				Primaryclanid:            "15",
				Profilestate:             "16",
				Profileurl:               "17",
				Realname:                 "18",
				Timecreated:              "19",
			},
			wantErr: false,
		},
		{
			name:    "Try to retrieve PlayerSummary from fixtures for ID: no_exist",
			steamID: "no_exist",
			want:    steamclient.PlayerSummary{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			prepareDB()
			got, err := db.getPlayerSummary(tt.steamID)

			if (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.getPlayerSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("DataStorage.getPlayerSummary() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDataStorage_getRecentlyPlayedGames(t *testing.T) {
	tests := []struct {
		name    string
		steamID string
		want    steamclient.RecentlyPlayedGames
		wantErr bool
	}{
		{
			name:    "Retrieve RecentlyPlayedGames from fixtures for ID: all_columns",
			steamID: "all_columns",
			want: steamclient.RecentlyPlayedGames{
				SteamID:                "all_columns",
				Appid:                  "0",
				ImgIconURL:             "1",
				ImgLogoURL:             "2",
				Name:                   "3",
				Playtime2Weeks:         "4",
				PlaytimeForever:        "5",
				PlaytimeLinuxForever:   "6",
				PlaytimeMacForever:     "7",
				PlaytimeWindowsForever: "8",
			},
			wantErr: false,
		},

		{
			name:    "Try to retrieve RecentlyPlayedGames from fixtures for ID: no_exist",
			steamID: "no_exist",
			want:    steamclient.RecentlyPlayedGames{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			prepareDB()
			got, err := db.getRecentlyPlayedGames(tt.steamID)

			if (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.getRecentlyPlayedGames() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("DataStorage.getRecentlyPlayedGames() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDataStorage_getPlayerHistory(t *testing.T) {
	tests := []struct {
		name    string
		steamID string
		want    steamclient.PlayerHistory
		wantErr bool
	}{
		{
			name:    "Retrieve PlayerHistory from fixtures for ID: all_columns",
			steamID: "all_columns",
			want: steamclient.PlayerHistory{
				SteamID: "all_columns",
				Data: []steamclient.PlayerHistoryEntry{
					{
						SteamID:                    "all_columns",
						HitRatio:                   "0",
						LastMatchADR:               "1",
						LastMatchContributionScore: "2",
						LastMatchDamage:            "3",
						LastMatchDeaths:            "4",
						LastMatchKD:                "5",
						LastMatchKills:             "6",
						LastMatchRounds:            "7",
						Playtime2Weeks:             "8",
						Time:                       time.Date(2020, time.June, 23, 0, 0, 9, 00, time.UTC),
						TotalADR:                   "10",
						TotalKD:                    "11",
						TotalKills:                 "12",
						TotalKillsHeadshot:         "13",
						TotalShotsFired:            "14",
						TotalShotsHit:              "15",
					},
				},
			},
			wantErr: false,
		},
		{
			name:    "Retrieve PlayerHistory from fixtures for ID: 3",
			steamID: "3",
			want: steamclient.PlayerHistory{
				SteamID: "3",
				Data: []steamclient.PlayerHistoryEntry{

					{
						SteamID:                    "3",
						Time:                       time.Date(2020, time.June, 23, 0, 0, 3, 00, time.UTC),
						HitRatio:                   "300",
						LastMatchADR:               "301",
						LastMatchContributionScore: "302",
						LastMatchDamage:            "303",
						LastMatchDeaths:            "304",
						LastMatchKD:                "305",
						LastMatchKills:             "306",
						LastMatchRounds:            "307",
						Playtime2Weeks:             "308",
						TotalADR:                   "309",
						TotalKD:                    "310",
						TotalKills:                 "311",
						TotalKillsHeadshot:         "312",
						TotalShotsFired:            "313",
						TotalShotsHit:              "314",
					},
					{
						SteamID:                    "3",
						Time:                       time.Date(2020, time.June, 23, 0, 0, 4, 00, time.UTC),
						HitRatio:                   "360",
						LastMatchADR:               "361",
						LastMatchContributionScore: "362",
						LastMatchDamage:            "363",
						LastMatchDeaths:            "364",
						LastMatchKD:                "365",
						LastMatchKills:             "366",
						LastMatchRounds:            "367",
						Playtime2Weeks:             "368",
						TotalADR:                   "369",
						TotalKD:                    "370",
						TotalKills:                 "371",
						TotalKillsHeadshot:         "372",
						TotalShotsFired:            "373",
						TotalShotsHit:              "374",
					},
					{
						SteamID:                    "3",
						Time:                       time.Date(2020, time.June, 23, 0, 0, 5, 00, time.UTC),
						HitRatio:                   "330",
						LastMatchADR:               "331",
						LastMatchContributionScore: "332",
						LastMatchDamage:            "333",
						LastMatchDeaths:            "334",
						LastMatchKD:                "335",
						LastMatchKills:             "336",
						LastMatchRounds:            "337",
						Playtime2Weeks:             "338",
						TotalADR:                   "339",
						TotalKD:                    "340",
						TotalKills:                 "341",
						TotalKillsHeadshot:         "342",
						TotalShotsFired:            "343",
						TotalShotsHit:              "344",
					},
				},
			},
			wantErr: false,
		},
		{
			name:    "Try to retrieve UserStatsForGamefrom fixtures for ID: no_exist",
			steamID: "no_exist",
			want: steamclient.PlayerHistory{
				SteamID: "no_exist",
				Data:    []steamclient.PlayerHistoryEntry{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			prepareDB()
			got, err := db.getPlayerHistory(tt.steamID)

			if (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.getPlayerHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("DataStorage.getPlayerHistory() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDataStorage_getUserStatsForGameExtra(t *testing.T) {
	tests := []struct {
		name    string
		steamID string
		want    steamclient.GameExtras
		wantErr bool
	}{
		{
			name:    "Retrieve Extra from fixtures (ID: all_columns)",
			steamID: "all_columns",
			want: steamclient.GameExtras{
				SteamID:      "all_columns",
				TotalKD:      "extra0",
				LastMatchKD:  "extra1",
				HitRatio:     "extra2",
				PlayedHours:  "extra3",
				TotalADR:     "extra4",
				LastMatchADR: "extra5",
			},
			wantErr: false,
		},
		{
			name:    "Retrieve Extra from fixtures (ID: no_exist)",
			steamID: "no_exist",
			want:    steamclient.GameExtras{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			prepareDB()
			got, err := db.getUserStatsForGameExtra(tt.steamID)

			if (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.getUserStatsForGameExtra() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("DataStorage.getUserStatsForGameExtra() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDataStorage_getUserStatsForGameStats(t *testing.T) {
	tests := []struct {
		name    string
		steamID string
		want    steamclient.GameStats
		wantErr bool
	}{
		{
			name:    "Retrieve Stats from fixtures (ID: all_columns)",
			steamID: "all_columns",
			want: steamclient.GameStats{
				SteamID:                                   "all_columns",
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
				LastMatchFavweaponHits:                    "24",
				LastMatchFavweaponID:                      "25",
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
			wantErr: false,
		},
		{
			name:    "Retrieve Stats from fixtures (ID: no_exist)",
			steamID: "no_exist",
			want:    steamclient.GameStats{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			prepareDB()
			got, err := db.getUserStatsForGameStats(tt.steamID)

			if (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.getUserStatsForGameStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("DataStorage.getUserStatsForGameStats() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
