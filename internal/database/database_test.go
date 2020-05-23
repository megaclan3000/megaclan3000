package database

import (
	"reflect"
	"testing"

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
	if db, err = NewDataStorage("../../test/database/test.db"); err != nil {
		panic(err)
	}

	if fixtures, err = testfixtures.New(
		testfixtures.Database(db.db),
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

func TestDataStorage_GetPlayerInfoBySteamID(t *testing.T) {

	// Needed here because other tests modify the database
	prepareDB()

	// The methods called here are tested on their own, so we might aswell use
	// the output instead of typing all fields again by hand
	summary, err := db.GetPlayerSummary("all_columns")
	if err != nil {
		panic(err)
	}
	stats, err := db.GetUserStatsForGame("all_columns")
	if err != nil {
		panic(err)
	}
	recent, err := db.GetRecentlyPlayedGames("all_columns")
	if err != nil {
		panic(err)
	}
	history, err := db.GetPlayerHistory("all_columns")
	if err != nil {
		panic(err)
	}

	tests := []struct {
		name    string
		steamID string
		want    steamclient.PlayerInfo
		wantErr bool
	}{
		{
			name:    "Retrieve PlayerInfo from fixtures (ID: all_columns)",
			steamID: "all_columns",
			want: steamclient.PlayerInfo{
				PlayerSummary:       summary,
				UserStatsForGame:    stats,
				RecentlyPlayedGames: recent,
				PlayerHistory:       history,
			},
			wantErr: false,
		},
		{
			name:    "Retrieve PlayerInfo from fixtures (ID: no_exist)",
			steamID: "no_exist",
			want: steamclient.PlayerInfo{
				PlayerSummary:       steamclient.PlayerSummary{},
				UserStatsForGame:    steamclient.UserStatsForGame{},
				RecentlyPlayedGames: steamclient.RecentlyPlayedGames{},
				PlayerHistory:       steamclient.PlayerHistory{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prepareDB()
			got, err := db.GetPlayerInfoBySteamID(tt.steamID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.GetPlayerInfoBySteamID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("DataStorage.GetPlayerInfoBySteamID() mismatch (-want +got):\n%s", diff)
			}
		})
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
