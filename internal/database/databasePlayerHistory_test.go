package database

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/pinpox/megaclan3000/internal/steamclient"
)

func TestDataStorage_GetPlayerHistory(t *testing.T) {
	tests := []struct {
		name    string
		steamID string
		want    steamclient.PlayerHistory
		wantErr bool
	}{
		{
			name:    "Retrieve PlayerHistory from fixtures",
			steamID: "123456789",
			want: steamclient.PlayerHistory{
				SteamID: "123456789",
				Data: []steamclient.PlayerHistoryEntry{
					{
						Time:    3000,
						TotalKD: "0.3",
					},
					{
						Time:    2000,
						TotalKD: "0.2",
					},
					{
						Time:    1000,
						TotalKD: "0.1",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			prepareDB()
			got, err := db.GetPlayerHistory(tt.steamID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.GetPlayerHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataStorage.GetPlayerHistory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataStorage_GetPlayerHistoryLatestTime(t *testing.T) {
	tests := []struct {
		name    string
		steamID string
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			prepareDB()
			got, err := db.GetPlayerHistoryLatestTime(tt.steamID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.GetPlayerHistoryLatestTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DataStorage.GetPlayerHistoryLatestTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataStorage_UpdatePlayerHistory(t *testing.T) {
	type fields struct {
		db         *sql.DB
		statements map[string]*sql.Stmt
	}
	type args struct {
		pi steamclient.PlayerInfo
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
			if err := ds.UpdatePlayerHistory(tt.args.pi); (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.UpdatePlayerHistory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
