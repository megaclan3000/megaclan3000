package database

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/pinpox/megaclan3000/internal/steamclient"
)

func TestDataStorage_UpdatePlayerHistory(t *testing.T) {
	tests := []struct {
		name    string
		pi      steamclient.PlayerHistoryEntry
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
						Time:                       time.Date(2090, time.June, 0, 0, 0, 0, 0, time.UTC),
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
					},
				},
			},
			pi: steamclient.PlayerHistoryEntry{

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
				Time: time.Date(2090, time.June, 0, 0, 0, 0, 0, time.UTC),
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
						Time:                       time.Date(2020, time.June, 23, 0, 0, 9, 0, time.UTC),
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

						Time: time.Date(2090, time.June, 0, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			pi: steamclient.PlayerHistoryEntry{
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

				Time: time.Date(2090, time.June, 0, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prepareDB()

			if err := db.UpdatePlayerHistory(tt.pi); (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.UpdatePlayerHistory() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got, err := db.getPlayerHistory(tt.pi.SteamID); err == nil {

				if diff := cmp.Diff(tt.want, got); diff != "" {
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
	}{
		{
			name:    "Get Time for existing ID (ID: all_columns)",
			steamID: "all_columns",
			want:    time.Date(2020, time.June, 23, 0, 0, 9, 00, time.UTC),
		},
		{
			name:    "Get Time for existing ID (ID: 1)",
			steamID: "1",
			want:    time.Date(2020, time.June, 23, 0, 0, 1, 00, time.UTC),
		},
		{
			name:    "Get Time for existing ID (ID: 2)",
			steamID: "2",
			want:    time.Date(2020, time.June, 23, 0, 0, 2, 00, time.UTC),
		},
		{
			name:    "Get Time for existing ID (ID: 3)",
			steamID: "3",
			want:    time.Date(2020, time.June, 23, 0, 0, 5, 00, time.UTC),
		},
		{
			name:    "Get Time for existing ID (ID: 41)",
			steamID: "41",
			want:    time.Date(2020, time.June, 23, 0, 0, 8, 00, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prepareDB()

			got := db.GetPlayerHistoryLatestTime(tt.steamID)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataStorage.GetPlayerHistoryLatestTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
