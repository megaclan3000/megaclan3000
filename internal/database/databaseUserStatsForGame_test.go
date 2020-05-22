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
			name:    "Retrieval of UserStatsForGame from fixtures (ID: all_columns)",
			steamID: "all_columns",
			want: steamclient.UserStatsForGame{
				SteamID: "all_columns",
				Stats:   steamclient.GameStats{
					//TODO
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
