package database

import (
	"database/sql"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/pinpox/megaclan3000/internal/steamclient"
)

func TestDataStorage_UpdatePlayerSummary(t *testing.T) {
	type fields struct {
		db         *sqlx.DB
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
