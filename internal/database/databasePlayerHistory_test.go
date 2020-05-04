package database

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/pinpox/megaclan3000/internal/steamclient"
)

func TestDataStorage_GetPlayerHistory(t *testing.T) {
	type fields struct {
		db         *sql.DB
		statements map[string]*sql.Stmt
	}
	type args struct {
		steamID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    steamclient.PlayerHistory
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
			got, err := ds.GetPlayerHistory(tt.args.steamID)
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
