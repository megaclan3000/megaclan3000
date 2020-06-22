package database

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/pinpox/megaclan3000/internal/steamclient"
)

func TestDataStorage_GetPlayerHistory(t *testing.T) {
	tests := []struct {
		name    string
		steamID string
		want    steamclient.PlayerHistory
		wantErr bool
	}{

		{
			name:    "Retrieve PlayerHistory from fixtures (ID: all_columns)",
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
						Time:                       "9",
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
			name:    "Retrieve PlayerHistory from fixtures (ID: 1)",
			steamID: "1",
			want: steamclient.PlayerHistory{
				SteamID: "1",
				Data: []steamclient.PlayerHistoryEntry{
					{
						SteamID:                    "1",
						Time:                       "100",
						TotalKills:                 "101",
						TotalADR:                   "102",
						TotalShotsHit:              "103",
						TotalShotsFired:            "104",
						TotalKillsHeadshot:         "105",
						TotalKD:                    "106",
						LastMatchContributionScore: "107",
						LastMatchDamage:            "108",
						LastMatchDeaths:            "109",
						LastMatchKills:             "1010",
						LastMatchRounds:            "1011",
						LastMatchKD:                "1012",
						LastMatchADR:               "1013",
						HitRatio:                   "1014",
						Playtime2Weeks:             "1015",
					},
				},
			},
			wantErr: false,
		},

		{
			name:    "Retrieve PlayerHistory from fixtures (ID: 2)",
			steamID: "2",
			want: steamclient.PlayerHistory{
				SteamID: "2",
				Data: []steamclient.PlayerHistoryEntry{
					{
						SteamID:                    "2",
						Time:                       "200",
						TotalKills:                 "201",
						TotalADR:                   "202",
						TotalShotsHit:              "203",
						TotalShotsFired:            "204",
						TotalKillsHeadshot:         "205",
						TotalKD:                    "206",
						LastMatchContributionScore: "207",
						LastMatchDamage:            "208",
						LastMatchDeaths:            "209",
						LastMatchKills:             "2010",
						LastMatchRounds:            "2011",
						LastMatchKD:                "2012",
						LastMatchADR:               "2013",
						HitRatio:                   "2014",
						Playtime2Weeks:             "2015",
					},
				},
			},
			wantErr: false,
		},

		{
			name:    "Retrieve PlayerHistory from fixtures (ID: 3)",
			steamID: "3",
			want: steamclient.PlayerHistory{
				SteamID: "3",
				Data: []steamclient.PlayerHistoryEntry{
					{
						SteamID:                    "3",
						Time:                       "300",
						TotalKills:                 "301",
						TotalADR:                   "302",
						TotalShotsHit:              "303",
						TotalShotsFired:            "304",
						TotalKillsHeadshot:         "305",
						TotalKD:                    "306",
						LastMatchContributionScore: "307",
						LastMatchDamage:            "308",
						LastMatchDeaths:            "309",
						LastMatchKills:             "3010",
						LastMatchRounds:            "3011",
						LastMatchKD:                "3012",
						LastMatchADR:               "3013",
						HitRatio:                   "3014",
						Playtime2Weeks:             "3015",
					},
					{
						SteamID:                    "3",
						Time:                       "3003",
						TotalKills:                 "3013",
						TotalADR:                   "3023",
						TotalShotsHit:              "3033",
						TotalShotsFired:            "3043",
						TotalKillsHeadshot:         "3053",
						TotalKD:                    "3063",
						LastMatchContributionScore: "3073",
						LastMatchDamage:            "3083",
						LastMatchDeaths:            "3093",
						LastMatchKills:             "30103",
						LastMatchRounds:            "30113",
						LastMatchKD:                "30123",
						LastMatchADR:               "30133",
						HitRatio:                   "30143",
						Playtime2Weeks:             "30153",
					},
					{
						SteamID:                    "3",
						Time:                       "30034",
						TotalKills:                 "30134",
						TotalADR:                   "30234",
						TotalShotsHit:              "30334",
						TotalShotsFired:            "30434",
						TotalKillsHeadshot:         "30534",
						TotalKD:                    "30634",
						LastMatchContributionScore: "30734",
						LastMatchDamage:            "30834",
						LastMatchDeaths:            "30934",
						LastMatchKills:             "301034",
						LastMatchRounds:            "301134",
						LastMatchKD:                "301234",
						LastMatchADR:               "301334",
						HitRatio:                   "301434",
						Playtime2Weeks:             "301534",
					},
				},
			},
			wantErr: false,
		},

		{
			name:    "Retrieve PlayerHistory from fixtures (ID: 4)",
			steamID: "4",
			want: steamclient.PlayerHistory{
				SteamID: "4",
				Data: []steamclient.PlayerHistoryEntry{
					{
						SteamID:                    "4",
						Time:                       "400",
						TotalKills:                 "401",
						TotalADR:                   "402",
						TotalShotsHit:              "403",
						TotalShotsFired:            "404",
						TotalKillsHeadshot:         "405",
						TotalKD:                    "406",
						LastMatchContributionScore: "407",
						LastMatchDamage:            "408",
						LastMatchDeaths:            "409",
						LastMatchKills:             "4010",
						LastMatchRounds:            "4011",
						LastMatchKD:                "4012",
						LastMatchADR:               "4013",
						HitRatio:                   "4014",
						Playtime2Weeks:             "4015",
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			prepareDB()

			got, err := db.GetPlayerInfoBySteamID(tt.steamID)

			if err != nil {
				t.Logf("DataStorage.GetPlayerHistory() error = %v, wantErr %v", err, tt.wantErr)
			}

			if diff := cmp.Diff(got.PlayerHistory, tt.want); diff != "" {
				t.Errorf("DataStorage.GetPlayerHistory() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestDataStorage_UpdatePlayerHistory(t *testing.T) {
	tests := []struct {
		name    string
		pi      steamclient.PlayerInfo
		want    steamclient.PlayerHistory
		wantErr bool
	}{
		{
			name: "All values inserted for new ID (ID: all_new)",
			want: steamclient.PlayerHistory{
				SteamID: "all_new",
				Data: []steamclient.PlayerHistoryEntry{
					{
						SteamID:                    "all_new",
						HitRatio:                   "inserted9",
						LastMatchADR:               "inserted10",
						LastMatchContributionScore: "inserted0",
						LastMatchDamage:            "inserted1",
						LastMatchDeaths:            "inserted2",
						LastMatchKD:                "inserted13",
						LastMatchKills:             "inserted3",
						LastMatchRounds:            "inserted2",
						Playtime2Weeks:             "inserted14",
						TotalADR:                   "inserted12",
						TotalKD:                    "inserted13",
						TotalKills:                 "inserted5",
						TotalKillsHeadshot:         "inserted6",
						TotalShotsFired:            "inserted7",
						TotalShotsHit:              "inserted8",

						//This value is set by the database, not tested
						Time: "database",
					},
				},
			},
			pi: steamclient.PlayerInfo{
				PlayerSummary: steamclient.PlayerSummary{
					SteamID: "all_new",
				},
				UserStatsForGame: steamclient.UserStatsForGame{
					Stats: steamclient.GameStats{
						LastMatchContributionScore: "inserted0",
						LastMatchDamage:            "inserted1",
						LastMatchDeaths:            "inserted2",
						LastMatchKills:             "inserted3",
						LastMatchRounds:            "inserted2",
						TotalKills:                 "inserted5",
						TotalKillsHeadshot:         "inserted6",
						TotalShotsFired:            "inserted7",
						TotalShotsHit:              "inserted8",
					},
					Extra: steamclient.GameExtras{
						HitRatio:     "inserted9",
						LastMatchADR: "inserted10",
						LastMatchKD:  "inserted13",
						TotalADR:     "inserted12",
						TotalKD:      "inserted13",
					},
				},
				RecentlyPlayedGames: steamclient.RecentlyPlayedGames{
					Playtime2Weeks: "inserted14",
				},
				PlayerHistory: steamclient.PlayerHistory{},
			},
		},
		{

			//TODO add more cases for other id's

			name: "All values inserted for existing ID (ID: all_columns)",
			want: steamclient.PlayerHistory{
				SteamID: "all_columns",
				Data: []steamclient.PlayerHistoryEntry{
					{
						SteamID:                    "all_columns",
						HitRatio:                   "inserted9",
						LastMatchADR:               "inserted10",
						LastMatchContributionScore: "inserted0",
						LastMatchDamage:            "inserted1",
						LastMatchDeaths:            "inserted2",
						LastMatchKD:                "inserted13",
						LastMatchKills:             "inserted3",
						LastMatchRounds:            "inserted2",
						Playtime2Weeks:             "inserted14",
						TotalADR:                   "inserted12",
						TotalKD:                    "inserted13",
						TotalKills:                 "inserted5",
						TotalKillsHeadshot:         "inserted6",
						TotalShotsFired:            "inserted7",
						TotalShotsHit:              "inserted8",

						//This value is set by the database, not tested
						Time: "database",
					},
					{
						SteamID:                    "all_columns",
						Time:                       "database",
						TotalKD:                    "11",
						TotalADR:                   "10",
						LastMatchADR:               "1",
						TotalKills:                 "12",
						TotalKillsHeadshot:         "13",
						TotalShotsHit:              "15",
						TotalShotsFired:            "14",
						LastMatchContributionScore: "2",
						LastMatchDamage:            "3",
						LastMatchDeaths:            "4",
						LastMatchKills:             "6",
						LastMatchRounds:            "7",
						LastMatchKD:                "5",
						HitRatio:                   "0",
						Playtime2Weeks:             "8",
					},
				},
			},
			pi: steamclient.PlayerInfo{
				PlayerSummary: steamclient.PlayerSummary{
					SteamID: "all_columns",
				},
				UserStatsForGame: steamclient.UserStatsForGame{
					Stats: steamclient.GameStats{
						LastMatchContributionScore: "inserted0",
						LastMatchDamage:            "inserted1",
						LastMatchDeaths:            "inserted2",
						LastMatchKills:             "inserted3",
						LastMatchRounds:            "inserted2",
						TotalKills:                 "inserted5",
						TotalKillsHeadshot:         "inserted6",
						TotalShotsFired:            "inserted7",
						TotalShotsHit:              "inserted8",
					},
					Extra: steamclient.GameExtras{
						HitRatio:     "inserted9",
						LastMatchADR: "inserted10",
						LastMatchKD:  "inserted13",
						TotalADR:     "inserted12",
						TotalKD:      "inserted13",
					},
				},
				RecentlyPlayedGames: steamclient.RecentlyPlayedGames{
					Playtime2Weeks: "inserted14",
				},
				PlayerHistory: steamclient.PlayerHistory{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prepareDB()

			if err := db.UpdatePlayerHistory(tt.pi); (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.UpdatePlayerHistory() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got, err := db.GetPlayerInfoBySteamID(tt.pi.PlayerSummary.SteamID); err == nil {

				//Replace db-generated time field to exclude from test
				for k := range got.PlayerHistory.Data {
					got.PlayerHistory.Data[k].Time = "database"
				}

				if diff := cmp.Diff(tt.want, got.PlayerHistory); diff != "" {
					t.Errorf("DataStorage.UpdatePlayerHistory() mismatch (-want +got):\n%s", diff)
				}
			} else {
				if tt.wantErr == (err != nil) {
					t.Errorf("DataStorage.UpdatePlayerHistory() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestDataStorage_GetPlayerHistoryLatestTime(t *testing.T) {
	tests := []struct {
		name    string
		steamID string
		want    time.Time
		wantErr bool
	}{
		{
			name:    "Get Time for existing ID (ID: all_columns)",
			steamID: "all_columns",
			want:    time.Date(0001, 01, 01, 00, 00, 00, 00, time.UTC),
			wantErr: false,
		},
		{
			name:    "Get Time for existing ID (ID: 1)",
			steamID: "1",
			want:    time.Date(0001, 01, 01, 00, 00, 00, 00, time.UTC),
			wantErr: false,
		},
		{
			name:    "Get Time for existing ID (ID: 2)",
			steamID: "2",
			want:    time.Date(0001, 01, 01, 00, 00, 00, 00, time.UTC),
			wantErr: false,
		},
		{
			name:    "Get Time for existing ID (ID: 3)",
			steamID: "3",
			want:    time.Date(0001, 01, 01, 00, 00, 00, 00, time.UTC),
			wantErr: false,
		},
		{
			name:    "Get Time for existing ID (ID: 41)",
			steamID: "41",
			want:    time.Date(0001, 01, 01, 00, 00, 00, 00, time.UTC),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prepareDB()

			got, err := db.GetPlayerHistoryLatestTime(tt.steamID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.GetPlayerHistoryLatestTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataStorage.GetPlayerHistoryLatestTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
