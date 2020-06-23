package database

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pinpox/megaclan3000/internal/steamclient"
)

func TestDataStorage_UpdateRecentlyPlayedGames(t *testing.T) {
	tests := []struct {
		name    string
		steamID string
		want    steamclient.RecentlyPlayedGames
		wantErr bool
	}{
		{
			name:    "Change all values (ID: all_columns)",
			steamID: "all_columns",
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
			if err := db.UpdateRecentlyPlayedGames(tt.want); (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.UpdateRecentlyPlayedGames() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := db.GetPlayerInfoBySteamID(tt.steamID)

			if err != nil {

				t.Errorf("Error retrieving RecentlyPlayedGames: %v", err)
			}

			if diff := cmp.Diff(tt.want, got.RecentlyPlayedGames); diff != "" {
				t.Errorf("DataStorage.UpdateRecentlyPlayedGames() mismatch (-want +got):\n%s", diff)
			}

		})
	}
}
