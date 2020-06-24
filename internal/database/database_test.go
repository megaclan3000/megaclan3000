package database

import (
	"reflect"
	"testing"
	"time"

	"github.com/go-testfixtures/testfixtures/v3"
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataStorage.GetAllPlayers() = %v, want %v", got, tt.want)
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataStorage.getPlayerSummary() = %v, want %v", got, tt.want)
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataStorage.getRecentlyPlayedGames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataStorage_getUserStatsForGame(t *testing.T) {
	tests := []struct {
		name    string
		steamID string
		want    steamclient.UserStatsForGame
		wantErr bool
	}{
		{
			name:    "Retrieve UserStatsForGame from fixtures for ID: all_columns",
			steamID: "all_columns",
			want:    steamclient.UserStatsForGame{
				//TODO
			},
			wantErr: false,
		},

		{
			name:    "Try to retrieve UserStatsForGamefrom fixtures for ID: no_exist",
			steamID: "no_exist",
			want:    steamclient.UserStatsForGame{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			prepareDB()
			got, err := db.getUserStatsForGame(tt.steamID)

			if (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.getUserStatsForGame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataStorage.getUserStatsForGame() = %v, want %v", got, tt.want)
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
						HitRatio:                   "0",
						LastMatchADR:               "1",
						LastMatchContributionScore: "2",
						LastMatchDamage:            "3",
						LastMatchDeaths:            "4",
						LastMatchKD:                "5",
						LastMatchKills:             "6",
						LastMatchRounds:            "7",
						Playtime2Weeks:             "8",
						TotalADR:                   "10",
						TotalKD:                    "11",
						TotalKills:                 "12",
						TotalKillsHeadshot:         "13",
						TotalShotsFired:            "14",
						TotalShotsHit:              "15",
					},
					{
						SteamID:                    "3",
						Time:                       time.Date(2020, time.June, 23, 0, 0, 4, 00, time.UTC),
						HitRatio:                   "0",
						LastMatchADR:               "1",
						LastMatchContributionScore: "2",
						LastMatchDamage:            "3",
						LastMatchDeaths:            "4",
						LastMatchKD:                "5",
						LastMatchKills:             "6",
						LastMatchRounds:            "7",
						Playtime2Weeks:             "8",
						TotalADR:                   "10",
						TotalKD:                    "11",
						TotalKills:                 "12",
						TotalKillsHeadshot:         "13",
						TotalShotsFired:            "14",
						TotalShotsHit:              "15",
					},
					{
						SteamID:                    "3",
						Time:                       time.Date(2020, time.June, 23, 0, 0, 5, 00, time.UTC),
						HitRatio:                   "0",
						LastMatchADR:               "1",
						LastMatchContributionScore: "2",
						LastMatchDamage:            "3",
						LastMatchDeaths:            "4",
						LastMatchKD:                "5",
						LastMatchKills:             "6",
						LastMatchRounds:            "7",
						Playtime2Weeks:             "8",
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
			name:    "Try to retrieve UserStatsForGamefrom fixtures for ID: no_exist",
			steamID: "no_exist",
			want:    steamclient.PlayerHistory{},
			wantErr: true,
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataStorage.getPlayerHistory() = %v, want %v", got, tt.want)
			}
		})
	}
}
