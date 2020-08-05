package steamclient

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// SteamConfig struct holds the values as read from the configuration file
// config.json on startup
type SteamConfig struct {
	SteamAPIKey     string   `json:"SteamAPIKey"`
	SteamIDs        []string `json:"SteamIDs"`
	HistoryInterval int      `json:"HistoryInterval"`
	UpdateInterval  int      `json:"UpdateInterval"`
	MongoUser       string   `json:"mongo_user"`
	MongoPass       string   `json:"mongo_pass"`
	MongoHost       string   `json:"mongo_host"`
	MongoTestDbName string   `json:"test_db_name"`
	MongoDbName     string   `json:"db_name"`
}

// NewSteamConfig creates a new configuration struct from a path pointing to
// the json configuration file. It will return an error if the file does not exist
func NewSteamConfig(configPath string) (SteamConfig, error) {

	conf := SteamConfig{}

	jsonFile, err := os.Open(configPath)
	if err != nil {
		return conf, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &conf)
	return conf, err
}
