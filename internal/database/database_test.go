package database

import (
	"reflect"
	"testing"

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

// TODO WRITE THIS TEST, IMPORTANT!
func TestDataStorage_GetPlayerInfoBySteamID(t *testing.T) {
	tests := []struct {
		name    string
		steamID string
		want    steamclient.PlayerInfo
		wantErr bool
	}{
		{
			name:    "Retrieve complete PlayerInfo for all_columns",
			steamID: "all_columns",
			want:    steamclient.PlayerInfo{
				//TODO
			},
			wantErr: false,
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataStorage.GetPlayerInfoBySteamID() = %v, want %v", got, tt.want)
			}
		})
	}
}
