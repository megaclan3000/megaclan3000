package main

import (
	// "reflect"
	"testing"

	"github.com/megaclan3000/megaclan3000/internal/demoparser"
	"github.com/megaclan3000/megaclan3000/internal/steamclient"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func TestNewDataStorage(t *testing.T) {

	conf, err := steamclient.NewSteamConfig("./config.json")

	if err != nil {
		panic(err)
	}

	mongoURI := "mongodb+srv://" + conf.MongoUser + ":" + conf.MongoPass + "@" + conf.MongoHost + "/" + conf.MongoTestDbName + "?retryWrites=true&w=majority"

	tests := []struct {
		name    string
		want    *DataStorage
		wantErr bool
	}{
		{
			name:    "Connect to test database",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Just check if the connection succeeds and no errors are thrown
			_, err := NewDataStorage(mongoURI, conf.MongoTestDbName)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDataStorage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDataStorage_SaveMatch(t *testing.T) {

	conf, err := steamclient.NewSteamConfig("./config.json")

	if err != nil {
		panic(err)
	}

	mongoURI := "mongodb+srv://" + conf.MongoUser + ":" + conf.MongoPass + "@" + conf.MongoHost + "/" + conf.MongoTestDbName + "?retryWrites=true&w=majority"

	tests := []struct {
		name          string
		matchDemoPath string
		want          interface{}
		wantErr       bool
	}{
		{
			name:          "Save demo1.dem",
			matchDemoPath: "internal/demoparser/testdata/demo1.dem",
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ds := prepareTestDB(mongoURI, conf.MongoTestDbName)
			parser := demoparser.NewMyParser()
			match, err := parser.Parse(tt.matchDemoPath)

			if err != nil {
				t.Errorf("DataStorage.SaveMatch() aborted, demo parsing error: %v", err)
			}

			err = ds.SaveMatch(match)

			if (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.SaveMatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func prepareTestDB(mongoURI, dbName string) *DataStorage {
	var ds *DataStorage
	var err error

	//TODO extact mongo URI
	if ds, err = NewDataStorage(mongoURI, dbName); err != nil {
		panic(err)
	}

	if err = ds.database.Drop(ds.ctx); err != nil {
		panic(err)
	}

	return ds

}
