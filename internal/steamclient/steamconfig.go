package steamclient

import (
	"encoding/json"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

// SteamConfig struct holds the values as read from the configuration file
// config.json on startup
type SteamConfig struct {
	SteamAPIKey     string   `json:"SteamAPIKey"`
	SteamIDs        []string `json:"SteamIDs"`
	HistoryInterval int      `json:"HistoryInterval"`
	UpdateInterval  int      `json:"UpdateInterval"`
	PostgresUser    string   `json:"db_user"`
	PostgresPass    string   `json:"db_pass"`
	PostgresDbName  string   `json:"db_name"`
	PostgresDbHost  string   `json:"db_host"`
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
	if err != nil {
		b, _ := ioutil.ReadAll(jsonFile)
		log.Println("Error during startup. Config file used:", configPath)
		log.Println(b)
	}
	return conf, err
}
