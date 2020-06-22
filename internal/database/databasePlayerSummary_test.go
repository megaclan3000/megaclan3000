package database

// import (
// 	"testing"

// 	"github.com/google/go-cmp/cmp"
// 	"github.com/pinpox/megaclan3000/internal/steamclient"
// )

// TODO Fix this test or move to sqlx
// func TestDataStorage_UpdatePlayerSummary(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		ps      steamclient.PlayerSummary
// 		wantErr bool
// 	}{
// 		{
// 			name: "Update all colums for existing ID (ID: all_columns)",
// 			ps: steamclient.PlayerSummary{
// 				SteamID:                  "all_columns",
// 				Avatar:                   "0",
// 				Avatarfull:               "1",
// 				Avatarmedium:             "2",
// 				Cityid:                   "3",
// 				Commentpermission:        "4",
// 				Communityvisibilitystate: "5",
// 				Gameextrainfo:            "6",
// 				Gameid:                   "7",
// 				Gameserverip:             "8",
// 				Lastlogoff:               "9",
// 				Loccityid:                "10",
// 				Loccountrycode:           "11",
// 				Locstatecode:             "12",
// 				Personaname:              "13",
// 				Personastate:             "14",
// 				Primaryclanid:            "15",
// 				Profilestate:             "16",
// 				Profileurl:               "17",
// 				Realname:                 "18",
// 				Timecreated:              "19",
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "Update all colums for new ID (ID: all_new)",
// 			ps: steamclient.PlayerSummary{
// 				SteamID:                  "all_new",
// 				Avatar:                   "new0",
// 				Avatarfull:               "new1",
// 				Avatarmedium:             "new2",
// 				Cityid:                   "new3",
// 				Commentpermission:        "new4",
// 				Communityvisibilitystate: "new5",
// 				Gameextrainfo:            "new6",
// 				Gameid:                   "new7",
// 				Gameserverip:             "new8",
// 				Lastlogoff:               "new9",
// 				Loccityid:                "new10",
// 				Loccountrycode:           "new11",
// 				Locstatecode:             "new12",
// 				Personaname:              "new13",
// 				Personastate:             "new14",
// 				Primaryclanid:            "new15",
// 				Profilestate:             "new16",
// 				Profileurl:               "new17",
// 				Realname:                 "new18",
// 				Timecreated:              "new19",
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			prepareDB()

// 			if err := db.UpdatePlayerSummary(tt.ps); (err != nil) != tt.wantErr {
// 				t.Errorf("DataStorage.UpdatePlayerSummary() error = %v, wantErr %v", err, tt.wantErr)
// 			}

// 			if got, err := db.GetPlayerSummary(tt.ps.SteamID); err == nil {
// 				if diff := cmp.Diff(tt.ps, got); diff != "" {
// 					t.Errorf("DataStorage.UpdatePlayerSummary() mismatch (-want +got):\n%s", diff)
// 				}
// 			} else {
// 				if tt.wantErr == (err != nil) {
// 					t.Errorf("DataStorage.UpdatePlayerSummary() error = %v, wantErr %v", err, tt.wantErr)
// 				}
// 			}
// 		})
// 	}
// }
