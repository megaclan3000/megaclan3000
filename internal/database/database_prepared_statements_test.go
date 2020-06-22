package database

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"testing"
)

func TestDataStorage_getUpdatePreparedstatements(t *testing.T) {
	type fields struct {
		db         *sqlx.DB
		statements map[string]*sql.Stmt
	}
	tests := []struct {
		name    string
		fields  fields
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
			if err := ds.getUpdatePreparedstatements(); (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.getUpdatePreparedstatements() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDataStorage_getInsertPreparedstatements(t *testing.T) {
	type fields struct {
		db         *sqlx.DB
		statements map[string]*sql.Stmt
	}
	tests := []struct {
		name    string
		fields  fields
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
			if err := ds.getInsertPreparedstatements(); (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.getInsertPreparedstatements() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDataStorage_getSelectPreparedstatements(t *testing.T) {
	type fields struct {
		db         *sqlx.DB
		statements map[string]*sql.Stmt
	}
	tests := []struct {
		name    string
		fields  fields
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
			if err := ds.getSelectPreparedstatements(); (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.getSelectPreparedstatements() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
