package database

import (
	"database/sql"
	"reflect"

	"testing"

	"github.com/pinpox/megaclan3000/internal/steamclient"
)

func TestDataStorage_GetPlayerSummary(t *testing.T) {

	tests := []struct {
		name    string
		steamID string
		want    steamclient.PlayerSummary
		wantErr bool
	}{
		{
			name:    "Test retrieval of PlayerSummary from fixtures",
			steamID: "123456789",
			want: steamclient.PlayerSummary{
				SteamID:                  "123456789",
				Communityvisibilitystate: "1",
				Profilestate:             "2",
				Personaname:              "3",
				Profileurl:               "4",
				Avatar:                   "5",
				Avatarmedium:             "6",
				Avatarfull:               "7",
				Lastlogoff:               "8",
				Personastate:             "9",
				Primaryclanid:            "10",
				Timecreated:              "11",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			prepareDB()
			got, err := db.GetPlayerSummary(tt.steamID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.GetPlayerSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataStorage.GetPlayerSummary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataStorage_UpdatePlayerSummary(t *testing.T) {
	type fields struct {
		db         *sql.DB
		statements map[string]*sql.Stmt
	}
	type args struct {
		ps steamclient.PlayerSummary
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &DataStorage{
				db:         tt.fields.db,
				statements: tt.fields.statements,
			}
			if err := ds.UpdatePlayerSummary(tt.args.ps); (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.UpdatePlayerSummary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
