package database

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pinpox/megaclan3000/internal/steamclient"
)

func TestDataStorage_GetRecentlyPlayedGames(t *testing.T) {
	tests := []struct {
		name    string
		steamID string
		want    steamclient.RecentlyPlayedGames
		wantErr bool
	}{

		{
			name:    "Test retrieval of RecentlyPlayedGames from fixtures (ID: no_exist)",
			steamID: "no_exist",
			want:    steamclient.RecentlyPlayedGames{},
			wantErr: false,
		},
		{
			name:    "Test retrieval of RecentlyPlayedGames from fixtures (ID: all_columns)",
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prepareDB()
			got, err := db.GetRecentlyPlayedGames(tt.steamID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.GetRecentlyPlayedGames() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("DataStorage.GetRecentlyPlayedGames() mismatch (-want +got):\n%s", diff)
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("DataStorage.GetRecentlyPlayedGames() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func TestDataStorage_UpdateRecentlyPlayedGames(t *testing.T) {
	tests := []struct {
		name    string
		steamID string
		rpg     steamclient.RecentlyPlayedGames
		want    steamclient.RecentlyPlayedGames
		wantErr bool
	}{
		{
			name:    "Change all values (ID: all_columns)",
			steamID: "all_columns",
			rpg: steamclient.RecentlyPlayedGames{
				SteamID:                "all_columns",
				Appid:                  "changed-0",
				ImgIconURL:             "changed-1",
				ImgLogoURL:             "changed-2",
				Name:                   "changed-3",
				Playtime2Weeks:         "changed-4",
				PlaytimeForever:        "changed-5",
				PlaytimeLinuxForever:   "changed-6",
				PlaytimeMacForever:     "changed-7",
				PlaytimeWindowsForever: "changed-8",
			},
			want: steamclient.RecentlyPlayedGames{
				SteamID:                "all_columns",
				Appid:                  "changed-0",
				ImgIconURL:             "changed-1",
				ImgLogoURL:             "changed-2",
				Name:                   "changed-3",
				Playtime2Weeks:         "changed-4",
				PlaytimeForever:        "changed-5",
				PlaytimeLinuxForever:   "changed-6",
				PlaytimeMacForever:     "changed-7",
				PlaytimeWindowsForever: "changed-8",
			},
			wantErr: false,
		},
		// {
		// 	name:    "Try to change all values (ID: no_exist)",
		// 	steamID: "all_columns",
		// 	rpg: steamclient.RecentlyPlayedGames{
		// 		SteamID:                "all_columns",
		// 		Appid:                  "changed-0",
		// 		ImgIconURL:             "changed-1",
		// 		ImgLogoURL:             "changed-2",
		// 		Name:                   "changed-3",
		// 		Playtime2Weeks:         "changed-4",
		// 		PlaytimeForever:        "changed-5",
		// 		PlaytimeLinuxForever:   "changed-6",
		// 		PlaytimeMacForever:     "changed-7",
		// 		PlaytimeWindowsForever: "changed-8",
		// 	},
		// 	want:    steamclient.RecentlyPlayedGames{},
		// 	wantErr: true,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prepareDB()
			if err := db.UpdateRecentlyPlayedGames(tt.rpg); (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.UpdateRecentlyPlayedGames() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got, err := db.GetRecentlyPlayedGames(tt.steamID); err == nil {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("DataStorage.UpdateRecentlyPlayedGames() mismatch (-want +got):\n%s", diff)
				}
			} else {
				if !(tt.wantErr == (err != nil)) {
					t.Errorf("DataStorage.UpdateRecentlyPlayedGames() error = %v, wantErr %v", err, tt.wantErr)
				}
			}

		})
	}
}
