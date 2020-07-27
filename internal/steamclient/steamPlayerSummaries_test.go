package steamclient

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"io/ioutil"
	"os"
	"testing"
)

func TestSteamClient_parsePlayerSummary(t *testing.T) {

	steamConfig := SteamConfig{
		SteamAPIKey:     "000000000",
		SteamIDs:        []string{},
		HistoryInterval: 0,
		UpdateInterval:  0,
	}
	tests := []struct {
		name     string
		want     PlayerSummary
		dataFile string
		wantErr  bool
	}{
		{
			name: "Parse data for ID: 76561197962156894",
			want: PlayerSummary{
				SteamID:                  76561197962156894,
				Personaname:              "mac5",
				Profileurl:               "https://steamcommunity.com/profiles/76561197962156894/",
				Avatar:                   "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/01/01b6a06e6946ef4c2525ba425c29d010795ec1ff.jpg",
				Avatarmedium:             "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/01/01b6a06e6946ef4c2525ba425c29d010795ec1ff_medium.jpg",
				Avatarfull:               "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/01/01b6a06e6946ef4c2525ba425c29d010795ec1ff_full.jpg",
				Personastate:             "0",
				Communityvisibilitystate: "3",
				Profilestate:             "1",
				Lastlogoff:               "1591043574",
				Primaryclanid:            "103582791435284539",
				Timecreated:              "1066582822",
			},
			dataFile: "../../test/steamclient/GetPlayerSummaries76561197962156894.json",
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
			var data playerSummariesData
			if err = json.Unmarshal(byteValue, &data); err != nil {
				panic(err)
			}

			// Create a SteamClient
			sc := &SteamClient{Config: steamConfig}

			// Try to parse
			got, err := sc.parsePlayerSummary(data)

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("SteamClient.parsePlayerSummary() mismatch (-want +got):\n%s", diff)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("SteamClient.parsePlayerSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
