package steamclient

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"io/ioutil"
	"os"
	"testing"
)

func TestSteamClient_ParseRecentlyPlayedGames(t *testing.T) {

	steamConfig := SteamConfig{
		SteamAPIKey:     "000000000",
		SteamIDs:        []string{},
		HistoryInterval: 0,
		UpdateInterval:  0,
	}
	tests := []struct {
		name     string
		want     RecentlyPlayedGames
		steamID  string
		dataFile string
		wantErr  bool
	}{
		{
			name:     "Parse data for ID (randolf): 76561198092006615",
			steamID:  "76561198092006615",
			dataFile: "../../test/steamclient/GetRecentlyPlayedGames76561198092006615.json",
			wantErr:  false,
			want: RecentlyPlayedGames{
				SteamID:                "76561198092006615",
				Appid:                  "730",
				Name:                   "Counter-Strike: Global Offensive",
				Playtime2Weeks:         "11",
				PlaytimeForever:        "1091",
				ImgIconURL:             "69f7ebe2735c366c65c0b33dae00e12dc40edbe4",
				ImgLogoURL:             "d0595ff02f5c79fd19b06f4d6165c3fda2372820",
				PlaytimeWindowsForever: "0",
				PlaytimeMacForever:     "0",
				PlaytimeLinuxForever:   "381",
			},
		},
		{
			name:     "Parse data for ID with no data (mac5): 76561197962156894",
			steamID:  "76561197962156894",
			dataFile: "../../test/steamclient/GetRecentlyPlayedGames76561197962156894.json",
			wantErr:  true,
			want: RecentlyPlayedGames{
				SteamID:                "76561197962156894",
				Appid:                  "",
				Name:                   "",
				Playtime2Weeks:         "",
				PlaytimeForever:        "",
				ImgIconURL:             "",
				ImgLogoURL:             "",
				PlaytimeWindowsForever: "",
				PlaytimeMacForever:     "",
				PlaytimeLinuxForever:   "",
			},
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
			var data recentlyPlayedGamesData
			json.Unmarshal(byteValue, &data)

			// Create a SteamClient
			sc := &SteamClient{Config: steamConfig}

			// Try to parse
			got, err := sc.ParseRecentlyPlayedGames(data, tt.steamID)

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("SteamClient.ParseRecentlyPlayedGames() mismatch (-want +got):\n%s", diff)
			}
			if (err != nil) != tt.wantErr {

				t.Errorf("SteamClient.ParseRecentlyPlayedGames() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
