package database

import (
	"database/sql"
	"reflect"
	"testing"

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
			name:    "Test retrieval of RecentlyPlayedGames from fixtures (ID: all_columns)",
			steamID: "all_columns",
			want:    steamclient.RecentlyPlayedGames{
				//TODO add test cases
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataStorage.GetRecentlyPlayedGames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataStorage_UpdateRecentlyPlayedGames(t *testing.T) {
	type fields struct {
		db         *sql.DB
		statements map[string]*sql.Stmt
	}
	type args struct {
		rpg steamclient.RecentlyPlayedGames
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
			if err := ds.UpdateRecentlyPlayedGames(tt.args.rpg); (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.UpdateRecentlyPlayedGames() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
